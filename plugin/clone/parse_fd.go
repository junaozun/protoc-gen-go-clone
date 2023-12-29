package clone

import (
	"strings"
	"unicode"

	"github.com/junaozun/protoc-gen-go-clone/generator"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

type IFieldDesc interface {
	clone(funcName string) ([]byte, error)
	isValueBase() bool
	getValueType() string
	getValuePrefix() string
	addFuncWay(temp *genTemp, funcWay map[string]IFieldDesc) (string, string)
	getFinalVal() string
}

type parseFd struct {
	gen           *generator.Generator
	sourceFileMap map[string]*filesAllMsg // fileName:fileMsg
	annObject     map[string][]string     // pkgName:有注释对象
}

// all files message
type filesAllMsg struct {
	isSourceFile bool // 是否传入的源文件
	isProto3     bool
	packageName  string                  // 文件所属包
	allObjects   map[string][]IFieldDesc // 对象名:fieldsDesc
}

type mapEntry struct {
	fieldName   string // 字段名
	keyType     string // key 类型 string
	valueType   string // value 类型 string,Apple
	valuePrefix string // value 的包名
	isBaseType  bool   // value 是否基础类型
	needPrefix  bool
	finalValue  string
}

type array struct {
	fieldName   string
	valueType   string
	valuePrefix string
	isBaseType  bool
	needPrefix  bool
	finalValue  string
}

type enum struct {
	isProto3    bool
	fieldName   string
	valueType   string
	valuePrefix string
	needPrefix  bool
	finalValue  string
}

type message struct {
	fieldName   string
	valueType   string
	valuePrefix string
	needPrefix  bool
	finalValue  string
}

type base struct {
	isProto3  bool
	fieldName string
	valueType string
}

func NewParseFd(gen *generator.Generator) *parseFd {
	return &parseFd{
		gen:           gen,
		sourceFileMap: make(map[string]*filesAllMsg),
		annObject:     make(map[string][]string),
	}
}

func (p *parseFd) ParseFileDesc(file *generator.FileDescriptor, cb func(announce []*descriptor.SourceCodeInfo_Location) []string) bool {
	fsGen := &filesAllMsg{
		isSourceFile: p.isSourceFile(file.GetName()),
		isProto3:     fileIsProto3(file.FileDescriptorProto),
		packageName:  file.GetPackage(),
		allObjects:   make(map[string][]IFieldDesc),
	}
	p.sourceFileMap[file.GetName()] = fsGen
	for _, m := range file.FileDescriptorProto.MessageType {
		middleMap := make(map[string]IFieldDesc)
		p.parseAllFields(m, fsGen, file.GetName(), file.GetPackage(), middleMap, "")
	}
	// get annotation object
	p.annObject[file.GetPackage()] = append(p.annObject[file.GetPackage()], cb(file.GetSourceCodeInfo().GetLocation())...)

	if len(p.sourceFileMap) != len(p.gen.Request.ProtoFile) {
		return false
	}
	return true
}

func convertFieldDesc(valueTypes string) (valueType string, valuePrefix string, isBaseType bool) {
	if strings.Contains(valueTypes, ".") {
		res := strings.Split(valueTypes, ".")
		if len(res) != 2 {
			panic("convertFieldDesc err")
		}
		return res[1], res[0], false
	}
	return valueTypes, "", true
}

func isLowerFirstChar(name string) bool {
	return unicode.IsLower(rune(name[0]))
}

// Given a type name defined in a .proto, return its name as we will print it.
func (p *parseFd) typeName(str string) string {
	return p.gen.TypeName(p.objectNamed(str))
}

// Given a type name defined in a .proto, return its object.
// Also record that we're using it, to guarantee the associated import.
func (g *parseFd) objectNamed(name string) generator.Object {
	g.gen.RecordTypeUse(name)
	return g.gen.ObjectNamed(name)
}

func (p *parseFd) parseAllFields(m *descriptor.DescriptorProto, fsGen *filesAllMsg, fileName, filePkgName string, middle map[string]IFieldDesc, parent string) {
	mName := generator.CamelCase(m.GetName())
	// 是内嵌类型的情况下，判断一下内嵌message首字母是大写还是小写
	if parent != "" {
		if isLowerFirstChar(m.GetName()) {
			mName = parent + mName
		} else {
			mName = parent + "_" + mName
		}
	}

	fsGen.allObjects[mName] = nil

	// enum 和 message都可能以嵌套形式存在
	for _, l := range m.NestedType {
		if !l.GetOptions().GetMapEntry() {
			p.parseAllFields(l, fsGen, fileName, filePkgName, middle, mName)
			continue
		}
		_, keyType := getValueTypeAndValue(l.GetField()[0])
		_, valueTypes := getValueTypeAndValue(l.GetField()[1]) // 带有前缀
		if valueTypes == "" {
			continue
		}
		valueType, valuePrefix, isBaseType := convertFieldDesc(valueTypes)
		var finalValue string
		aa := l.GetField()[1].GetTypeName()
		if aa != "" {
			finalValue = p.typeName(aa)
		}
		middle[l.GetName()] = &mapEntry{
			fieldName:   convertFieldName(l.GetName()[:len(l.GetName())-5]), // map类型尾部都拼接了Entry，所以这里要删掉,
			keyType:     keyType,
			valueType:   valueType,
			valuePrefix: valuePrefix,
			isBaseType:  isBaseType,
			finalValue:  finalValue,
		}
	}

	for _, f := range m.Field {
		iFieldDesc, valueTypes := getValueTypeAndValue(f)
		valueType, valuePrefix, isBaseType := convertFieldDesc(valueTypes)
		aa := f.GetTypeName()
		var finalValue string
		if aa != "" {
			finalValue = p.typeName(aa)
		}
		var (
			fd IFieldDesc
		)
		switch iFieldDesc.(type) {
		case *mapEntry:
			fd = middle[strings.Title(f.GetName())+"Entry"]
		case *array:
			fd = &array{
				fieldName:   convertFieldName(f.GetName()),
				valueType:   valueType,
				valuePrefix: valuePrefix,
				isBaseType:  isBaseType,
				finalValue:  finalValue,
			}
		case *enum:
			newStr := f.GetTypeName()[1:]
			ls := strings.Split(newStr, ".")
			var eName string
			if len(ls) == 2 {
				eName = strings.Title(ls[1])
			} else {
				// 对内嵌enum的处理
				eName = strings.Title(strings.Split(spiltValueType(newStr), ".")[1])
			}
			fd = &enum{
				isProto3:    fsGen.isProto3,
				fieldName:   convertFieldName(f.GetName()),
				valueType:   eName,
				valuePrefix: ls[0],
				finalValue:  finalValue,
			}
			fsGen.allObjects[eName] = append(fsGen.allObjects[eName], fd)
		case *message:
			fd = &message{
				fieldName:   convertFieldName(f.GetName()),
				valueType:   valueType,
				valuePrefix: valuePrefix,
				finalValue:  finalValue,
			}
		case *base:
			fd = &base{
				isProto3:  fsGen.isProto3,
				fieldName: convertFieldName(f.GetName()),
				valueType: valueType,
			}
		default:
			continue
		}
		fsGen.allObjects[mName] = append(fsGen.allObjects[mName], fd)
	}
}

func convertFieldName(fieldName string) string {
	var newString string
	var nextBig bool
	for _, v := range fieldName {
		if string(v) == "_" {
			nextBig = true
			continue
		}
		if nextBig {
			newString += strings.Title(string(v))
			nextBig = false
			continue
		}
		newString += string(v)
	}
	return newString
}

func fileIsProto3(file *descriptor.FileDescriptorProto) bool {
	return file.GetSyntax() == "proto3"
}

func getValueTypeAndValue(fl *descriptor.FieldDescriptorProto) (IFieldDesc, string) {
	if fl.GetLabel() == descriptor.FieldDescriptorProto_LABEL_REPEATED {
		switch fl.GetType() {
		case descriptor.FieldDescriptorProto_TYPE_STRING:
			return &array{}, "string"
		case descriptor.FieldDescriptorProto_TYPE_INT32:
			return &array{}, "int32"
		case descriptor.FieldDescriptorProto_TYPE_INT64:
			return &array{}, "int64"
		case descriptor.FieldDescriptorProto_TYPE_BYTES:
			return &array{}, "byte"
		case descriptor.FieldDescriptorProto_TYPE_UINT32:
			return &array{}, "uint32"
		case descriptor.FieldDescriptorProto_TYPE_UINT64:
			return &array{}, "uint64"
		case descriptor.FieldDescriptorProto_TYPE_BOOL:
			return &array{}, "bool"
		case descriptor.FieldDescriptorProto_TYPE_FLOAT:
			return &array{}, "float32"
		case descriptor.FieldDescriptorProto_TYPE_DOUBLE:
			return &array{}, "float64"
		case descriptor.FieldDescriptorProto_TYPE_ENUM:
			return &array{}, "enum"
		case descriptor.FieldDescriptorProto_TYPE_MESSAGE:
			// 可能是数组对象类型，也可能是map类型
			str := fl.GetTypeName()
			if len(str) == 0 {
				panic("message_type not found typeName")
			}
			newStr := str[1:]
			res := strings.Split(newStr, ".")
			if len(res) < 2 {
				panic("message_type split error")
			}
			if res[len(res)-1] == strings.Title(fl.GetName())+"Entry" {
				return &mapEntry{}, ""
			}
			if len(res) != 2 {
				newStr = spiltValueType(newStr)
			}
			return &array{}, newStr
		default:
			panic("label LABEL_REPEATED ,type_err")
		}
	} else if fl.GetLabel() == descriptor.FieldDescriptorProto_LABEL_OPTIONAL || fl.GetLabel() == descriptor.FieldDescriptorProto_LABEL_REQUIRED {
		switch fl.GetType() {
		case descriptor.FieldDescriptorProto_TYPE_STRING:
			return &base{}, "string"
		case descriptor.FieldDescriptorProto_TYPE_INT32:
			return &base{}, "int32"
		case descriptor.FieldDescriptorProto_TYPE_INT64:
			return &base{}, "int64"
		case descriptor.FieldDescriptorProto_TYPE_BYTES:
			return &base{}, "byte"
		case descriptor.FieldDescriptorProto_TYPE_UINT32:
			return &base{}, "uint32"
		case descriptor.FieldDescriptorProto_TYPE_UINT64:
			return &base{}, "uint64"
		case descriptor.FieldDescriptorProto_TYPE_BOOL:
			return &base{}, "bool"
		case descriptor.FieldDescriptorProto_TYPE_ENUM:
			return &enum{}, "enum"
		case descriptor.FieldDescriptorProto_TYPE_FLOAT:
			return &base{}, "float32"
		case descriptor.FieldDescriptorProto_TYPE_DOUBLE:
			return &base{}, "float64"
		case descriptor.FieldDescriptorProto_TYPE_MESSAGE:
			str := fl.GetTypeName()
			if len(str) == 0 {
				panic("message_type not found typeName")
			}
			newStr := str[1:]
			res := strings.Split(newStr, ".")
			if len(res) < 2 {
				panic("message_type split error")
			}
			if len(res) != 2 {
				newStr = spiltValueType(newStr)
			}
			return &message{}, newStr
		default:
			panic("label LABEL_OPTIONAL ,type_err")
		}
	} else {
		panic("getLabelType label type error")
	}
}

// 先忽略第一个分隔符. 从遇到第二个分隔符开始，遇到大写下划线拼接，遇到小写转成大写再拼接
func spiltValueType(valueType string) string {
	var (
		newStr string
	)
	str := strings.Split(valueType, ".")
	for i, v := range str {
		if i == 1 {
			v = "." + strings.Title(v)
		} else if i > 1 {
			if !isLowerFirstChar(v) {
				v = "_" + strings.Title(v)
			} else {
				v = strings.Title(v)
			}
		}
		newStr += v
	}
	return newStr
}

func (p *parseFd) isSourceFile(file string) bool {
	for _, sourceFile := range p.gen.Request.GetFileToGenerate() {
		if file == sourceFile {
			return true
		}
	}
	return false
}

func (m *mapEntry) isValueBase() bool {
	return m.isBaseType
}

func (a *array) isValueBase() bool {
	return a.isBaseType
}

func (e *enum) isValueBase() bool {
	return true
}

func (m *message) isValueBase() bool {
	return false
}

func (b *base) isValueBase() bool {
	return true
}

func (m *mapEntry) getValueType() string {
	return m.valueType
}

func (a *array) getValueType() string {
	return a.valueType
}

func (e *enum) getValueType() string {
	return e.valueType
}

func (m *message) getValueType() string {
	return m.valueType
}

func (b *base) getValueType() string {
	return b.valueType
}

func (m *mapEntry) getValuePrefix() string {
	return m.valuePrefix
}

func (a *array) getValuePrefix() string {
	return a.valuePrefix
}

func (e *enum) getValuePrefix() string {
	return ""
}

func (m *message) getValuePrefix() string {
	return m.valuePrefix
}

func (b *base) getValuePrefix() string {
	return ""
}

func (m *mapEntry) getFinalVal() string {
	return m.finalValue
}

func (a *array) getFinalVal() string {
	return a.finalValue
}

func (e *enum) getFinalVal() string {
	return e.finalValue
}

func (m *message) getFinalVal() string {
	return m.finalValue
}

func (b *base) getFinalVal() string {
	return b.fieldName
}

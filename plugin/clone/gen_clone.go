package clone

import (
	"sort"
	"strings"

	"github.com/junaozun/protoc-gen-go-clone/pkg/tmpl"
)

type genClone struct {
	parseFd        *parseFd
	annotationFlag map[string]*announceObj // 注释对象名->announceObj
	funcWayStore   map[string]IFieldDesc   // funcName->IFieldDesc
}

type announceObj struct {
	objPkgName string       // 注释对象所在包名
	genPkgName string       // 最终生成文件所在包名
	allRefers  []*referObjs // 对象直接或间接引用的所有对象
}

// 哪个包下哪个文件中的对象
type referObjs struct {
	objName     string // 对象名
	valuePrefix string // 对象前缀
	parentObj   string
	finalValue  string
}

type convertTemps struct {
	key   string
	value IFieldDesc
}

func NewGenClone(parseFd *parseFd) *genClone {
	return &genClone{
		parseFd:        parseFd,
		annotationFlag: make(map[string]*announceObj),
		funcWayStore:   make(map[string]IFieldDesc),
	}
}

func (g *genClone) ContainString(imp string) bool {
	for _, v := range g.annotationFlag {
		for _, m := range v.allRefers {
			if m.valuePrefix == imp {
				return true
			}
		}
	}
	return false
}

func (g *genClone) GenClone(pkgName string) {
	for k, v := range g.parseFd.annObject {
		for _, vv := range v {
			g.annotationFlag[vv] = &announceObj{
				objPkgName: k,
				genPkgName: pkgName,
			}
		}
	}

	for _, fileInfo := range g.parseFd.sourceFileMap {
		if !fileInfo.isSourceFile {
			continue
		}
		for structName, allFields := range fileInfo.allObjects {
			g.parseObject(structName, allFields)
		}
	}

	g.genParentClone()
	g.genSonClone()
	g.genWayClone()
}

func (g *genClone) parseObject(structName string, allFields []IFieldDesc) {
	// 解析出这个结构体直接引用和间接引用的所有对象
	for _, v := range allFields {
		// 过滤非对象字段
		if v.isValueBase() {
			continue
		}
		// 过滤一下未注释的对象
		mm, ok := g.annotationFlag[structName]
		if !ok {
			return
		}
		// 过滤一下已经存在的对象名
		if containString(mm.allRefers, v.getValueType(), v.getValuePrefix()) {
			continue
		}

		prefix := v.getValuePrefix()
		g.annotationFlag[structName].allRefers = append(g.annotationFlag[structName].allRefers, &referObjs{
			objName:     v.getValueType(),
			valuePrefix: prefix,
			parentObj:   structName,
			finalValue:  v.getFinalVal(),
		})
		// 查找子对象
		rs, _ := g.findObjAllFields(prefix, v.getValueType())
		g.parseObject(structName, rs)
	}
}

func (g *genClone) findObjAllFields(valuePrefix, valueType string) ([]IFieldDesc, bool) {
	for _, fileMsg := range g.parseFd.sourceFileMap {
		if fileMsg.packageName != valuePrefix {
			continue
		}
		obj, ok := fileMsg.allObjects[strings.Title(valueType)]
		if !ok {
			continue
		}
		return obj, fileMsg.isProto3
	}
	panic("not findObjAllFields :" + valuePrefix + valueType)
}

type tplTemp struct {
	FuncName []string
	ObName   string
	Fields   map[string]string
}

type genTemp struct {
	genPkgName  string
	isProto3    bool
	objName     string
	finalValue  string
	valuePrefix string
	parentObj   string
	fDesc       []IFieldDesc
}

func (g *genTemp) isNeedPrefix(prefix string) bool {
	return g.genPkgName != prefix
}

// generate export way
func (g *genClone) genParentClone() {
	parents := make([]*genTemp, 0, 10)
	for k, v := range g.annotationFlag {
		allFields, isProto3 := g.findObjAllFields(v.objPkgName, k)
		parents = append(parents, &genTemp{
			genPkgName: v.genPkgName,
			parentObj:  k,
			isProto3:   isProto3,
			objName:    k,
			fDesc:      allFields,
		})
	}
	sort.Slice(parents, func(i, j int) bool {
		return parents[i].objName < parents[j].objName
	})
	for _, par := range parents {
		input := &tplTemp{
			ObName: par.objName,
			Fields: make(map[string]string),
		}
		for _, v := range par.fDesc {
			filedName, value := v.addFuncWay(par, g.funcWayStore)
			input.Fields[filedName] = value
		}
		b, err := tmpl.Render(tempParent, input)
		if err != nil {
			g.parseFd.gen.Error(err)
		}
		g.parseFd.gen.P(string(b))
	}
}

// return fieldName,value
func (m *mapEntry) addFuncWay(temp *genTemp, funcWay map[string]IFieldDesc) (string, string) {
	fn := strings.Title(m.fieldName)
	strings.Split(fn, "_")
	value := strings.Title(m.valueType)
	if !m.isBaseType {
		value = m.valuePrefix + "_" + value
	}
	m.needPrefix = temp.isNeedPrefix(m.valuePrefix)
	funcName := "clone_" + strings.Title(temp.parentObj) + "_Map_" + m.keyType + "_" + value
	if _, ok := funcWay[funcName]; !ok {
		funcWay[funcName] = m
	}
	return fn, funcName + "(in." + fn + ")"
}

func (a *array) addFuncWay(temp *genTemp, funcWay map[string]IFieldDesc) (string, string) {
	fn := strings.Title(a.fieldName)

	if a.isBaseType {
		return fn, "append(in." + fn + "[:0:0], in." + fn + "...)"
	}
	funcName := "clone_" + strings.Title(temp.parentObj) + "_Array_" + a.valuePrefix + "_" + strings.Title(a.valueType)
	a.needPrefix = temp.isNeedPrefix(a.valuePrefix)
	if _, ok := funcWay[funcName]; !ok {
		funcWay[funcName] = a
	}
	return fn, funcName + "(in." + fn + ")"
}

func (e *enum) addFuncWay(temp *genTemp, funcWay map[string]IFieldDesc) (string, string) {
	fn := strings.Title(e.fieldName)
	funcName := "clone" + "Enum" + temp.parentObj + e.valueType
	if temp.isProto3 {
		return fn, "in." + fn
	}
	e.needPrefix = temp.isNeedPrefix(e.valuePrefix)
	// pb2需要对enum生成clone方法
	if _, ok := funcWay[funcName]; !ok {
		funcWay[funcName] = e
	}
	return fn, funcName + "(in." + fn + ")"
}

func (m *message) addFuncWay(temp *genTemp, funcWay map[string]IFieldDesc) (string, string) {
	fn := strings.Title(m.fieldName)
	m.needPrefix = temp.isNeedPrefix(m.valuePrefix)
	funcName := "clone_" + strings.Title(temp.parentObj) + "_" + m.valuePrefix + "_" + strings.Title(m.valueType)
	return fn, funcName + "(in." + fn + ")"
}

func (b *base) addFuncWay(temp *genTemp, funcWay map[string]IFieldDesc) (string, string) {
	fn := strings.Title(b.fieldName)
	if temp.isProto3 {
		return fn, "in." + fn
	}
	funcName := "clone" + strings.Title(b.valueType) + "Pointer"
	// pb2需要对base生成clone方法
	if _, ok := funcWay[funcName]; !ok {
		funcWay[funcName] = b
	}
	return fn, funcName + "(in." + fn + ")"
}

// generate no export ways
func (g *genClone) genSonClone() {
	sons := make([]*genTemp, 0, 10000)
	for ob, v := range g.annotationFlag {
		for _, son := range v.allRefers {
			allFields, isProto3 := g.findObjAllFields(son.valuePrefix, son.objName)
			sons = append(sons, &genTemp{
				genPkgName:  v.genPkgName,
				isProto3:    isProto3,
				objName:     son.objName,
				finalValue:  son.finalValue,
				valuePrefix: son.valuePrefix,
				parentObj:   ob,
				fDesc:       allFields,
			})
		}

	}
	sort.Slice(sons, func(i, j int) bool {
		return sons[i].objName < sons[j].objName
	})
	for _, son := range sons {
		value := strings.Title(son.objName)
		funcName := []string{son.parentObj, son.valuePrefix, strings.Title(son.objName)}
		if son.valuePrefix != son.genPkgName {
			value = son.valuePrefix + "." + value
		}
		if son.finalValue != "" {
			value = son.finalValue
		}
		input := &tplTemp{
			FuncName: funcName,
			ObName:   value,
			Fields:   make(map[string]string),
		}

		for _, v := range son.fDesc {
			filedName, valueFunc := v.addFuncWay(son, g.funcWayStore)
			input.Fields[filedName] = valueFunc
		}
		b, err := tmpl.Render(tempSon, input)
		if err != nil {
			g.parseFd.gen.Error(err)
		}
		g.parseFd.gen.P(string(b))
	}
}

func (g *genClone) genWayClone() {
	sortWays := make([]*convertTemps, 0, len(g.funcWayStore))
	for k, v := range g.funcWayStore {
		sortWays = append(sortWays, &convertTemps{
			key:   k,
			value: v,
		})
	}
	sort.Slice(sortWays, func(i, j int) bool {
		return sortWays[i].key < sortWays[j].key
	})
	for _, v := range sortWays {
		b, err := v.value.clone(v.key)
		if err != nil {
			g.parseFd.gen.Error(err)
		}
		g.parseFd.gen.P(string(b))
	}
}

func (m *mapEntry) clone(funcName string) ([]byte, error) {
	valueType := strings.Title(m.valueType)
	data := map[string]interface{}{
		"funcName": funcName,
		"keyType":  m.keyType,
		"objName":  m.valueType,
	}
	if !m.isBaseType {
		obName := "*" + valueType
		data["sonMsg"] = []string{getParentObjectName(funcName), m.valuePrefix, valueType}
		if m.finalValue != "" {
			obName = "*" + m.finalValue
		}
		data["objName"] = obName
	}
	return tmpl.Render(tempMap, data)
}

func (a *array) clone(funcName string) ([]byte, error) {
	valueType := strings.Title(a.valueType)
	data := map[string]interface{}{
		"funcName": funcName,
		"objName":  valueType,
		"sonMsg":   []string{getParentObjectName(funcName), a.valuePrefix, valueType},
	}
	if a.finalValue != "" {
		data["objName"] = a.finalValue
	}
	return tmpl.Render(tempArray, data)
}

func (e *enum) clone(funcName string) ([]byte, error) {
	final := e.valueType
	if e.finalValue != "" {
		final = e.finalValue
	}
	data := map[string]interface{}{
		"funcName": funcName,
		"objName":  final,
	}
	return tmpl.Render(tempEnum, data)
}

func (m *message) clone(funcName string) ([]byte, error) {
	return nil, nil
}

func (b *base) clone(funcName string) ([]byte, error) {
	if b.isProto3 {
		return nil, nil
	}
	data := map[string]interface{}{
		"funcName":  funcName,
		"valueType": b.valueType,
	}
	return tmpl.Render(tempBase, data)
}

func getParentObjectName(funcName string) string {
	return strings.Split(funcName, "_")[1]
}

func containString(arr []*referObjs, valueType string, valuePrefix string) bool {
	for _, v := range arr {
		if v.objName == valueType && v.valuePrefix == valuePrefix {
			return true
		}
	}
	return false
}

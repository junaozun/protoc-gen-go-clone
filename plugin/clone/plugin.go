package clone

import (
	"strings"

	"gitlab.uuzu.com/war/pbtool/generator"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

// Plugin is an implementation of the Go protocol buffer compiler's
// plugin architecture.  It generates bindings for Plugin support.
type Plugin struct {
	gen      *generator.Generator
	parseFd  *parseFd
	importFd map[string]bool
}

// Name returns the name of this plugin, "Plugin".
func (g *Plugin) Name() string {
	return "clone"
}

// The names for packages imported in the generated code.
// They may vary from the final path component of the import path
// if the name is used by other packages.

// Init initializes the plugin.
func (g *Plugin) Init(gen *generator.Generator) {
	g.gen = gen
	g.parseFd = NewParseFd(gen)
	g.importFd = make(map[string]bool)
}

// Given a type name defined in a .proto, return its object.
// Also record that we're using it, to guarantee the associated import.
func (g *Plugin) objectNamed(name string) generator.Object {
	g.gen.RecordTypeUse(name)
	return g.gen.ObjectNamed(name)
}

// Given a type name defined in a .proto, return its name as we will print it.
func (g *Plugin) typeName(str string) string {
	return g.gen.TypeName(g.objectNamed(str))
}

// P forwards to g.gen.P.
func (g *Plugin) P(args ...interface{}) { g.gen.P(args...) }

// Generate generates code for the services in the given file.
func (g *Plugin) Generate(file *generator.FileDescriptor) {
	ok := g.parseFd.ParseFileDesc(file, func(comments []*descriptor.SourceCodeInfo_Location) []string {
		return parseSpecifiedComments(comments)
	})
	for k := range g.gen.GetImport() {
		g.importFd[string(k)] = true
	}

	if !ok {
		return
	}

	clone := NewGenClone(g.parseFd)
	clone.GenClone(file.GetPackage())
	for k := range g.importFd {
		// 需要过滤一下无用的import
		m := strings.Split(k, "/")
		imp := m[len(m)-1]
		if !clone.ContainString(imp) {
			continue
		}
		g.gen.AddImport(generator.GoImportPath(k))
	}

}

// GenerateImports generates the import declaration for this file.
func (g *Plugin) GenerateImports(file *generator.FileDescriptor) {
	// g.P("var _ = context.Backgroud()")
}

// user指定解析注释格式
func parseSpecifiedComments(comments []*descriptor.SourceCodeInfo_Location) []string {
	var res []string
	for _, lc := range comments {
		if lc.GetLeadingComments() == "" {
			continue
		}
		var strName string
		for i, v := range lc.GetLeadingComments() {
			if string(v) == "|" {
				strName = lc.GetLeadingComments()[i+1:]
				break
			}
		}
		if strName == "" {
			continue
		}
		strName = strings.Replace(strName, " ", "", -1)
		// 去除换行符
		strName = strings.Replace(strName, "\n", "", -1)
		res = append(res, strName)
	}
	return res
}

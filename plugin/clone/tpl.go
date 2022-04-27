package clone

var tempMap = `func {{.funcName}}(in map[{{.keyType}}]{{.objName}}) map[{{.keyType}}]{{.objName}} {
				if in == nil {
					return nil
				}
				a := make(map[{{.keyType}}]{{.objName}},len(in))
				for k,v := range in {
                {{if .sonMsg}}a[k] = clone{{- range $i, $v := .sonMsg -}}
		     	_{{$v}}{{- end -}}(v)
				{{else}}a[k] = v{{end}}}
				return a
				}
`

var tempArray = `func {{.funcName}}(in []*{{.objName}}) []*{{.objName}} {
				if in == nil {
				return nil
				}
				a := make([]*{{.objName}},len(in))
				for k,v := range in {
				a[k] = clone{{- range $i, $v := .sonMsg -}}
		     	_{{$v}}
	          	{{- end -}}(v)
				}
				return a
				}
`

var tempEnum = `func {{.funcName}}(m *{{.objName}}) *{{.objName}} {
				if m == nil {
					return nil
				}
				p := new({{.objName}})
				*p = *m
				return p
				}
`
var tempBase = `func {{.funcName}}(m *{{.valueType}}) *{{.valueType}} {
				if m == nil {
					return nil
				}
				p := new({{.valueType}})
				*p = *m
				return p
				}
`

var tempParent = `func(in *{{.ObName}}) Clone() *{{.ObName}} {
				if in == nil {
				return nil
				}
				out := &{{.ObName}}{}{{range $index, $value := .Fields}}
                out.{{$index}} = {{$value}}{{end}}
				return out
				}
				
`

var tempSon = `func clone{{- range $i, $v := .FuncName -}}
		     	_{{$v}}
	          	{{- end -}}(in *{{.ObName}}) *{{.ObName}} {
				if in == nil {
				return nil
				}
				out := &{{.ObName}}{}{{range $index, $value := .Fields}}
                out.{{$index}} = {{$value}}{{end}}
				return out
				}
				
`

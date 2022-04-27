package tmpl

import (
	"bytes"
	"text/template"
)

// Render 渲染模板
func Render(tmp string, data interface{}) ([]byte, error) {
	tpl, err := template.New("tmpl").Parse(tmp)
	if err != nil {
		return nil, err
	}
	return RenderTemplate(tpl, data)
}

// RenderTemplate 模板执行
func RenderTemplate(t *template.Template, data interface{}) ([]byte, error) {
	w := &bytes.Buffer{}
	if err := t.Execute(w, data); nil != err {
		return []byte(""), err
	}
	return w.Bytes(), nil
}

package Tmpl

import (
	"bytes"
	"text/template"
)

func Tmpl(Arr map[string]string, Tpl string) string {
	var body bytes.Buffer
	tmpl, err := template.New("go2Msg").Parse(Tpl)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(&body, Arr)
	if err != nil {
		panic(err)
	}
	return body.String()
}

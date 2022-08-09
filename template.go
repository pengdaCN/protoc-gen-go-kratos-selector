package main

import (
	_ "embed"
	"github.com/Masterminds/sprig/v3"
	"text/template"
)

//go:embed tmpl.go.tmpl
var tmplStr string
var tmpl *template.Template

func init() {
	tmpl = template.Must(template.New("selector").Funcs(sprig.FuncMap()).Parse(tmplStr))
}

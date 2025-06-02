package main

import (
	"os"
	"text/template"
)

func main() {
	var layoutTmpl = "./layout.tex.tmpl"
	layout, err := template.New(layoutTmpl).ParseFiles(layoutTmpl)
	if err != nil {
		panic(err)
	}
	err = layout.Execute(os.Stdout, nil)
	if err != nil {
		panic(err)
	}
	var stylingTmpl = "./styling.cls.tmpl"
	style, err := template.New(stylingTmpl).ParseFiles(stylingTmpl)
	if err != nil {
		panic(err)
	}
	err = style.Execute(os.Stdout, nil)
	if err != nil {
		panic(err)
	}
}

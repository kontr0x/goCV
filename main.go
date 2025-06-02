package main

import (
	"fmt"

	"github.com/coffeemakingtoaster/cv-gen/pkg/content"
)

func main() {
	layoutTmpl := "layout.tex.tmpl"
	stylingTmpl := "resume.cls.tmpl"

	data, err := content.ParseContentFromYaml("./content.yaml")
	if err != nil {
		panic(err)
	}
	for _, entry := range data {
		dir := fmt.Sprintf("./%s-out", entry.Content.Version)
		err := content.RenderTemplate(layoutTmpl, stylingTmpl, dir, entry)
		if err != nil {
			panic(err)
		}
		err = content.BuildTemplate(dir)
		if err != nil {
			panic(err)
		}
	}
}

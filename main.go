package main

import (
	"fmt"
	"os"

	"github.com/coffeemakingtoaster/cv-gen/pkg/content"
)

var defaultContentYamlPath = "./content.yaml"

func main() {
	contentPath := defaultContentYamlPath
	if len(os.Args) > 1 {
		contentPath = os.Args[1]
	}
	fmt.Printf("Reading content from %s\n", contentPath)
	data, err := content.ParseContentFromYaml(contentPath)
	if err != nil {
		panic(err)
	}
	for _, entry := range data {
		dir := fmt.Sprintf("./%s-out", entry.Content.Version)
		// TODO: Add command line flag parsing to allow for custom passed paths
		err := content.RenderTemplate("", "", dir, entry)
		if err != nil {
			panic(err)
		}
		err = content.BuildTemplate(dir)
		if err != nil {
			panic(err)
		}
	}
}

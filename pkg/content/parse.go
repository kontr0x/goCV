package content

import (
	_ "embed"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"

	"gopkg.in/yaml.v3"
)

//go:embed layout.tex.tmpl
var layoutTemplate string

//go:embed resume.cls.tmpl
var styleTemplate string

func ParseContentFromYaml(path string) ([]TemplateData, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return []TemplateData{}, err
	}
	var resume Resume
	if err := yaml.Unmarshal(data, &resume); err != nil {
		return []TemplateData{}, err
	}
	result := []TemplateData{}
	for version, content := range resume.Versions {
		res := TemplateData{
			Style:   StyleTemplateData{Version: version},
			Content: ContentTemplateData{StaticContent: resume.Static, Version: version, Content: content},
		}
		result = append(result, res)

	}
	return result, nil
}

func loadTemplate(name, path, fallback string) (*template.Template, error) {
	if len(path) == 0 {
		return template.New(name).Parse(fallback)
	}
	return template.New(filepath.Base(path)).ParseFiles(path)
}

func renderTemplate(template *template.Template, targetpath string, data interface{}) error {
	var err error
	var f *os.File
	f, err = os.Create(targetpath)
	if err != nil {
		return err
	}
	err = template.Execute(f, data)
	if err != nil {
		return err
	}
	err = f.Close()
	return err
}

func ensureCleanDir(targetPath string) error {
	if _, err := os.Stat(targetPath); !os.IsNotExist(err) {
		err := os.RemoveAll(targetPath)
		if err != nil {
			return err
		}
	}
	err := os.MkdirAll(targetPath, 0755)
	return err
}

func RenderTemplate(layoutPath, stylePath, targetDirPath string, data TemplateData) error {
	fmt.Printf("Rendering for version: %s\n", data.Content.Version)
	err := ensureCleanDir(targetDirPath)
	if err != nil {
		return err
	}
	layoutTmpl, err := loadTemplate("layoutTemplate", layoutPath, layoutTemplate)
	if err != nil {
		return err
	}
	err = renderTemplate(layoutTmpl, fmt.Sprintf("%s/layout.tex", targetDirPath), data.Content)
	if err != nil {
		return err
	}
	styleTmpl, err := loadTemplate("styleTemplate", stylePath, styleTemplate)
	if err != nil {
		return err
	}
	err = renderTemplate(styleTmpl, fmt.Sprintf("%s/resume.cls", targetDirPath), data.Content)
	if err != nil {
		return err
	}
	return nil
}

func BuildTemplate(workDir string) error {
	cmd := exec.Command("latexmk", "-pdf", "./layout.tex")
	cmd.Dir = workDir
	cmd.Stderr = os.Stdout
	err := cmd.Run()
	return err
}

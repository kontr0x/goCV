package content

import (
	"fmt"
	"os"
	"os/exec"
	"text/template"

	"gopkg.in/yaml.v3"
)

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

func renderTemplate(templatePath, targetpath string, data interface{}) error {
	tmpl, err := template.New(templatePath).ParseFiles(templatePath)
	if err != nil {
		return err
	}
	var f *os.File
	f, err = os.Create(targetpath)
	if err != nil {
		return err
	}
	err = tmpl.Execute(f, data)
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
	err = renderTemplate(layoutPath, fmt.Sprintf("%s/layout.tex", targetDirPath), data.Content)
	if err != nil {
		return err
	}
	err = renderTemplate(stylePath, fmt.Sprintf("%s/resume.cls", targetDirPath), data.Content)
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

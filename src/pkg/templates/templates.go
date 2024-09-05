package templates

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"
)

type Scaffold struct {
	Name        string     `yaml:"scaffold_name"`
	Description string     `yaml:"description"`
	Params      []Param    `yaml:"params"`
	Templates   []Template `yaml:"templates"`
}

type Param struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Type        string `yaml:"type"`
	Default     string `yaml:"default"`
}

type Template struct {
	Source      string `yaml:"source"`
	Destination string `yaml:"destination"`
}

func TemplateRun() {
	yamlFile, err := ioutil.ReadFile("./templates/go/go-sveltekit/template.yaml")
	if err != nil {
		fmt.Println("Error reading YAML file:", err)
		return
	}

	var scaffold Scaffold
	err = yaml.Unmarshal(yamlFile, &scaffold)
	if err != nil {
		fmt.Println("Error unmarshalling YAML:", err)
		return
	}

	fmt.Printf("Scaffold: %+v\n", scaffold)
}

func FetchTemplates() []string {
	files, err := os.ReadDir("./template")
	if err != nil {
		fmt.Println("Error reading templates:", err)
		return nil
	}

	var templates []string
	for _, file := range files {
		if file.IsDir() {
			templates = append(templates, file.Name())
		}
	}
	return templates
}

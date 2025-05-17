package templates

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/sevenreup/panga/src/pkg/engine"
	"gopkg.in/yaml.v3"
)

func GetTemplate(template string) *engine.Scaffold {
	yamlFile, err := ioutil.ReadFile(template)
	if err != nil {
		fmt.Println("Error reading YAML file:", err)
		return nil
	}

	var scaffold engine.Scaffold
	err = yaml.Unmarshal(yamlFile, &scaffold)
	if err != nil {
		fmt.Println("Error unmarshalling YAML:", err)
		return nil
	}

	return &scaffold
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

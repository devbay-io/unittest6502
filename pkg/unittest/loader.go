package unittest

import (
	"os"
	"strings"

	"github.com/go-playground/validator"
	"gopkg.in/yaml.v3"
)

func UnitTestFromAsmFile(filePath string) (UnitTest, error) {
	b, err := os.ReadFile(filePath) // just pass the file name
	if err != nil {
		return UnitTest{}, err
	}
	return UnitTestFromString(string(b))
}

func UnitTestFromString(unitTestString string) (UnitTest, error) {
	var unitTest UnitTest
	var yamlContent []string

	lines := strings.Split(unitTestString, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, ";>") || strings.HasPrefix(line, "; >") {
			line = strings.TrimPrefix(line, ";>")
			line = strings.TrimPrefix(line, "; >")
			yamlContent = append(yamlContent, line)
		}
	}
	yamlData := strings.Join(yamlContent, "\n")
	err := yaml.Unmarshal([]byte(yamlData), &unitTest)
	if err != nil {
		return UnitTest{}, err
	}
	validate := validator.New()
	err = validate.Struct(unitTest)
	if err != nil {
		return UnitTest{}, err
	}

	return unitTest, nil
}

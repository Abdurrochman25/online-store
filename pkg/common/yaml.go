package common

import (
	"os"

	"gopkg.in/yaml.v3"
)

func YamlDecodeFile(filePath string, output interface{}) error {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(file, output); err != nil {
		return err
	}
	return nil
}

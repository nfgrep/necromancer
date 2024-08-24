package config

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

type EntitiesConfig struct {
	Entities map[string]yaml.Node `yaml:"entities"`
}

func ParseEntitiesConfig(data []byte) (*EntitiesConfig, error) {
	var config EntitiesConfig
	err := yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling YAML: %v", err)
	}
	return &config, nil
}

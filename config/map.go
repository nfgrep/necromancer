package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type MapData [][]string

type MapConfig struct {
	Maps map[string]MapData `yaml:"maps"`
}

func ParseMaps(configFile string) (map[string]MapData, error) {
	data, err := os.ReadFile(configFile)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %v", err)
	}

	var config MapConfig
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling YAML: %v", err)
	}

	return config.Maps, nil
}

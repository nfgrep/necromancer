package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type EntitiesConfig struct {
	Entities map[string]EntityConfig `yaml:"entities"`
}

// TODO: factor these into individual types?
// All possible fields for an entity
type EntityConfig struct {
	Type           string     `yaml:"type,omitempty"`
	Symbol         string     `yaml:"symbol"`
	TerminalSymbol string     `yaml:"terminal_symbol,omitempty"`
	Height         int        `yaml:"height,omitempty"`
	Texture        [][]string `yaml:"texture,omitempty"`
}

func ParseEntities(fpath string) (map[string]EntityConfig, error) {
	data, err := os.ReadFile(fpath)
	if err != nil {
		return nil, err
	}

	var config EntitiesConfig
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	// Set default type if not specified
	for id, entity := range config.Entities {
		if entity.Type == "" {
			entity.Type = id
			config.Entities[id] = entity
		}
	}

	return config.Entities, nil
}

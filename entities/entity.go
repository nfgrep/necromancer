package entities

import (
	"fmt"
	"os"
	"reflect"

	"github.com/nfgrep/necromancer/config"
)

// TODO: old code, remove
type EntityConfig struct {
	Type           string `yaml:"type,omitempty"`
	Symbol         string `yaml:"symbol"`
	TerminalSymbol string `yaml:"terminal_symbol"`
	Height         int    `yaml:"height"`
}

type EntityType string

type BaseEntity struct {
	Type   EntityType `yaml:"type"`
	Symbol string     `yaml:"symbol"`
}

type Entity interface{}

var entityTypes = make(map[EntityType]reflect.Type)

func RegisterEntityType(entityType EntityType, entityStruct Entity) {
	entityTypes[entityType] = reflect.TypeOf(entityStruct).Elem()
}

func ParseEntities(config *config.EntitiesConfig) (map[string]Entity, error) {

	entities := make(map[string]Entity)

	for id, node := range config.Entities {
		var base BaseEntity
		err := node.Decode(&base)
		if err != nil {
			return nil, fmt.Errorf("error decoding entity %s: %v", id, err)
		}

		if base.Type == "" {
			base.Type = EntityType(id)
		}

		entityType, ok := entityTypes[base.Type]
		if !ok {
			return nil, fmt.Errorf("unknown entity type for %s: %s", id, base.Type)
		}

		entity := reflect.New(entityType).Interface().(Entity)
		err = node.Decode(entity)
		if err != nil {
			return nil, fmt.Errorf("error decoding entity %s: %v", id, err)
		}

		entities[id] = entity
	}

	return entities, nil
}

func ParseEntitiesFromFile(fpath string) (map[string]Entity, error) {
	data, err := os.ReadFile(fpath)
	if err != nil {
		return nil, err
	}

	config, err := config.ParseEntitiesConfig(data)
	if err != nil {
		return nil, err
	}

	return ParseEntities(config)
}

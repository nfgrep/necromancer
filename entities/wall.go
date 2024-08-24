package entities

const WallType EntityType = "wall"

type WallEntity struct {
	BaseEntity     `yaml:",inline"`
	TerminalSymbol string     `yaml:"terminal_symbol,omitempty"`
	Height         int        `yaml:"height,omitempty"`
	Texture        [][]string `yaml:"texture,omitempty"`
}

func init() {
	RegisterEntityType(WallType, &WallEntity{})
}

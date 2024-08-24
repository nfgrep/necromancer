package entities

const PlayerType EntityType = "player"

type PlayerEntity struct {
	BaseEntity `yaml:",inline"`
}

func init() {
	RegisterEntityType(PlayerType, &PlayerEntity{})
}

package entities

const AirType EntityType = "air"

type AirEntity struct {
	BaseEntity `yaml:",inline"`
}

func init() {
	RegisterEntityType(AirType, &AirEntity{})
}

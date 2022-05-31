package enums

type TileTypeEnum int
type TaskTypeEnum int
type MouseModeEnum int

const (
	TileTypeEmpty TileTypeEnum = iota
	TileTypeDirt
	TileTypeGrass
	TileTypeRock
	TileTypeWater
)

const (
	TaskTypeNone TaskTypeEnum = iota
	TaskTypeMoveTo
	TaskTypeGather
	TaskTypeBuild
)

const (
	MouseModeNone MouseModeEnum = iota
	MouseModeBuild
	MouseModeGather
)

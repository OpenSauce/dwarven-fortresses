package enums

type TileTypeEnum int
type TaskTypeEnum int

const (
	Empty TileTypeEnum = iota
	Dirt
	Grass
	Rock
	Water

	None TaskTypeEnum = iota
	MoveTo
	Gather
	Build
)

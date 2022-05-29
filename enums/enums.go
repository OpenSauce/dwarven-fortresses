package enums

type TileTypeEnum int
type TaskTypeEnum int

const (
	Dirt TileTypeEnum = iota
	Grass
	Rock
	Water

	None TaskTypeEnum = iota
	MoveTo
	Gather
	Build
)

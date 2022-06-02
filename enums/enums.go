package enums

type TileTypeEnum int
type InputModeEnum int
type GuiActionEnum int
type ResourceTypeEnum int
type DropTypeEnum int

const (
	TileTypeEmpty TileTypeEnum = iota
	TileTypeDirt
	TileTypeGrass
	TileTypeRock
	TileTypeRockFloor
	TileTypeWater
	TileTypeStairDown
	TileTypeStairUp
)

const (
	InputModeNone InputModeEnum = iota
	InputModeBuild
	InputModeGather
	InputModeChop
	InputModeMine
)

const (
	GuiActionNone GuiActionEnum = iota
	GuiActionChop
	GuiActionStair
	GuiActionMine
)

const (
	ResourceTypeNone ResourceTypeEnum = iota
	ResourceTypeTree
)

const (
	DropTypeNone DropTypeEnum = iota
	DropTypeLog
)

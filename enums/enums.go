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
	TileTypeWater
)

const (
	InputModeNone InputModeEnum = iota
	InputModeBuild
	InputModeGather
	InputModeChop
)

const (
	GuiActionNone GuiActionEnum = iota
	GuiActionChop
	GuiActionStair
)

const (
	ResourceTypeNone ResourceTypeEnum = iota
	ResourceTypeTree
)

const (
	DropTypeNone DropTypeEnum = iota
	DropTypeLog
)

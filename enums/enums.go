package enums

type TileTypeEnum int
type InputModeEnum int
type GuiActionEnum int
type ResourceTypeEnum int
type DropTypeEnum int
type TaskTypeEnum int
type DesignationTypeEnum int
type ItemTypeEnum int

const (
	TileTypeEmpty TileTypeEnum = iota
	TileTypeDirt
	TileTypeGrass
	TileTypeRock
	TileTypeRockFloor
	TileTypeWater
	TileTypeStairDown
	TileTypeStairUp
	TileTypeStockpile
)

const (
	InputModeNone InputModeEnum = iota
	InputModeBuild
	InputModeGather
	InputModeChop
	InputModeMine
	InputModeStockpile
	InputModeHaul
)

const (
	GuiActionNone GuiActionEnum = iota
	GuiActionChop
	GuiActionStair
	GuiActionMine
	GuiActionStockpile
)

const (
	ResourceTypeNone ResourceTypeEnum = iota
	ResourceTypeTree
)

const (
	DropTypeNone DropTypeEnum = iota
	DropTypeLog
)

const (
	TaskTypeNone TaskTypeEnum = iota
	TaskTypePickUp
	TaskTypeHaul
	TaskTypeDrop
	TaskTypeChop
	TaskTypeMine
	TaskTypeBuild
	TaskTypeAddToStockpile
)

const (
	DesignationTypeNone DesignationTypeEnum = iota
	DesignationTypeStockpile
)

const (
	ItemTypeNone ItemTypeEnum = iota
	ItemTypeLog
	ItemTypeStone
)

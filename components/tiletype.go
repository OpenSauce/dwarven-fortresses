package components

import "github.com/tomknightdev/dwarven-fortresses/enums"

type TileType struct {
	enums.TileTypeEnum
}

func NewTileType(ttv enums.TileTypeEnum) TileType {
	return TileType{
		TileTypeEnum: ttv,
	}
}

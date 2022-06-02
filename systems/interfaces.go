package systems

import (
	"github.com/OpenSauce/paths"
	"github.com/tomknightdev/dwarven-fortresses/components"
	"github.com/tomknightdev/dwarven-fortresses/enums"
)

type GameMap interface {
	GetGrids() map[int]*paths.Grid
	GetTilesByZ(int) []struct {
		components.Position
		components.TileType
		components.Sprite
	}
	GetResourcesByZ(int) []struct {
		components.Position
		components.Sprite
	}
	GetTilesByType(enums.TileTypeEnum) []components.Position
	UpdateTile(enums.TileTypeEnum, int, enums.TileTypeEnum)
	GetTileByTypeIndexFromPos(enums.TileTypeEnum, components.Position) int
}

package components

import (
	"github.com/OpenSauce/paths"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tomknightdev/dwarven-fortresses/assets"
	"github.com/tomknightdev/dwarven-fortresses/enums"
)

type GameMapSingleton struct {
	WorldGenerated bool
	OffScreen      *ebiten.Image
	TilesByZ       map[int][]struct {
		Position
		TileType
		Sprite
	}
	TilesByType  map[enums.TileTypeEnum][]Position
	ResourcesByZ map[int][]struct {
		Position
		Sprite
	}
	Grids             map[int]*paths.Grid
	TilesToUpdateChan chan struct {
		FromTileType enums.TileTypeEnum
		ToTileType   enums.TileTypeEnum
		TileIndex    int
	}
}

func NewGameMapSingleton() GameMapSingleton {
	gm := GameMapSingleton{
		Grids: make(map[int]*paths.Grid),
		TilesByZ: map[int][]struct {
			Position
			TileType
			Sprite
		}{},
		TilesByType: make(map[enums.TileTypeEnum][]Position),
		ResourcesByZ: make(map[int][]struct {
			Position
			Sprite
		}),
		TilesToUpdateChan: make(chan struct {
			FromTileType enums.TileTypeEnum
			ToTileType   enums.TileTypeEnum
			TileIndex    int
		}, 100),
		OffScreen: ebiten.NewImage(assets.WorldWidth*assets.CellSize, assets.WorldHeight*assets.CellSize),
	}

	return gm
}

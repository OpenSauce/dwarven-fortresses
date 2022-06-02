package helpers

import (
	"fmt"
	"math/rand"

	"github.com/OpenSauce/paths"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sedyh/mizu/pkg/engine"
	"github.com/tomknightdev/dwarven-fortresses/assets"
	"github.com/tomknightdev/dwarven-fortresses/components"
	"github.com/tomknightdev/dwarven-fortresses/enums"
)

type GameMap struct {
	TilesByZ map[int][]struct {
		components.Position
		components.TileType
		components.Sprite
	}
	TilesByType map[enums.TileTypeEnum][]components.Position
	Grids       map[int]*paths.Grid
	World       engine.World
}

func (gm GameMap) GetGrids() map[int]*paths.Grid {
	return gm.Grids
}

func (gm GameMap) GetTilesByZ(z int) []struct {
	components.Position
	components.TileType
	components.Sprite
} {
	return gm.TilesByZ[z]
}

func (gm GameMap) GetTilesByType(tt enums.TileTypeEnum) []components.Position {
	return gm.TilesByType[tt]
}

// NewGameMap creates the world map and stores each tile information
func NewGameMap(world engine.World) GameMap {
	w := GameMap{
		Grids: make(map[int]*paths.Grid),
		TilesByZ: map[int][]struct {
			components.Position
			components.TileType
			components.Sprite
		}{},
		TilesByType: make(map[enums.TileTypeEnum][]components.Position),
		World:       world,
	}

	for z := 1; z <= assets.WorldLevels; z++ {
		g := paths.NewGrid(assets.WorldWidth, assets.WorldHeight, assets.CellSize, assets.CellSize)
		for x := 0; x < assets.WorldWidth; x++ {
			for y := 0; y < assets.WorldHeight; y++ {
				if x == 0 || x == assets.WorldWidth-1 || y == 0 || y == assets.WorldHeight-1 {
					c := g.Get(x, y)
					c.Walkable = false // There's a weird issue with pathfinding where it panics if the border cells are walkable
				}

				t := struct {
					components.Position
					components.TileType
					components.Sprite
				}{
					Position: components.NewPosition(x, y, z),
				}

				if z == 5 {
					t.TileType = components.NewTileType(enums.TileTypeDirt)
					t.Image = assets.Images["dirt0"]
				} else if z < 5 {
					t.TileType = components.NewTileType(enums.TileTypeRock)
					t.Image = assets.Images["rock"]
				} else {
					t.TileType = components.NewTileType(enums.TileTypeEmpty)
				}

				w.TilesByZ[z] = append(w.TilesByZ[z], t)
				w.TilesByType[t.TileTypeEnum] = append(w.TilesByType[t.TileTypeEnum], t.Position)
			}
		}
		w.Grids[z] = g
	}

	return w
}

func (g GameMap) UpdateTile(fromTileType enums.TileTypeEnum, tileByTypeIndex int, newTileType enums.TileTypeEnum) {
	// t.tileUpdateChan <- struct {
	// 	tileByTypeIndex int
	// 	newTileType     enums.TileTypeEnum
	// 	fromTileType    enums.TileTypeEnum
	// }{
	// 	tileByTypeIndex: tileByTypeIndex,
	// 	newTileType:     newTileType,
	// 	fromTileType:    fromTileType,
	// }

	tile := g.GetTilesByType(fromTileType)[tileByTypeIndex]
	tileMap := g.World.View(components.TileMap{}, components.Sprite{}, components.Position{}).Filter()
	for _, tm := range tileMap {
		var tmPos *components.Position
		var tmSprite *components.Sprite

		tm.Get(&tmPos, &tmSprite)
		if tmPos.Z == tile.Z {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(tile.X*assets.CellSize), float64(tile.Y*assets.CellSize))
			r := rand.Intn(3)
			tmSprite.Image.DrawImage(assets.Images[fmt.Sprintf("grass%d", r)], op)

			// Update map
			g.TilesByType[fromTileType] = append(g.TilesByType[fromTileType][:tileByTypeIndex], g.TilesByType[fromTileType][tileByTypeIndex+1:]...)
			g.TilesByType[newTileType] = append(g.TilesByType[newTileType], tile)
			break
		}
	}
}

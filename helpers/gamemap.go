package helpers

import (
	"fmt"
	"math/rand"
	"time"

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
	TilesByType  map[enums.TileTypeEnum][]components.Position
	ResourcesByZ map[int][]struct {
		components.Position
		components.Sprite
	}
	Grids map[int]*paths.Grid
	World engine.World
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

func (gm GameMap) GetResourcesByZ(z int) []struct {
	components.Position
	components.Sprite
} {
	return gm.ResourcesByZ[z]
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
		ResourcesByZ: make(map[int][]struct {
			components.Position
			components.Sprite
		}),
		World: world,
	}

	// Setup world tiles
	for z := 1; z <= assets.WorldLevels; z++ {
		g := paths.NewGrid(assets.WorldWidth, assets.WorldHeight, assets.CellSize, assets.CellSize)
		for x := 0; x < assets.WorldWidth; x++ {
			for y := 0; y < assets.WorldHeight; y++ {
				c := g.Get(x, y)
				if x == 0 || x == assets.WorldWidth-1 || y == 0 || y == assets.WorldHeight-1 {

					// c.Walkable = false // There's a weird issue with pathfinding where it panics if the border cells are walkable
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
					c.Walkable = false
				} else {
					t.TileType = components.NewTileType(enums.TileTypeEmpty)
					c.Walkable = false
				}

				w.TilesByZ[z] = append(w.TilesByZ[z], t)
				w.TilesByType[t.TileTypeEnum] = append(w.TilesByType[t.TileTypeEnum], t.Position)
			}
		}
		w.Grids[z] = g
	}

	// Setup resource tiles
	rand.Seed(time.Now().UnixNano())

	for _, tile := range w.TilesByType[enums.TileTypeDirt] {
		if rand.Intn(100) < 5 {
			g := w.Grids[tile.Z]
			c := g.Get(tile.X, tile.Y)
			c.Walkable = false

			t := struct {
				components.Position
				components.Sprite
			}{
				Position: components.NewPosition(tile.X, tile.Y, tile.Z),
				Sprite:   components.NewSprite(assets.Images["tree0"], 0),
			}

			w.ResourcesByZ[tile.Z] = append(w.ResourcesByZ[tile.Z], t)
		}
	}

	return w
}

func (g GameMap) UpdateTile(fromTileType enums.TileTypeEnum, tileByTypeIndex int, newTileType enums.TileTypeEnum) {
	tile := g.GetTilesByType(fromTileType)[tileByTypeIndex]
	tileMap := g.World.View(components.TileMap{}, components.Sprite{}, components.Position{}).Filter()
	rand.Seed(time.Now().UnixNano())

	for _, tm := range tileMap {
		var tmPos *components.Position
		var tmSprite *components.Sprite

		tm.Get(&tmPos, &tmSprite)
		if tmPos.Z == tile.Z {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(tile.X*assets.CellSize), float64(tile.Y*assets.CellSize))

			switch newTileType {
			case enums.TileTypeGrass:
				r := rand.Intn(3)
				tmSprite.Image.DrawImage(assets.Images[fmt.Sprintf("grass%d", r)], op)
			case enums.TileTypeRockFloor:
				tmSprite.Image.DrawImage(assets.Images["rockfloor"], op)
				cell := g.Grids[tile.Z].Get(tile.X, tile.Y)
				cell.Walkable = true
			}
			// Update map
			g.TilesByType[fromTileType] = append(g.TilesByType[fromTileType][:tileByTypeIndex], g.TilesByType[fromTileType][tileByTypeIndex+1:]...)
			g.TilesByType[newTileType] = append(g.TilesByType[newTileType], tile)
			break
		}
	}
}

func (g GameMap) GetTileByTypeIndexFromPos(tt enums.TileTypeEnum, pos components.Position) (int, error) {
	for i, t := range g.TilesByType[tt] {
		if t.X == pos.X && t.Y == pos.Y && t.Z == pos.Z {
			return i, nil
		}
	}

	return 0, fmt.Errorf("unable to find %v at %v", tt, pos)
}

func (g GameMap) AddTileByType(tileType enums.TileTypeEnum, pos components.Position) {
	g.TilesByType[tileType] = append(g.TilesByType[tileType], pos)
}

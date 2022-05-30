package scenes

import (
	"github.com/OpenSauce/paths"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sedyh/mizu/pkg/engine"
	"github.com/tomknightdev/dwarven-fortresses/assets"
	"github.com/tomknightdev/dwarven-fortresses/components"
	"github.com/tomknightdev/dwarven-fortresses/entities"
	"github.com/tomknightdev/dwarven-fortresses/enums"
	"github.com/tomknightdev/dwarven-fortresses/systems"
)

type Game struct {
	grids map[int]*paths.Grid
}

func (g *Game) Setup(w engine.World) {
	w.AddComponents(
		components.GameMap{},
		components.Position{},
		components.Sprite{},
		components.Move{},
		components.Input{},
		components.Zoom{},
		components.TileType{},
		components.Task{},
		components.Worker{},
		components.TileMap{},
	)

	gameMap := components.NewGameMap(assets.WorldWidth, assets.WorldHeight, assets.WorldLevels, assets.CellSize)

	w.AddSystems(
		systems.NewRender(assets.WorldWidth, assets.WorldHeight, assets.CellSize, nil),
		systems.NewPathfinder(gameMap.Grids),
		systems.NewInput(),
		systems.NewActor(),
		systems.NewNature(),
	)

	// World
	for z := 0; z < assets.WorldLevels; z++ {
		tmImage := ebiten.NewImage(assets.WorldWidth*assets.CellSize, assets.WorldHeight*assets.CellSize)
		for _, t := range gameMap.Tiles[z] {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(t.X*assets.CellSize), float64(t.Y*assets.CellSize))
			if z == 5 {
				tmImage.DrawImage(assets.Images["dirt0"], op)
			} else if z < 5 {
				tmImage.DrawImage(assets.Images["rock"], op)

			}

			w.AddEntities(&entities.Tile{
				Position: components.NewPosition(t.X, t.Y, t.Z),
				// Sprite:   components.NewSprite(assets.Images["grass0"]),
				TileType: components.NewTileType(enums.Dirt),
			})
		}
		w.AddEntities(&entities.TileMap{
			Sprite:   components.NewSprite(tmImage),
			Position: components.NewPosition(0, 0, z),
			TileMap:  components.NewTileMap(),
		})
	}

	// Actors
	for i := 0; i < assets.DwarfCount; i++ {
		w.AddEntities(&entities.Actor{
			Position: components.NewPosition(1, 1, 5),
			Sprite:   components.NewSprite(assets.Images["dwarf"]),
			Move:     components.NewMove(1, 1, 5),
			Worker:   components.NewWorker(),
		})
	}

	// Input
	cx, cy := ebiten.CursorPosition()
	w.AddEntities(&entities.Input{
		MousePos:    components.NewPosition(cx, cy, 5),
		CursorImage: components.NewSprite(assets.Images["cursor"]),
		Input:       components.NewInput(),
	})

	// Camera
	w.AddEntities(&entities.Camera{
		Zoom:     components.NewZoom(),
		Position: components.NewPosition(0, 0, 5),
	})
}

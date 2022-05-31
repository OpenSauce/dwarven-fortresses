package scenes

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sedyh/mizu/pkg/engine"
	"github.com/tomknightdev/dwarven-fortresses/assets"
	"github.com/tomknightdev/dwarven-fortresses/components"
	"github.com/tomknightdev/dwarven-fortresses/entities"
	"github.com/tomknightdev/dwarven-fortresses/enums"
	"github.com/tomknightdev/dwarven-fortresses/systems"
)

type Game struct {
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
		components.Gui{},
	)

	gameMap := components.NewGameMap(assets.WorldWidth, assets.WorldHeight, assets.WorldLevels, assets.CellSize)

	w.AddSystems(
		systems.NewRender(assets.WorldWidth, assets.WorldHeight, assets.CellSize, nil),
		systems.NewPathfinder(gameMap.Grids),
		systems.NewInput(),
		systems.NewActor(),
		systems.NewNature(),
		systems.NewGui(),
	)

	setupWorld(w, gameMap)

	// Actors
	for i := 0; i < assets.StartingDwarfCount; i++ {
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
		CursorImage: components.NewSprite(assets.Images["empty"]),
		Input:       components.NewInput(),
	})

	// Camera
	w.AddEntities(&entities.Camera{
		Zoom:     components.NewZoom(),
		Position: components.NewPosition(0, 0, 5),
	})

	setupGui(w)
}

func setupWorld(w engine.World, gameMap components.GameMap) {
	for z := 0; z < assets.WorldLevels; z++ {
		tmImage := ebiten.NewImage(assets.WorldWidth*assets.CellSize, assets.WorldHeight*assets.CellSize)
		for _, t := range gameMap.Tiles[z] {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(t.X*assets.CellSize), float64(t.Y*assets.CellSize))
			if z == 5 {
				tmImage.DrawImage(assets.Images["dirt0"], op)
				w.AddEntities(&entities.Tile{
					Position: components.NewPosition(t.X, t.Y, t.Z),
					TileType: components.NewTileType(enums.TileTypeDirt),
				})
			} else if z < 5 {
				tmImage.DrawImage(assets.Images["rock"], op)
				w.AddEntities(&entities.Tile{
					Position: components.NewPosition(t.X, t.Y, t.Z),
					TileType: components.NewTileType(enums.TileTypeRock),
				})
			} else {
				w.AddEntities(&entities.Tile{
					Position: components.NewPosition(t.X, t.Y, t.Z),
					TileType: components.NewTileType(enums.TileTypeEmpty),
				})
			}

		}
		w.AddEntities(&entities.TileMap{
			Sprite:   components.NewSprite(tmImage),
			Position: components.NewPosition(0, 0, z),
			TileMap:  components.NewTileMap(),
		})
	}
}

func setupGui(w engine.World) {
	w.AddEntities(&entities.Gui{
		Gui:    components.NewGui(10, 200, 3.0),
		Sprite: components.NewSprite(assets.Images["stairdown"]),
	})
}

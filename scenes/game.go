package scenes

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sedyh/mizu/pkg/engine"
	"github.com/tomknightdev/dwarven-fortresses/assets"
	"github.com/tomknightdev/dwarven-fortresses/components"
	"github.com/tomknightdev/dwarven-fortresses/entities"
	"github.com/tomknightdev/dwarven-fortresses/enums"
	"github.com/tomknightdev/dwarven-fortresses/helpers"
	"github.com/tomknightdev/dwarven-fortresses/systems"
)

type Game struct {
	gameMap systems.GameMap
}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) Setup(w engine.World) {

	g.gameMap = helpers.NewGameMap(w)

	w.AddComponents(
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

	w.AddSystems(
		systems.NewRender(assets.WorldWidth, assets.WorldHeight, assets.CellSize, nil),
		systems.NewPathfinder(g.gameMap.GetGrids()),
		systems.NewInput(),
		systems.NewActor(),
		systems.NewNature(g.gameMap),
		systems.NewGui(),
		systems.NewTileMap(),
	)

	setupWorld(w, g.gameMap)

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

func setupWorld(w engine.World, gameMap systems.GameMap) {
	for z := 0; z < assets.WorldLevels; z++ {
		tmImage := ebiten.NewImage(assets.WorldWidth*assets.CellSize, assets.WorldHeight*assets.CellSize)
		for _, t := range gameMap.GetTilesByZ(z) {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(t.X*assets.CellSize), float64(t.Y*assets.CellSize))

			if t.TileTypeEnum != enums.TileTypeEmpty {
				tmImage.DrawImage(t.Image, op)
			}
			// w.AddEntities(&entities.Tile{
			// 	Position: components.NewPosition(t.X, t.Y, t.Z),
			// 	TileType: components.NewTileType(t.TileTypeEnum),
			// })
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

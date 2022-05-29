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
		components.WorldMap{},
		components.Position{},
		components.Sprite{},
		components.Move{},
		components.Input{},
		components.Zoom{},
		components.TileType{},
		components.Task{},
		components.Worker{},
	)

	world := components.NewWorldMap(assets.WorldWidth, assets.WorldHeight, assets.WorldLevels, assets.CellSize)

	w.AddSystems(
		systems.NewRender(assets.WorldWidth, assets.WorldHeight, assets.CellSize),
		systems.NewPathfinder(world.Grids),
		systems.NewInput(),
		systems.NewActor())

	// World
	for _, t := range world.Tiles {
		w.AddEntities(&entities.Tile{
			Position: components.NewPosition(t.X, t.Y, t.Z),
			Sprite:   components.NewSprite(assets.Images["grass0"]),
			TileType: components.NewTileType(enums.Grass),
		})
	}

	// Actors
	for i := 0; i < assets.DwarfCount; i++ {
		w.AddEntities(&entities.Actor{
			Position: components.NewPosition(1, 1, 0),
			Sprite:   components.NewSprite(assets.Images["dwarf"]),
			Move:     components.NewMove(1, 1, 0),
			Worker:   components.NewWorker(),
		})
	}

	// Input
	cx, cy := ebiten.CursorPosition()
	w.AddEntities(&entities.Input{
		MousePos:    components.NewPosition(cx, cy, 0),
		CursorImage: components.NewSprite(assets.Images["cursor"]),
		Input:       components.NewInput(),
	})

	// Camera
	w.AddEntities(&entities.Camera{
		Zoom:     components.NewZoom(),
		Position: components.NewPosition(0, 0, 0),
	})
}

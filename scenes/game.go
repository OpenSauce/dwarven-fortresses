package scenes

import (
	"github.com/OpenSauce/paths"
	"github.com/sedyh/mizu/pkg/engine"
	"github.com/tomknightdev/dwarven-fortresses/assets"
	"github.com/tomknightdev/dwarven-fortresses/components"
	"github.com/tomknightdev/dwarven-fortresses/entities"
)

type Game struct {
	grids map[int]*paths.Grid
}

func (g *Game) Setup(w engine.World) {
	w.AddComponents(
		components.Construct{},
		components.Position{},
		components.Sprite{},
	)

	world := components.NewConstruct(assets.WorldWidth, assets.WorldHeight, assets.WorldLevels)

	w.AddSystems()

	for z := 0; z < world.Levels; z++ {
		for x := 0; x < world.Width; x++ {
			for y := 0; y < world.Height; y++ {
				w.AddEntities(&entities.Tile{
					Position: components.NewPosition(x, y, z),
					Sprite:   components.NewSprite(assets.Images["grass0"]),
				})
			}
		}
	}
}

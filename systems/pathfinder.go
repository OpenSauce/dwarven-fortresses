package systems

import (
	"github.com/OpenSauce/paths"
	"github.com/sedyh/mizu/pkg/engine"
	"github.com/tomknightdev/dwarven-fortresses/assets"
	"github.com/tomknightdev/dwarven-fortresses/components"
)

type Pathfinder struct {
	grids map[int]*paths.Grid
}

func NewPathfinder(grids map[int]*paths.Grid) *Pathfinder {
	return &Pathfinder{
		grids: grids,
	}
}

func (p *Pathfinder) Update(w engine.World) {
	view := w.View(components.Move{}, components.Position{})
	view.Each((func(e engine.Entity) {
		var move *components.Move
		var pos *components.Position
		e.Get(&move, &pos)

		if pos.X != move.X || pos.Y != move.Y || pos.Z != move.Z {
			move.Arrived = false
			path := p.grids[pos.Z].GetPath(float64(pos.X*assets.CellSize), float64(pos.Y*assets.CellSize), float64(move.X*assets.CellSize), float64(move.Y*assets.CellSize), true, true)

			c := path.Next()
			pos.X = c.X
			pos.Y = c.Y
			path.Advance()
		} else {
			move.Arrived = true
		}
	}))
}

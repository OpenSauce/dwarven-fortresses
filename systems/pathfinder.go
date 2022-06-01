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
	p := &Pathfinder{
		grids: grids,
	}

	return p
}

func (p *Pathfinder) Update(w engine.World) {
	view := w.View(components.Move{}, components.Position{})
	view.Each((func(e engine.Entity) {
		var move *components.Move
		var pos *components.Position
		e.Get(&move, &pos)

		if move.CurrentEnergy < move.TotalEnergy {
			move.CurrentEnergy++
			return
		}

		if move.GettingRoute {
			return
		}

		if pos.X != move.X || pos.Y != move.Y || pos.Z != move.Z {
			move.Arrived = false

			if move.CurrentPath == nil {
				move.GettingRoute = true
				path := p.grids[pos.Z].GetPath(float64(pos.X*assets.CellSize), float64(pos.Y*assets.CellSize), float64(move.X*assets.CellSize), float64(move.Y*assets.CellSize), true, true)

				if path == nil {
					move.Arrived = true
					move.GettingRoute = false

					return
				}
				move.CurrentPath = path
			}

			c := move.CurrentPath.Next()
			pos.X = c.X
			pos.Y = c.Y
			move.CurrentPath.Advance()
			move.CurrentEnergy = 0
			move.GettingRoute = false

			for _, c := range move.CurrentPath.Cells[move.CurrentPath.CurrentIndex:] {
				if !c.Walkable {
					move.CurrentPath = nil
					break
				}
			}
		} else {
			move.CurrentPath = nil
			move.Arrived = true
		}
	}))
}

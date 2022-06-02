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

		if (move.Adjacent && IsAdjacent(*move, *pos)) || Matches(*move, *pos) {
			move.CurrentPath = nil
			move.Arrived = true
		} else {
			move.Arrived = false

			if move.CurrentPath == nil {
				move.GettingRoute = true

				var path *paths.Path

				if move.Adjacent {
					adjacents := GetAdjacents(*move)

					for _, a := range adjacents {
						path = p.grids[pos.Z].GetPath(float64(pos.X*assets.CellSize), float64(pos.Y*assets.CellSize), float64(a.X*assets.CellSize), float64(a.Y*assets.CellSize), true, true)
						if path != nil {
							break
						}
					}
				} else {
					path = p.grids[pos.Z].GetPath(float64(pos.X*assets.CellSize), float64(pos.Y*assets.CellSize), float64(move.X*assets.CellSize), float64(move.Y*assets.CellSize), true, true)
				}

				if path == nil {
					// move.Arrived = true
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
		}
	}))
}

func IsAdjacent(dest components.Move, current components.Position) bool {
	if current.X >= dest.X-1 && current.X <= dest.X+1 && current.Y >= dest.Y-1 && current.Y <= dest.Y+1 {
		return true
	}

	return false
}

func GetAdjacents(dest components.Move) []components.Position {
	var adjacents []components.Position

	for x := -1; x < 2; x++ {
		for y := -1; y < 2; y++ {
			if x == dest.X && y == dest.Y {
				continue
			}
			adjacents = append(adjacents, components.NewPosition(dest.X+x, dest.Y+y, dest.Z))
		}
	}

	return adjacents
}

func Matches(a components.Move, b components.Position) bool {
	if a.X == b.X && a.Y == b.Y && a.Z == b.Z {
		return true
	}

	return false
}

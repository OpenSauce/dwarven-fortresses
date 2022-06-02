package systems

import (
	"github.com/OpenSauce/paths"
	"github.com/sedyh/mizu/pkg/engine"
	"github.com/tomknightdev/dwarven-fortresses/assets"
	"github.com/tomknightdev/dwarven-fortresses/components"
)

type Pathfinder struct {
	grids   map[int]*paths.Grid
	GameMap GameMap
}

func NewPathfinder(grids map[int]*paths.Grid, gameMap GameMap) *Pathfinder {
	p := &Pathfinder{
		grids:   grids,
		GameMap: gameMap,
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
			if len(move.CurrentPaths) > 1 {
				move.CurrentPaths = move.CurrentPaths[1:]
				pos.Z = move.CurrentPaths[0].Level

			} else {
				move.CurrentPaths = nil
				move.Arrived = true
			}
		} else {
			move.Arrived = false

			if move.CurrentPaths == nil {
				move.GettingRoute = true

				var paths []struct {
					*paths.Path
					Level int
				}

				if move.Adjacent {
					adjacents := GetAdjacents(*move)

					for _, a := range adjacents {
						paths = p.GetPath(*pos, a)
						if len(paths) > 0 {
							break
						}
						// path = p.grids[pos.Z].GetPath(float64(pos.X*assets.CellSize), float64(pos.Y*assets.CellSize), float64(a.X*assets.CellSize), float64(a.Y*assets.CellSize), true, true)
						// if path != nil {
						// 	break
						// }
					}
				} else {
					paths = p.GetPath(*pos, components.NewPosition(move.X, move.Y, move.Z))
					// path = p.grids[pos.Z].GetPath(float64(pos.X*assets.CellSize), float64(pos.Y*assets.CellSize), float64(move.X*assets.CellSize), float64(move.Y*assets.CellSize), true, true)
				}

				if len(paths) == 0 {
					// move.Arrived = true
					move.GettingRoute = false

					return
				}
				move.CurrentPaths = paths
			}

			c := move.CurrentPaths[0].Next()
			pos.X = c.X
			pos.Y = c.Y
			move.CurrentPaths[0].Advance()
			move.CurrentEnergy = 0
			move.GettingRoute = false

			for _, c := range move.CurrentPaths[0].Cells[move.CurrentPaths[0].CurrentIndex:] {
				if !c.Walkable {
					move.CurrentPaths = nil
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

func (p Pathfinder) GetPath(startPos components.Position, endPos components.Position) []struct {
	*paths.Path
	Level int
} {
	if startPos.Z == endPos.Z {
		path := p.grids[endPos.Z].GetPath(float64(startPos.X)*assets.CellSize, float64(startPos.Y)*assets.CellSize, float64(endPos.X)*assets.CellSize, float64(endPos.Y)*assets.CellSize, true, true)

		if path != nil {
			return []struct {
				*paths.Path
				Level int
			}{
				{
					path,
					endPos.Z,
				},
			}
		}
	}

	// stairsUp := p.GameMap.GetTilesByType(enums.TileTypeStairUp)
	// stairsDown := p.GameMap.GetTilesByType(enums.TileTypeStairDown)

	// // Routes is a map of total cost and all paths between start and end positions
	// routes := make(map[int][]*paths.Path)

	// // This will define whether we go down or up in the search
	// startZ := startPos.Z
	// endZ := endPos.Z
	// reverse := false

	// if endZ > startZ {
	// 	startZ = startZ + endZ
	// 	endZ = startZ - endZ
	// 	startZ = startZ - endZ
	// 	reverse = true
	// }

	// nextStartPos := startPos
	// nextEndPos := endPos

	// var currentLevelStairsUp []*paths.Path
	// var currentLevelStairsDown []*paths.Path

	// // Work backwards from end pos and then reverse if needed
	// for z := endZ; z <= startZ; z-- {
	// 	// Get all stairs that are on same level and can reach end pos
	// 	currentLevelStairsUp, currentLevelStairsDown = p.navigateLevel(z, nextEndPos.X, nextEndPos.Y, stairsUp, stairsDown)

	// 	if z == startZ {
	// 		if len(currentLevelStairsUp) > 0 {

	// 		}
	// 	}
	// }

	return nil
}

func (p *Pathfinder) navigateLevel(x, y, z int, stairsUp, stairsDown []components.Position) ([]*paths.Path, []*paths.Path) {
	// Get all stairs that are on same level and can reach end pos
	var currentLevelStairsUp []*paths.Path
	for _, s := range stairsUp {
		if s.Z == z {
			path := p.grids[z].GetPath(float64(s.X)*assets.CellSize, float64(s.Y)*assets.CellSize, float64(x)*assets.CellSize, float64(y)*assets.CellSize, true, true)

			if path != nil {
				currentLevelStairsUp = append(currentLevelStairsUp, path)
			}
		}
	}
	var currentLevelStairsDown []*paths.Path
	for _, s := range stairsDown {
		if s.Z == z {
			path := p.grids[z].GetPath(float64(s.X)*assets.CellSize, float64(s.Y)*assets.CellSize, float64(x)*assets.CellSize, float64(y)*assets.CellSize, true, true)

			if path != nil {
				currentLevelStairsDown = append(currentLevelStairsDown, path)
			}
		}
	}

	return currentLevelStairsUp, currentLevelStairsDown
}

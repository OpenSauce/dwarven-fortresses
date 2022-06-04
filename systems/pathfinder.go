package systems

import (
	"log"

	"github.com/OpenSauce/paths"
	"github.com/sedyh/mizu/pkg/engine"
	"github.com/tomknightdev/dwarven-fortresses/assets"
	"github.com/tomknightdev/dwarven-fortresses/components"
	"github.com/tomknightdev/dwarven-fortresses/entities"
	"github.com/tomknightdev/dwarven-fortresses/enums"
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

		if (move.Adjacent && IsAdjacent(*move, *pos)) || p.GameMap.Matches(components.NewPosition(move.X, move.Y, move.Z), *pos) {
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

				var paths []components.Path

				if move.Adjacent {
					adjacents := p.GetAdjacents(*move)

					for _, a := range adjacents {
						paths = p.GetPath(*pos, a)
						if len(paths) > 0 {
							if len(paths[0].Cells) == 0 {
								log.Println("path with no cells found")
								continue
							}

							break
						}
					}

					// Path to job not found
					if len(paths) == 0 {
						var wk *components.Worker
						e.Get(&wk)
						wk.HasJob = false

						job, found := w.GetEntity(wk.JobID)
						if found {
							var j *components.Task
							var jPos *components.Position
							job.Get(&j, &jPos)
							j.Claimed = false

							w.AddEntities(&entities.Job{
								Position: *jPos,
								Task:     *j,
							})

							w.RemoveEntity(job)
						}
					}

				} else {
					paths = p.GetPath(*pos, components.NewPosition(move.X, move.Y, move.Z))
				}

				if len(paths) == 0 {
					// move.Arrived = true
					move.GettingRoute = false

					return
				}
				move.CurrentPaths = paths
			}

			if move.CurrentPaths[0].AtEnd() {
				if len(move.CurrentPaths) > 1 {
					move.CurrentPaths = move.CurrentPaths[1:]
					pos.Z = move.CurrentPaths[0].Level

				} else {
					move.CurrentPaths = nil
					move.Arrived = true
				}
			} else {
				c := move.CurrentPaths[0].Next()

				if c == nil {
					panic("no next cell in path")
				}

				pos.X = c.X
				pos.Y = c.Y
				move.CurrentPaths[0].Advance()

				for _, c := range move.CurrentPaths[0].Cells[move.CurrentPaths[0].CurrentIndex:] {
					if !c.Walkable {
						move.CurrentPaths = nil
						break
					}
				}
			}

			move.CurrentEnergy = 0
			move.GettingRoute = false
		}
	}))
}

func IsAdjacent(dest components.Move, current components.Position) bool {
	if current.X >= dest.X-1 && current.X <= dest.X+1 && current.Y >= dest.Y-1 && current.Y <= dest.Y+1 && current.Z == dest.Z {
		return true
	}

	return false
}

func (p *Pathfinder) GetAdjacents(dest components.Move) []components.Position {
	var adjacents []components.Position

	for x := -1; x < 2; x++ {
		for y := -1; y < 2; y++ {
			if (x == 0 && y == 0) || dest.X+x < 0 || dest.Y+y < 0 || dest.X+x > assets.WorldWidth-1 || dest.Y+y > assets.WorldHeight-1 {
				continue
			}

			if p.grids[dest.Z].Get(dest.X+x, dest.Y+y).Walkable {
				adjacents = append(adjacents, components.NewPosition(dest.X+x, dest.Y+y, dest.Z))
			}
		}
	}

	return adjacents
}

func (p Pathfinder) GetPath(startPos components.Position, endPos components.Position) []components.Path {
	if startPos.Z == endPos.Z {
		path := p.grids[endPos.Z].GetPath(float64(startPos.X)*assets.CellSize, float64(startPos.Y)*assets.CellSize, float64(endPos.X)*assets.CellSize, float64(endPos.Y)*assets.CellSize, true, false)

		if path != nil && len(path.Cells) > 0 {
			return []components.Path{
				{
					Path:  path,
					Level: endPos.Z,
				},
			}
		}
	}

	// Use golden path approach, assume that the dwarf can reach the end pos from any stair
	// true = down, false = up
	travTiles := make(map[bool][]components.Position)

	travTiles[false] = append(travTiles[false], p.GameMap.GetTilesByType(enums.TileTypeStairUp)...)
	travTiles[true] = append(travTiles[true], p.GameMap.GetTilesByType(enums.TileTypeStairDown)...)

	direction := true
	if endPos.Z > startPos.Z {
		direction = false
	}

	var paths []components.Path
	var reached bool
	// Find route to each stair from current location, checking if each can get to final destination
	for _, tt := range travTiles[direction] {
		path := p.grids[startPos.Z].GetPath(float64(startPos.X)*assets.CellSize, float64(startPos.Y)*assets.CellSize, float64(tt.X)*assets.CellSize, float64(tt.Y)*assets.CellSize, true, false)

		if path == nil {
			continue
		}

		paths = append(paths, components.Path{
			Path:  path,
			Level: startPos.Z,
		})

		zChange := 1
		if direction {
			zChange = -1
		}

		paths, reached = p.traverseLevel(paths, travTiles[direction], direction, components.NewPosition(tt.X, tt.Y, tt.Z+zChange), endPos)
		if reached {
			return paths
		}
	}

	if !reached {
		panic("unable to reach final destination")
	}
	return paths
}

func (p Pathfinder) traverseLevel(paths []components.Path, travTiles []components.Position, direction bool, startPos, finalDest components.Position) ([]components.Path, bool) {
	// First thing to try is if we can reach final destination
	if startPos.Z == finalDest.Z {
		path := p.grids[finalDest.Z].GetPath(float64(startPos.X)*assets.CellSize, float64(startPos.Y)*assets.CellSize, float64(finalDest.X)*assets.CellSize, float64(finalDest.Y)*assets.CellSize, true, false)

		if path != nil {
			paths = append(paths, components.Path{
				Path:  path,
				Level: startPos.Z,
			})
			return paths, true
		}
	}

	for _, tt := range travTiles {
		if tt.Z == startPos.Z {
			path := p.grids[startPos.Z].GetPath(float64(startPos.X)*assets.CellSize, float64(startPos.Y)*assets.CellSize, float64(tt.X)*assets.CellSize, float64(tt.Y)*assets.CellSize, true, false)

			if path != nil {
				paths = append(paths, components.Path{
					Path:  path,
					Level: tt.Z,
				})

				// Final destination not found, next start pos will be the opposite stair in the same location
				zChange := 1
				if direction {
					zChange = -1
				}

				paths, reachedFinalDest := p.traverseLevel(paths, travTiles, direction, components.NewPosition(tt.X, tt.Y, tt.Z+zChange), finalDest)
				if reachedFinalDest {
					return paths, reachedFinalDest
				}
			}
		}
	}

	return nil, false
}

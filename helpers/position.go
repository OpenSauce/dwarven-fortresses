package helpers

import (
	"github.com/OpenSauce/paths"
	"github.com/tomknightdev/dwarven-fortresses/assets"
	"github.com/tomknightdev/dwarven-fortresses/components"
)

func Matches(a components.Position, b components.Position) bool {
	if a.X == b.X && a.Y == b.Y && a.Z == b.Z {
		return true
	}

	return false
}

func GetAdjacents(grid *paths.Grid, pos components.Position, walkable bool) []components.Position {
	var adjacents []components.Position

	for x := -1; x < 2; x++ {
		for y := -1; y < 2; y++ {
			if (x == 0 && y == 0) || pos.X+x < 0 || pos.Y+y < 0 || pos.X+x > assets.WorldWidth-1 || pos.Y+y > assets.WorldHeight-1 {
				continue
			}

			if grid.Get(pos.X+x, pos.Y+y).Walkable == walkable {
				adjacents = append(adjacents, components.NewPosition(pos.X+x, pos.Y+y, pos.Z))
			}
		}
	}

	return adjacents
}

func IsAdjacent(dest components.Move, current components.Position) bool {
	if current.X >= dest.X-1 && current.X <= dest.X+1 && current.Y >= dest.Y-1 && current.Y <= dest.Y+1 && current.Z == dest.Z {
		return true
	}

	return false
}

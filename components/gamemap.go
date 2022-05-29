package components

import (
	"github.com/OpenSauce/paths"
)

type GameMap struct {
	Tiles map[int][]struct {
		X, Y, Z int
	}
	Grids map[int]*paths.Grid
}

// NewGameMap creates the world map and stores each tile information
func NewGameMap(width, height, levels, cellSize int) GameMap {
	w := GameMap{
		Grids: make(map[int]*paths.Grid),
		Tiles: map[int][]struct {
			X int
			Y int
			Z int
		}{},
	}

	for z := 1; z <= levels; z++ {
		g := paths.NewGrid(width, height, cellSize, cellSize)
		for x := 0; x < width; x++ {
			for y := 0; y < height; y++ {
				if x == 0 || x == width-1 || y == 0 || y == height-1 {
					c := g.Get(x, y)
					c.Walkable = false
				}

				w.Tiles[z] = append(w.Tiles[z], struct {
					X, Y, Z int
				}{
					x,
					y,
					z,
				})
			}
		}
		w.Grids[z] = g
	}

	return w
}

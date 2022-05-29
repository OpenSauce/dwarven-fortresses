package components

import (
	"github.com/OpenSauce/paths"
)

type WorldMap struct {
	Tiles []struct {
		X, Y, Z int
	}
	Grids map[int]*paths.Grid
}

// NewWorldMap creates the world map and stores each tile information
func NewWorldMap(width, height, levels, cellSize int) WorldMap {
	w := WorldMap{
		Grids: make(map[int]*paths.Grid),
	}

	for z := 0; z < levels; z++ {
		g := paths.NewGrid(width, height, cellSize, cellSize)
		for x := 0; x < width; x++ {
			if x == width - 1 {
				
			}

			for y := 0; y < height; y++ {
				w.Tiles = append(w.Tiles, struct {
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

package main

import (
	"github.com/solarlune/paths"
)

type GameMap struct {
	grid *paths.Grid
}

func NewGameMap(gridWidth, gridHeight, cellWidth, cellHeight int) *GameMap {
	m := GameMap{
		grid: paths.NewGrid(gridWidth, gridHeight, cellWidth, cellHeight),
	}

	for _, c := range m.grid.AllCells() {
		if c.X == 0 || c.Y == 0 || c.X == gridWidth-1 || c.Y == gridHeight-1 {
			c.Walkable = false
		}
	}

	return &m
}

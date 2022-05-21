package main

import (
	"math/rand"

	"github.com/solarlune/paths"
)

type GameMap struct {
	grid        *paths.Grid
	getPathChan chan getPathRequest
}

type getPathRequest struct {
	startX, startY, endX, endY int
	responseChan               chan *paths.Path
}

func NewGameMap(gridWidth, gridHeight, cellWidth, cellHeight int) *GameMap {
	gm := GameMap{
		grid:        paths.NewGrid(gridWidth, gridHeight, cellWidth, cellHeight),
		getPathChan: make(chan getPathRequest),
	}

	// mapCentre := gridWidth / 2

	for _, c := range gm.grid.AllCells() {
		if c.X == 0 || c.Y == 0 || c.X == gridWidth-1 || c.Y == gridHeight-1 {
			c.Walkable = false
		} else {
			r := rand.Intn(100)
			if r > 90 {
				c.Walkable = false
			}
			// Take the map centre, deduct current tile from it, the lower the number the closer to the centre
			// xCost := math.Abs(float64(mapCentre - c.X))
			// yCost := math.Abs(float64(mapCentre - c.Y))
			// cost := float64(gridWidth) - float64(xCost+yCost)
			// c.Cost = math.Max(cost, 1)
		}
	}

	go gm.handleGetPathRequests()

	return &gm
}

func (g *GameMap) handleGetPathRequests() {
	for r := range g.getPathChan {
		r.responseChan <- g.grid.GetPathFromCells(g.grid.Get(r.startX, r.startY), g.grid.Get(r.endX, r.endY), true, true)
	}
}

func (g *GameMap) GetPath(startX, startY, endX, endY int) *paths.Path {
	responseChan := make(chan *paths.Path)
	defer close(responseChan)
	g.getPathChan <- getPathRequest{
		startX, startY, endX, endY, responseChan,
	}
	return <-responseChan
}

func (g *GameMap) GetMapDimensions() (int, int) {
	return g.grid.Width(), g.grid.Height()
}

func (g *GameMap) IsWalkable(x, y int) bool {
	return g.grid.Get(x, y).Walkable
}

func (g *GameMap) GetCellCost(x, y int) int {
	return int(g.grid.Get(x, y).Cost)
}

package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tomknightdev/paths"
	"golang.org/x/exp/rand"
)

type GameMap struct {
	grid        *paths.Grid
	tiles       map[*paths.Cell]*Tile
	getPathChan chan getPathRequest
}

type Tile struct {
	cell     *paths.Cell
	resource *Resource
	drawn    bool
	// XPos, YPos int
}

func (t *Tile) Update() error {
	if t.resource.resourceType == Dirt {
		r := rand.Intn(5000000)

		if r > 4999998 {
			t.resource = nil
			t.resource = CreateResource(2)
		} else if r > 4999995 {
			t.resource = nil
			t.resource = CreateResource(1)
		}
	}

	return nil
}

func (t *Tile) Gethered() {
	t.resource = nil
	t.resource = CreateResource(0)
}

type getPathRequest struct {
	startX, startY, endX, endY int
	responseChan               chan *paths.Path
}

func NewGameMap(gridWidth, gridHeight, cellWidth, cellHeight int) *GameMap {
	gm := GameMap{
		grid:        paths.NewGrid(gridWidth, gridHeight, cellWidth, cellHeight),
		tiles:       make(map[*paths.Cell]*Tile),
		getPathChan: make(chan getPathRequest),
	}

	for _, c := range gm.grid.AllCells() {
		r := rand.Intn(50)
		if r < 2 {
			r = 3
		} else if r < 10 {
			r = 2
		} else if r < 25 {
			r = 1
		} else {
			r = 0
		}

		c.Cost += float64(r)
		t := Tile{
			cell:     c,
			resource: CreateResource(ResourceType(r)),
		}
		gm.tiles[c] = &t

		if c.X == 0 || c.Y == 0 || c.X == gridWidth-1 || c.Y == gridHeight-1 {
			c.Walkable = false
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

func (g *GameMap) SwitchWalkable(x, y int) {
	c := g.grid.Get(x, y)
	c.Walkable = !c.Walkable
}

func (g *GameMap) Update() error {
	return nil
}

func (g *GameMap) Draw(screen *ebiten.Image) {
	camXPos := int(Cam.X) / cellWidth
	camYPos := int(Cam.Y) / cellHeight

	camWidth := Cam.Width / 2 / 8 / int(Cam.Scale)
	camHeight := Cam.Height / 2 / 8 / int(Cam.Scale)

	for c, t := range g.tiles {
		if c.X < camXPos-camWidth || c.X > camXPos+camWidth || c.Y < camYPos-camHeight || c.Y > camYPos+camHeight {
			t.drawn = false
			continue
		}

		t.drawn = true

		// Draw the tile
		op := Cam.GetTranslation(float64(c.X*cellWidth), float64(c.Y*cellHeight))

		Cam.Surface.DrawImage(t.resource.image, op)
		if t.resource.queued {
			Cam.Surface.DrawImage(cursorImage, op)
		}

	}
}

func (g *GameMap) DrawnTileCount() int {
	count := 0
	for _, t := range g.tiles {
		if t.drawn {
			count++
		}
	}
	return count
}

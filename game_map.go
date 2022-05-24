package main

import (
	"image/color"

	"github.com/OpenSauce/paths"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/exp/rand"
)

type GameMap struct {
	grids         map[int]*paths.Grid
	tiles         map[*paths.Cell]*Tile
	tilesByZLevel map[int][]*Tile
	getPathChan   chan getPathRequest
}

type Tile struct {
	cell     *paths.Cell
	resource *Resource
	drawn    bool
	zLevel   int
	// XPos, YPos int
}

func (t *Tile) Update() error {
	// if t.resource.resourceType == Dirt {
	// 	r := rand.Intn(5000000)

	// 	if r > 4999998 {
	// 		t.resource = nil
	// 		t.resource = CreateResource(2)
	// 	} else if r > 4999995 {
	// 		t.resource = nil
	// 		t.resource = CreateResource(1)
	// 	}
	// }

	return nil
}

func (t *Tile) Gethered() {
	t.resource = nil
	t.resource = CreateResource(0)
}

type getPathRequest struct {
	startX, startY, startZ, endX, endY, endZ int
	responseChan                             chan *paths.Path
}

func NewGameMap(gridWidth, gridHeight, cellWidth, cellHeight int) *GameMap {
	gm := GameMap{
		grids:         make(map[int]*paths.Grid),
		tiles:         make(map[*paths.Cell]*Tile),
		tilesByZLevel: make(map[int][]*Tile),
		getPathChan:   make(chan getPathRequest),
	}

	for i := 0; i < 10; i++ {
		gm.grids[i] = paths.NewGrid(gridWidth, gridHeight, cellWidth, cellHeight)

		for _, c := range gm.grids[i].AllCells() {
			resourceType := Empty

			if i < 5 {
				resourceType = Rock
			} else if i == 5 {
				r := rand.Intn(50)
				if r < 2 {
					resourceType = Water
				} else if r < 10 {
					resourceType = Tree
				} else if r < 25 {
					resourceType = Grass
				} else {
					resourceType = Dirt
					r = 0
				}

				// c.Cost += float64(r)
			}

			t := Tile{
				cell:     c,
				resource: CreateResource(resourceType),
				zLevel:   i,
			}
			gm.tiles[c] = &t
			gm.tilesByZLevel[i] = append(gm.tilesByZLevel[i], &t)

			if c.X == 0 || c.Y == 0 || c.X == gridWidth-1 || c.Y == gridHeight-1 {
				c.Walkable = false
			}
		}
	}

	go gm.handleGetPathRequests()

	return &gm
}

func (g *GameMap) handleGetPathRequests() {
	for r := range g.getPathChan {
		grid := g.grids[r.endZ]
		r.responseChan <- grid.GetPathFromCells(grid.Get(r.startX, r.startY), grid.Get(r.endX, r.endY), true, true)
	}
}

func (g *GameMap) GetPath(startX, startY, startZ, endX, endY, endZ int) *paths.Path {
	responseChan := make(chan *paths.Path)
	defer close(responseChan)
	g.getPathChan <- getPathRequest{
		startX, startY, startZ, endX, endY, endZ, responseChan,
	}
	return <-responseChan
}

func (g *GameMap) GetMapDimensions() (int, int) {
	return g.grids[0].Width(), g.grids[0].Height()
}

func (g *GameMap) IsWalkable(x, y, z int) bool {
	return g.grids[z].Get(x, y).Walkable
}

func (g *GameMap) GetCellCost(x, y, z int) int {
	return int(g.grids[z].Get(x, y).Cost)
}

func (g *GameMap) SwitchWalkable(x, y, z int) {
	c := g.grids[z].Get(x, y)
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

	cl := CamZLevel
	if cl > 5 {
		cl = 5
	}
	for _, t := range g.tilesByZLevel[cl] {
		if t.resource.image == nil || t.cell.X < camXPos-camWidth || t.cell.X > camXPos+camWidth || t.cell.Y < camYPos-camHeight || t.cell.Y > camYPos+camHeight {
			t.drawn = false
			continue
		}

		t.drawn = true

		// Draw the tile
		op := Cam.GetTranslation(float64(t.cell.X*cellWidth), float64(t.cell.Y*cellHeight))

		if CamZLevel > cl {
			op.ColorM.Apply(color.RGBA{
				R: 100,
				G: 100,
				B: 100,
				A: 100,
			})
		}

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

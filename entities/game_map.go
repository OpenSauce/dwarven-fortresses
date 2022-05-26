package entities

import (
	"github.com/OpenSauce/paths"
	camera "github.com/melonfunction/ebiten-camera"
	"github.com/tomknightdev/dwarven-fortresses/components"
	"golang.org/x/exp/rand"
)

type ZTraversable int

const (
	NO ZTraversable = iota
	UP
	DOWN
)

type GameMap struct {
	grids         map[int]*paths.Grid
	tiles         map[*paths.Cell]*Tile
	tilesByZLevel map[int][]*Tile
	cam           *camera.Camera
	cellSize      int
}

type Tile struct {
	cell         *paths.Cell
	resource     *Resource
	drawn        bool
	zLevel       int
	zTraversable ZTraversable
	cellSize     int
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

func (t *Tile) Draw(cam *camera.Camera) {
	op := cam.GetTranslation(float64(t.cell.X*t.cellSize), float64(t.cell.Y*t.cellSize))

	// if camZLevel > cl {
	// 	op.ColorM.Apply(color.RGBA{
	// 		R: 100,
	// 		G: 100,
	// 		B: 100,
	// 		A: 100,
	// 	})
	// }

	t.resource.Draw(cam, op)

	if t.resource.queued {
		cam.Surface.DrawImage(components.CursorImage, op)
	}
}

func (t *Tile) Gathered() {
	t.resource = nil
	t.resource = NewResource(0)
	t.cell.Walkable = true
}

func (t *Tile) SetType(tileType string) {
	if tileType == "stairDown" {
		t.zTraversable = DOWN
		t.resource.image = components.StairDownImage
	} else if tileType == "stairUp" {
		t.zTraversable = UP
		t.resource.image = components.StairUpImage
	}
}

func NewGameMap(gridWidth, gridHeight, cellWidth, cellHeight int) *GameMap {
	gm := GameMap{
		grids:         make(map[int]*paths.Grid),
		tiles:         make(map[*paths.Cell]*Tile),
		tilesByZLevel: make(map[int][]*Tile),
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
				resource: NewResource(resourceType),
				zLevel:   i,
			}
			gm.tiles[c] = &t
			gm.tilesByZLevel[i] = append(gm.tilesByZLevel[i], &t)

			if c.X == 0 || c.Y == 0 || c.X == gridWidth-1 || c.Y == gridHeight-1 {
				c.Walkable = false
			}
		}
	}

	return &gm
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

func (g *GameMap) GetGridForZLevel(zLevel int) *paths.Grid {
	return g.grids[zLevel]
}

func (g *GameMap) GetTilesByZLevel(zLevel int) []*Tile {
	return g.tilesByZLevel[zLevel]
}

func (g *GameMap) GetTileForCell(cell *paths.Cell) *Tile {
	return g.tiles[cell]
}

func (g *GameMap) Update() error {
	for _, t := range g.tiles {
		t.Update()
	}
	return nil
}

func (g *GameMap) Draw(cam *camera.Camera, camZLevel int) {
	camXPos := int(g.cam.X) / g.cellSize
	camYPos := int(g.cam.Y) / g.cellSize

	camWidth := g.cam.Width / 2 / 8 / int(g.cam.Scale)
	camHeight := g.cam.Height / 2 / 8 / int(g.cam.Scale)

	cl := camZLevel
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
		t.Draw(cam)
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

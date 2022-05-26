package components

import (
	"github.com/OpenSauce/paths"
)

type gameMap interface {
	GetGridForZLevel(int) *paths.Grid
	GetTilesByZLevel(int) []*Tile
	// struct {
	// 	zTraversable int // 0=NO, 1=UP, 2=DOWN
	// 	cell         *paths.Cell
	// }
}

type Pathfinder struct {
	getPathChan chan pathRequest
	gm          gameMap
}

type pathRequest struct {
	startX, startY, startZ, endX, endY, endZ int
	responseChan                             chan []PathResponse
}

type PathResponse struct {
	*paths.Path
	zTraversable int
}

type PFTile struct {
	cell         *paths.Cell
	zTraversable int //0=NO, 1=UP, 2=DOWN
	zLevel       int
}

func NewPathfinder(gm gameMap) *Pathfinder {
	p := Pathfinder{
		getPathChan: make(chan pathRequest),
	}

	go p.handleGetPathRequests()

	return &p
}

func (p *Pathfinder) GetPath(startX, startY, startZ, endX, endY, endZ int) []PathResponse {
	responseChan := make(chan []PathResponse)
	defer close(responseChan)
	p.getPathChan <- pathRequest{
		startX, startY, startZ, endX, endY, endZ, responseChan,
	}
	return <-responseChan
}

func (p *Pathfinder) handleGetPathRequests() {
	for r := range p.getPathChan {

		// If both start and end tiles are on the same z level, only one grid needs to be traversed
		// PROBLEM: If the start and end z are the same, but are not connected (two seperate caves with stairs to the
		// surface for example)
		if r.startZ == r.endZ {
			grid := p.gm.GetGridForZLevel(r.endZ)
			r.responseChan <- []PathResponse{
				{grid.GetPathFromCells(grid.Get(r.startX, r.startY), grid.Get(r.endX, r.endY), true, true),
					0,
				},
			}
			continue
		}

		// On each level we need to map each traversable tile to the end position, if the end z level is different,
		// then the end position(s) is each traversable tile

		// Get map of traversable tile by type
		travTiles := make(map[int][]PFTile)
		for _, t := range p.gm.GetTilesByZLevel(r.startZ) {
			if t.zTraversable != 0 {
				travTiles[t.zTraversable] = append(travTiles[t.zTraversable], t)
			}
		}

		travPathsByZ := make(map[int][]struct {
			tile PFTile
			path *paths.Path
		})

		routesByCost := [][]PathResponse{}

		//Create paths to each traversable tile
		if r.endZ < r.startZ {
			grid := p.gm.GetGridForZLevel(r.startZ)
			for _, t := range travTiles[2] {
				p := grid.GetPathFromCells(grid.Get(r.startX, r.startY), grid.Get(t.cell.X, t.cell.Y), true, true)

				if p == nil {
					continue
				}

				travPathsByZ[t.zLevel] = append(travPathsByZ[t.zLevel], struct {
					tile PFTile
					path *paths.Path
				}{
					tile: t,
					path: p,
				})
			}

			if r.endZ == r.startZ-1 {
				// For each path found, calculate that tile to end pos
				grid := p.gm.GetGridForZLevel(r.endZ)
				for _, p := range travPathsByZ[r.startZ] {
					sx := p.path.Cells[p.path.Length()-1].X
					sy := p.path.Cells[p.path.Length()-1].Y
					path := grid.GetPathFromCells(grid.Get(sx, sy), grid.Get(r.endX, r.endY), true, true)

					if path == nil {
						continue
					}

					routesByCost = append(routesByCost, []PathResponse{
						{p.path, 2},
						{path, 0},
					})
				}
			}

			r.responseChan <- routesByCost[0]
		}
	}
}

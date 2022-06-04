package systems

import (
	"math/rand"
	"time"

	"github.com/sedyh/mizu/pkg/engine"
	"github.com/tomknightdev/dwarven-fortresses/enums"
)

type Nature struct {
	GrowTimer        int
	CurrentGrowTimer int
	GameMap          GameMap
}

func NewNature(gameMap GameMap) *Nature {
	return &Nature{
		GrowTimer:        100,
		CurrentGrowTimer: 0,
		GameMap:          gameMap,
	}
}

func (n *Nature) Update(w engine.World) {
	if n.CurrentGrowTimer < n.GrowTimer {
		n.CurrentGrowTimer++
		return
	}
	n.CurrentGrowTimer = 0

	// Pick a random tile, if dirt, make grass
	tiles := n.GameMap.GetTilesByType(enums.TileTypeDirt)
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(len(tiles))

	n.GameMap.UpdateTile(enums.TileTypeDirt, r, enums.TileTypeGrass)
}

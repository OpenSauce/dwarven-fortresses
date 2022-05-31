package systems

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sedyh/mizu/pkg/engine"
	"github.com/tomknightdev/dwarven-fortresses/assets"
	"github.com/tomknightdev/dwarven-fortresses/components"
	"github.com/tomknightdev/dwarven-fortresses/enums"
)

type GameMap interface {
	GetTilesByType(enums.TileTypeEnum) []components.Position
}

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
	tile := tiles[r]

	tileMap := w.View(components.TileMap{}, components.Sprite{}, components.Position{}).Filter()
	for _, tm := range tileMap {
		var tmPos *components.Position
		var tmSprite *components.Sprite

		tm.Get(&tmPos, &tmSprite)
		if tmPos.Z == tile.Z {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(tile.X*assets.CellSize), float64(tile.Y*assets.CellSize))
			r = rand.Intn(3)
			tmSprite.Image.DrawImage(assets.Images[fmt.Sprintf("grass%d", r)], op)
			break
		}
	}
}

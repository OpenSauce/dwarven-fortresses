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

type Nature struct {
	GrowTimer        int
	CurrentGrowTimer int
}

func NewNature() *Nature {
	return &Nature{
		GrowTimer:        100,
		CurrentGrowTimer: 0,
	}
}

func (n *Nature) Update(w engine.World) {
	if n.CurrentGrowTimer < n.GrowTimer {
		n.CurrentGrowTimer++
		return
	}
	n.CurrentGrowTimer = 0

	// Pick a random tile, if dirt, make grass
	tiles := w.View(components.TileType{}, components.Position{}).Filter()
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(len(tiles))

	var tt *components.TileType
	var pos *components.Position
	tiles[r].Get(&tt, &pos)
	if tt.TileTypeEnum == enums.TileTypeDirt {
		tileMap := w.View(components.TileMap{}, components.Sprite{}, components.Position{}).Filter()
		for _, tm := range tileMap {
			var tmPos *components.Position
			var tmSprite *components.Sprite

			tm.Get(&tmPos, &tmSprite)
			if tmPos.Z == pos.Z {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(pos.X*assets.CellSize), float64(pos.Y*assets.CellSize))
				r = rand.Intn(3)
				tmSprite.Image.DrawImage(assets.Images[fmt.Sprintf("grass%d", r)], op)
				break
			}
		}
	}
}

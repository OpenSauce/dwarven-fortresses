package systems

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sedyh/mizu/pkg/engine"
	"github.com/tomknightdev/dwarven-fortresses/components"
	"github.com/tomknightdev/dwarven-fortresses/enums"
	"github.com/tomknightdev/dwarven-fortresses/helpers"
)

type Nature struct {
}

func NewNature() *Nature {
	return &Nature{}
}

func (n *Nature) Update(w engine.World) {
	ne, found := w.View(components.NatureSingleton{}).Get()
	if !found {
		panic("unable to find entity with nature component")
	}
	var nc *components.NatureSingleton
	ne.Get(&nc)

	if nc.CurrentGrowTimer < nc.GrowTimer {
		nc.CurrentGrowTimer++
		return
	}
	nc.CurrentGrowTimer = 0

	gms, found := w.View(components.GameMapSingleton{}).Get()
	if !found {
		panic("game map singleton not found")
	}

	var gmComp *components.GameMapSingleton
	gms.Get(&gmComp)

	// Pick a random tile, if dirt, make grass
	tiles := gmComp.TilesByType[enums.TileTypeDirt]
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(len(tiles))

	gmComp.TilesToUpdateChan <- struct {
		FromTileType enums.TileTypeEnum
		ToTileType   enums.TileTypeEnum
		TileIndex    int
	}{
		enums.TileTypeDirt,
		enums.TileTypeGrass,
		r,
	}
}

func (n *Nature) Draw(w engine.World, screen *ebiten.Image) {
	ne, found := w.View(components.NatureSingleton{}).Get()
	if !found {
		panic("unable to find entity with nature component")
	}
	var nc *components.NatureSingleton
	ne.Get(&nc)

	ents := w.View(components.Nature{}, components.Sprite{}, components.Position{}).Filter()
	helpers.DrawImages(w, screen, nc.OffScreen, ents)

	// var s *components.Sprite
	// var p *components.Position

	// ents.Each(func(e engine.Entity) {
	// 	e.Get(&s, &p)

	// 	helpers.DrawImage(w, screen, *p, s.Image)
	// })
}

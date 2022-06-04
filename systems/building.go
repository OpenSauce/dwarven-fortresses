package systems

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sedyh/mizu/pkg/engine"
	"github.com/tomknightdev/dwarven-fortresses/components"
	"github.com/tomknightdev/dwarven-fortresses/helpers"
)

type Building struct {
}

func NewBuilding() *Building {
	return &Building{}
}

func (b *Building) Update(w engine.World) {

}

func (b *Building) Draw(w engine.World, screen *ebiten.Image) {
	ents := w.View(components.Building{}, components.Position{}, components.Sprite{})

	var p *components.Position
	var s *components.Sprite

	ents.Each(func(e engine.Entity) {
		e.Get(&p, &s)

		helpers.DrawImage(w, screen, *p, s.Image)
	})
}

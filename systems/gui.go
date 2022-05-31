package systems

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sedyh/mizu/pkg/engine"
	"github.com/tomknightdev/dwarven-fortresses/components"
)

type Gui struct {
}

func NewGui() *Gui {
	return &Gui{}
}

func (g *Gui) Draw(w engine.World, screen *ebiten.Image) {
	view := w.View(components.Gui{}, components.Sprite{})
	view.Each(func(e engine.Entity) {
		var gui *components.Gui
		var sprite *components.Sprite
		e.Get(&gui, &sprite)
		op := &ebiten.DrawImageOptions{}

		op.GeoM.Translate(float64(gui.X)/gui.Scale, float64(gui.Y)/gui.Scale)
		op.GeoM.Scale(gui.Scale, gui.Scale)
		screen.DrawImage(sprite.Image, op)
	})

}

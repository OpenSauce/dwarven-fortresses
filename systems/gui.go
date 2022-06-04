package systems

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sedyh/mizu/pkg/engine"
	"github.com/tomknightdev/dwarven-fortresses/components"
	"github.com/tomknightdev/dwarven-fortresses/enums"
)

type Gui struct {
}

func NewGui() *Gui {
	return &Gui{}
}

func (g *Gui) Update(w engine.World) {
	var inputSingleton *components.InputSingleton
	is, found := w.View(components.InputSingleton{}).Get()
	if !found {
		panic("input singleton was not found")
	}
	is.Get(&inputSingleton)

	if inputSingleton.IsMouseLeftPressed {
		guis := w.View(components.Gui{}, components.Sprite{}).Filter()
		for _, g := range guis {
			var gsp *components.Sprite
			var gui *components.Gui
			g.Get(&gsp, &gui)

			if gui.Within(inputSingleton.MousePosX, inputSingleton.MousePosY) {
				switch gui.Action {
				case enums.GuiActionStair:
					inputSingleton.InputMode = enums.InputModeBuild
				case enums.GuiActionChop:
					inputSingleton.InputMode = enums.InputModeChop
				case enums.GuiActionMine:
					inputSingleton.InputMode = enums.InputModeMine
				}
			}
		}
	}
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

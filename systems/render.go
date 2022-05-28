package systems

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/sedyh/mizu/pkg/engine"
	"github.com/tomknightdev/dwarven-fortresses/components"
)

type Render struct {
	width, height, cellSize int
	offscreen               *ebiten.Image
}

func NewRender(w, h, cs int) *Render {
	return &Render{
		width:     w,
		height:    h,
		cellSize:  cs,
		offscreen: ebiten.NewImage(w*cs, h*cs),
	}
}

// Render one frame
func (r *Render) Draw(w engine.World, screen *ebiten.Image) {
	// But choose the right entities yourself
	view := w.View(components.Position{}, components.Sprite{})
	view.Each(func(entity engine.Entity) {
		var pos *components.Position
		var spr *components.Sprite
		entity.Get(&pos, &spr)

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(pos.X*r.cellSize), float64(pos.Y*r.cellSize))
		r.offscreen.DrawImage(spr.Image, op)
	})

	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(r.offscreen, op)
	r.offscreen.Clear()

	msg := fmt.Sprintf("TPS: %0.2f FPS: %0.2f\n",
		ebiten.CurrentTPS(), ebiten.CurrentFPS())

	// camXPos := int(Cam.X) / cellWidth
	// camYPos := int(Cam.Y) / cellHeight

	// msg += fmt.Sprintf("CAM SIZE: %d / %d\n", camXPos, camYPos)
	// msg += fmt.Sprintf("TILES DRAWN: %d\n", g.gameMap.DrawnTileCount())
	// msg += fmt.Sprintf("CAM SCALE: %0.2f", Cam.Scale)
	// msg += fmt.Sprintf("CAM Z LEVEL: %d", CamZLevel)

	// for i, u := range g.units {
	// 	msg += fmt.Sprintf("%d: %d\n", i, u.Energy)
	// }

	ebitenutil.DebugPrint(screen, msg)
}

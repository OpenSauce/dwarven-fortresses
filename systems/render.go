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
	tilesByZ                map[int][]engine.Entity
}

func NewRender(w, h, cs int, tilesByZ map[int][]engine.Entity) *Render {
	return &Render{
		width:     w,
		height:    h,
		cellSize:  cs,
		offscreen: ebiten.NewImage(w*cs, h*cs),
		tilesByZ:  tilesByZ,
	}
}

func (r *Render) Draw(w engine.World, screen *ebiten.Image) {
	// Camera
	camera, found := w.View(components.Zoom{}, components.Position{}).Get()
	if !found {
		return
	}
	var zoom *components.Zoom
	var camPos *components.Position
	camera.Get(&zoom, &camPos)

	// camXPos := int(camPos.X/int(zoom.Value)) / r.cellSize
	// camYPos := int(camPos.Y/int(zoom.Value)) / r.cellSize

	// camWidth := 500 / r.cellSize / int(zoom.Value)
	// camHeight := 500 / r.cellSize / int(zoom.Value)

	// Entities with position and sprite components
	// var fifthLevel engine.Entity
	view := w.View(components.Position{}, components.Sprite{})
	view.Each(func(e engine.Entity) {
		var pos *components.Position
		var spr *components.Sprite
		e.Get(&pos, &spr)

		// if pos.Z == 5 && spr.Image.Bounds().Size().X > assets.CellSize {
		// 	fifthLevel = e
		// }

		if pos.Z != camPos.Z || camPos.Z > 5 {
			return
		}

		// if pos.Z != camPos.Z || pos.X < camXPos-camWidth || pos.X > camXPos+camWidth || pos.Y < camYPos-camHeight || pos.Y > camYPos+camHeight {
		// 	return
		// }

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(pos.X*r.cellSize), float64(pos.Y*r.cellSize))
		r.offscreen.DrawImage(spr.Image, op)
	})

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(zoom.Value, zoom.Value)

	// if camPos.Z > 5 && fifthLevel != nil {
	// 	var pos *components.Position
	// 	var spr *components.Sprite
	// 	fifthLevel.Get(&pos, &spr)

	// 	op := &ebiten.DrawImageOptions{}
	// 	op.GeoM.Translate(float64(pos.X*r.cellSize), float64(pos.Y*r.cellSize))
	// 	r.offscreen.DrawImage(spr.Image, op)
	// 	// This is how you adjust the ALPHA value
	// 	op.ColorM.Scale(1, 1, 1, 0.5)
	// }

	ww, wh := ebiten.WindowSize()
	op.GeoM.Translate(-float64(camPos.X-(ww/2)), -float64(camPos.Y-(wh/2)))
	// op.Filter = ebiten.FilterNearest
	screen.DrawImage(r.offscreen, op)
	r.offscreen.Clear()

	// Debug information
	msg := fmt.Sprintf("TPS: %0.2f FPS: %0.2f\n",
		ebiten.CurrentTPS(), ebiten.CurrentFPS())

	msg += fmt.Sprintf("JOB COUNT: %d", len(w.View(components.Task{}).Filter()))
	ebitenutil.DebugPrint(screen, msg)
}

package systems

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/sedyh/mizu/pkg/engine"
	"github.com/tomknightdev/dwarven-fortresses/components"
)

type Debug struct {
}

func NewDebug() *Debug {
	return &Debug{}
}

func (d *Debug) Update(w engine.World) {

}

func (d *Debug) Draw(w engine.World, screen *ebiten.Image) {
	// Debug information
	msg := fmt.Sprintf("TPS: %0.2f FPS: %0.2f\n",
		ebiten.CurrentTPS(), ebiten.CurrentFPS())

	msg += fmt.Sprintf("JOB COUNT: %d\n", len(w.View(components.Task{}).Filter()))

	ebitenutil.DebugPrint(screen, msg)

}

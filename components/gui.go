package components

import (
	"github.com/tomknightdev/dwarven-fortresses/assets"
	"github.com/tomknightdev/dwarven-fortresses/enums"
)

type Gui struct {
	X, Y   int
	Scale  float64
	Action enums.GuiActionEnum
}

func NewGui(x, y int, scale float64, action enums.GuiActionEnum) Gui {
	return Gui{
		X:      x,
		Y:      y,
		Scale:  scale,
		Action: action,
	}
}

func (g Gui) Within(x, y int) bool {
	sx, sy := g.scalePos()
	if x > g.X && x < sx && y > g.Y && y < sy {
		return true
	}

	return false
}

func (g Gui) scalePos() (int, int) {
	x := g.X + int(g.Scale)*assets.CellSize
	y := g.Y + int(g.Scale)*assets.CellSize
	return x, y
}

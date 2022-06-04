package components

import (
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

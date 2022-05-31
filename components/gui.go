package components

import "github.com/tomknightdev/dwarven-fortresses/assets"

type Gui struct {
	X, Y  int
	Scale float64
}

func NewGui(x, y int, scale float64) Gui {
	return Gui{
		X:     x,
		Y:     y,
		Scale: scale,
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

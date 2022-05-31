package entities

import (
	"github.com/tomknightdev/dwarven-fortresses/components"
)

type Input struct {
	MousePos    components.Position
	CursorImage components.Sprite
	Input       components.Input
}

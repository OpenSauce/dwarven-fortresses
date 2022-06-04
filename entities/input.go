package entities

import (
	"github.com/tomknightdev/dwarven-fortresses/components"
)

type Mouse struct {
	components.Position
	components.Sprite
	components.Mouse
}

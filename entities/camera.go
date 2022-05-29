package entities

import "github.com/tomknightdev/dwarven-fortresses/components"

type Camera struct {
	components.Zoom
	components.Input
	components.Position
}

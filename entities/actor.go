package entities

import "github.com/tomknightdev/dwarven-fortresses/components"

type Actor struct {
	components.Position
	components.Sprite
	components.Move
	components.Worker
	components.Inventory
}

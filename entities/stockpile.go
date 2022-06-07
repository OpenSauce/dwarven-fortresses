package entities

import "github.com/tomknightdev/dwarven-fortresses/components"

type Stockpile struct {
	components.Designation
	components.Position
	components.Sprite
	components.Inventory
}

package entities

import "github.com/tomknightdev/dwarven-fortresses/components"

type Tree struct {
	components.Choppable
	components.Position
	components.Sprite
	components.Resource
	components.Drops
	components.Nature
}

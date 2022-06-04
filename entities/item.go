package entities

import "github.com/tomknightdev/dwarven-fortresses/components"

type Item struct {
	components.Item
	components.Position
	components.Sprite
}

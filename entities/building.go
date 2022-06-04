package entities

import "github.com/tomknightdev/dwarven-fortresses/components"

type Building struct {
	components.Building
	components.Position
	components.Sprite
	components.TileType
}

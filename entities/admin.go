package entities

import "github.com/tomknightdev/dwarven-fortresses/components"

type Admin struct {
	components.InputSingleton
	components.GameMapSingleton
	components.NatureSingleton
}

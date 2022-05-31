package components

import "github.com/tomknightdev/dwarven-fortresses/enums"

type Input struct {
	MouseMode enums.MouseModeEnum
}

func NewInput() Input {
	return Input{}
}

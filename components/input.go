package components

import "github.com/tomknightdev/dwarven-fortresses/enums"

type Input struct {
	InputMode enums.InputModeEnum
}

func NewInput() Input {
	return Input{}
}

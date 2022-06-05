package components

import "github.com/tomknightdev/dwarven-fortresses/enums"

type InputSingleton struct {
	// Mouse values
	MouseLeftPressDuration  bool
	IsMouseLeftPressed      bool
	IsMouseLeftReleased     bool
	MouseRightPressDuration bool
	IsMouseRightPressed     bool
	IsMouseRightReleased    bool
	MouseWheel              float64
	MousePosX               int
	MousePosY               int
	MouseWorldPosX          int
	MouseWorldPosY          int
	LeftClickedTiles        []Position
	RightClickedTiles       []Position
	InGui                   bool

	// Keyboard values
	IsCameraLeftPressed     bool
	IsCameraRightPressed    bool
	IsCameraUpPressed       bool
	IsCameraDownPressed     bool
	IsCameraLowerPressed    bool
	IsCameraRaisePressed    bool
	IsResetInputModePressed bool

	// Input mode
	InputMode enums.InputModeEnum
}

func NewInputSingleton() InputSingleton {
	return InputSingleton{}
}

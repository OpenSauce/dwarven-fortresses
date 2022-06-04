package systems

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/sedyh/mizu/pkg/engine"
	"github.com/tomknightdev/dwarven-fortresses/components"
)

type Input struct {
}

func NewInput() *Input {
	return &Input{}
}

func (i *Input) Update(w engine.World) {
	var inputSingleton *components.InputSingleton
	is, found := w.View(components.InputSingleton{}).Get()
	if !found {
		panic("input singleton was not found")
	}
	is.Get(&inputSingleton)

	// Keyboard updates
	if inpututil.KeyPressDuration(ebiten.KeyD) > 0 {
		inputSingleton.IsCameraRightPressed = true
	} else {
		inputSingleton.IsCameraRightPressed = false
	}

	if inpututil.KeyPressDuration(ebiten.KeyA) > 0 {
		inputSingleton.IsCameraLeftPressed = true
	} else {
		inputSingleton.IsCameraLeftPressed = false
	}

	if inpututil.KeyPressDuration(ebiten.KeyW) > 0 {
		inputSingleton.IsCameraUpPressed = true
	} else {
		inputSingleton.IsCameraUpPressed = false
	}

	if inpututil.KeyPressDuration(ebiten.KeyS) > 0 {
		inputSingleton.IsCameraDownPressed = true
	} else {
		inputSingleton.IsCameraDownPressed = false
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyE) {
		inputSingleton.IsCameraLowerPressed = true
	} else {
		inputSingleton.IsCameraLowerPressed = false
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		inputSingleton.IsCameraRaisePressed = true
	} else {
		inputSingleton.IsCameraRaisePressed = false
	}

	// Mouse updates
	_, wy := ebiten.Wheel()
	inputSingleton.MouseWheel = wy

	inputSingleton.MousePosX, inputSingleton.MousePosY = ebiten.CursorPosition()

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		inputSingleton.IsMouseLeftPressed = true
	} else {
		inputSingleton.IsMouseLeftPressed = false
	}

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		inputSingleton.IsMouseLeftReleased = true
	} else {
		inputSingleton.IsMouseLeftReleased = false
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
		inputSingleton.IsMouseRightPressed = true
	} else {
		inputSingleton.IsMouseRightPressed = false
	}

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonRight) {
		inputSingleton.IsMouseRightReleased = true
	} else {
		inputSingleton.IsMouseRightReleased = false
	}
}

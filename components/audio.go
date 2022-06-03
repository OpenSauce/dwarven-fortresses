package components

import "github.com/hajimehoshi/ebiten/v2/audio"

type Audio struct {
	*audio.Player
}

func NewAudio() Audio {
	return Audio{}
}

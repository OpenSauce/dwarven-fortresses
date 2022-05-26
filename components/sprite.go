package components

import "github.com/hajimehoshi/ebiten/v2"

type Sprite struct {
	image *ebiten.Image
}

func NewSprite(image *ebiten.Image) Sprite {
	return Sprite{image}
}

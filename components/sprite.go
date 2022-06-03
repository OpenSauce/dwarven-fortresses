package components

import "github.com/hajimehoshi/ebiten/v2"

type Sprite struct {
	Image       *ebiten.Image
	RenderOrder int
}

func NewSprite(image *ebiten.Image, renderOrder int) Sprite {
	return Sprite{image, renderOrder}
}

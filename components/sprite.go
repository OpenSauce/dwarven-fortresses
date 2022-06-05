package components

import "github.com/hajimehoshi/ebiten/v2"

type Sprite struct {
	Image *ebiten.Image
	Drawn bool
}

func NewSprite(image *ebiten.Image) Sprite {
	return Sprite{
		Image: image,
		Drawn: true,
	}
}

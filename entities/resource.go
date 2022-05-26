package entities

import (
	"math/rand"

	_ "embed"

	"github.com/hajimehoshi/ebiten/v2"
	camera "github.com/melonfunction/ebiten-camera"
	"github.com/tomknightdev/dwarven-fortresses/components"
)

type ResourceType int

const (
	Dirt ResourceType = iota
	Grass
	Tree
	Water
	Rock
	Empty
)

var ()

type Resource struct {
	image        *ebiten.Image
	resourceType ResourceType
	queued       bool
}

func SetupResources(cellSize int) {

}

func NewResource(rt ResourceType) *Resource {
	r := Resource{
		resourceType: rt,
	}

	switch rt {
	case Dirt:
		r.image = components.DirtImages[rand.Intn(len(components.DirtImages))]
	case Grass:
		r.image = components.GrassImages[rand.Intn(len(components.GrassImages))]
	case Tree:
		r.image = components.TreeImages[rand.Intn(len(components.TreeImages))]
	case Water:
		r.image = components.WaterImages[rand.Intn(len(components.WaterImages))]
	case Rock:
		r.image = components.RockImages[rand.Intn(len(components.RockImages))]
	case Empty:
		r.image = nil
	}

	return &r
}

func (r *Resource) Update() error {
	return nil
}

func (r *Resource) Draw(cam *camera.Camera, op *ebiten.DrawImageOptions) {
	cam.Surface.DrawImage(r.image, op)
}

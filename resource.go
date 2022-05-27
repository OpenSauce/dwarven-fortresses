package main

import (
	"image"
	"math/rand"

	_ "embed"

	"github.com/hajimehoshi/ebiten/v2"
)

type ResourceType int

const (
	Dirt ResourceType = iota - 2
	Grass
	Tree
	Water
	Rock
	Empty
)

var (
	rockImages  []image.Rectangle
	dirtImages  []image.Rectangle
	grassImages []image.Rectangle
	treeImages  []image.Rectangle
	waterImages []image.Rectangle
)

type Resource struct {
	image        *ebiten.Image
	resourceType ResourceType
	queued       bool
}

func init() {
	rockImages = append(rockImages, image.Rect(20*cellWidth, 1*cellHeight, 21*cellWidth, 2*cellHeight))

	dirtImages = append(dirtImages, image.Rect(0*cellWidth, 0*cellHeight, 1*cellWidth, 1*cellHeight))
	dirtImages = append(dirtImages, image.Rect(1*cellWidth, 0*cellHeight, 2*cellWidth, 1*cellHeight))

	grassImages = append(grassImages, image.Rect(5*cellWidth, 0*cellHeight, 6*cellWidth, 1*cellHeight))
	grassImages = append(grassImages, image.Rect(6*cellWidth, 0*cellHeight, 7*cellWidth, 1*cellHeight))
	grassImages = append(grassImages, image.Rect(7*cellWidth, 0*cellHeight, 8*cellWidth, 1*cellHeight))

	treeImages = append(treeImages, image.Rect(0*cellWidth, 1*cellHeight, 1*cellWidth, 2*cellHeight))
	treeImages = append(treeImages, image.Rect(1*cellWidth, 1*cellHeight, 2*cellWidth, 2*cellHeight))
	treeImages = append(treeImages, image.Rect(2*cellWidth, 1*cellHeight, 3*cellWidth, 2*cellHeight))
	treeImages = append(treeImages, image.Rect(3*cellWidth, 1*cellHeight, 4*cellWidth, 2*cellHeight))
	treeImages = append(treeImages, image.Rect(4*cellWidth, 1*cellHeight, 5*cellWidth, 2*cellHeight))
	treeImages = append(treeImages, image.Rect(5*cellWidth, 1*cellHeight, 6*cellWidth, 2*cellHeight))

	waterImages = append(waterImages, image.Rect(14*cellWidth, 5*cellHeight, 15*cellWidth, 6*cellHeight))
}

func CreateResource(rt ResourceType) *Resource {
	r := Resource{
		resourceType: rt,
	}

	switch rt {
	case Dirt:
		r.image = TilesetImage.SubImage(dirtImages[rand.Intn(len(dirtImages))]).(*ebiten.Image)
	case Grass:
		r.image = TilesetImage.SubImage(grassImages[rand.Intn(len(grassImages))]).(*ebiten.Image)
	case Tree:
		r.image = TilesetImage.SubImage(treeImages[rand.Intn(len(treeImages))]).(*ebiten.Image)
	case Water:
		r.image = TilesetImage.SubImage(waterImages[rand.Intn(len(waterImages))]).(*ebiten.Image)
	case Rock:
		r.image = TilesetImage.SubImage(rockImages[rand.Intn(len(rockImages))]).(*ebiten.Image)
	case Empty:
		r.image = nil
	}

	return &r
}

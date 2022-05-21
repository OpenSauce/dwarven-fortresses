package main

import (
	"bytes"
	"image"
	"image/png"
	"log"
	"math/rand"

	_ "embed"

	"github.com/hajimehoshi/ebiten/v2"
)

type ResourceType int

const (
	Dirt ResourceType = iota
	Grass
	Tree
	Water
)

var (
	//go:embed resources/terrain.png
	terrain_sheet []byte
	terrainImage  *ebiten.Image

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
	img, err := png.Decode(bytes.NewReader(terrain_sheet))
	if err != nil {
		log.Fatal(err)
	}
	terrainImage = ebiten.NewImageFromImage(img)

	dirtImages = append(dirtImages, image.Rect(0*cellWidth, 0*cellHeight, 1*cellWidth, 1*cellHeight))
	dirtImages = append(dirtImages, image.Rect(1*cellWidth, 0*cellHeight, 2*cellWidth, 1*cellHeight))

	grassImages = append(grassImages, image.Rect(2*cellWidth, 0*cellHeight, 3*cellWidth, 1*cellHeight))
	grassImages = append(grassImages, image.Rect(3*cellWidth, 0*cellHeight, 4*cellWidth, 1*cellHeight))

	treeImages = append(treeImages, image.Rect(4*cellWidth, 0*cellHeight, 5*cellWidth, 1*cellHeight))
	treeImages = append(treeImages, image.Rect(5*cellWidth, 0*cellHeight, 6*cellWidth, 1*cellHeight))
	treeImages = append(treeImages, image.Rect(6*cellWidth, 0*cellHeight, 7*cellWidth, 1*cellHeight))

	waterImages = append(waterImages, image.Rect(7*cellWidth, 0*cellHeight, 8*cellWidth, 1*cellHeight))
}

func CreateResource(rt ResourceType) *Resource {
	r := Resource{
		resourceType: rt,
	}

	switch rt {
	case Dirt:
		r.image = terrainImage.SubImage(dirtImages[rand.Intn(len(dirtImages))]).(*ebiten.Image)
	case Grass:
		r.image = terrainImage.SubImage(grassImages[rand.Intn(len(grassImages))]).(*ebiten.Image)
	case Tree:
		r.image = terrainImage.SubImage(treeImages[rand.Intn(len(treeImages))]).(*ebiten.Image)
	case Water:
		r.image = terrainImage.SubImage(waterImages[rand.Intn(len(waterImages))]).(*ebiten.Image)
	}

	return &r
}

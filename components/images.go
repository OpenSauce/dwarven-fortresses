package components

import (
	"bytes"
	"image"
	"image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	//go:embed resources/map.png
	Tileset      []byte
	TilesetImage *ebiten.Image

	//go:embed resources/trans_map.png
	TransTileset      []byte
	TransTilesetImage *ebiten.Image

	StairUpImage   *ebiten.Image
	StairDownImage *ebiten.Image
	CursorImage    *ebiten.Image

	RockImages  []*ebiten.Image
	DirtImages  []*ebiten.Image
	GrassImages []*ebiten.Image
	TreeImages  []*ebiten.Image
	WaterImages []*ebiten.Image
)

func SetupImages(cellSize int) {
	img, err := png.Decode(bytes.NewReader(Tileset))
	if err != nil {
		log.Fatal(err)
	}
	TilesetImage = ebiten.NewImageFromImage(img)

	img, err = png.Decode(bytes.NewReader(TransTileset))
	if err != nil {
		log.Fatal(err)
	}
	TransTilesetImage = ebiten.NewImageFromImage(img)

	RockImages = append(RockImages, TilesetImage.SubImage(image.Rect(20*cellSize, 1*cellSize, 21*cellSize, 2*cellSize)).(*ebiten.Image))

	DirtImages = append(DirtImages, TilesetImage.SubImage(image.Rect(0*cellSize, 0*cellSize, 1*cellSize, 1*cellSize)).(*ebiten.Image))
	DirtImages = append(DirtImages, TilesetImage.SubImage(image.Rect(1*cellSize, 0*cellSize, 2*cellSize, 1*cellSize)).(*ebiten.Image))

	GrassImages = append(GrassImages, TilesetImage.SubImage(image.Rect(5*cellSize, 0*cellSize, 6*cellSize, 1*cellSize)).(*ebiten.Image))
	GrassImages = append(GrassImages, TilesetImage.SubImage(image.Rect(6*cellSize, 0*cellSize, 7*cellSize, 1*cellSize)).(*ebiten.Image))
	GrassImages = append(GrassImages, TilesetImage.SubImage(image.Rect(7*cellSize, 0*cellSize, 8*cellSize, 1*cellSize)).(*ebiten.Image))

	TreeImages = append(TreeImages, TilesetImage.SubImage(image.Rect(0*cellSize, 1*cellSize, 1*cellSize, 2*cellSize)).(*ebiten.Image))
	TreeImages = append(TreeImages, TilesetImage.SubImage(image.Rect(1*cellSize, 1*cellSize, 2*cellSize, 2*cellSize)).(*ebiten.Image))
	TreeImages = append(TreeImages, TilesetImage.SubImage(image.Rect(2*cellSize, 1*cellSize, 3*cellSize, 2*cellSize)).(*ebiten.Image))
	TreeImages = append(TreeImages, TilesetImage.SubImage(image.Rect(3*cellSize, 1*cellSize, 4*cellSize, 2*cellSize)).(*ebiten.Image))
	TreeImages = append(TreeImages, TilesetImage.SubImage(image.Rect(4*cellSize, 1*cellSize, 5*cellSize, 2*cellSize)).(*ebiten.Image))
	TreeImages = append(TreeImages, TilesetImage.SubImage(image.Rect(5*cellSize, 1*cellSize, 6*cellSize, 2*cellSize)).(*ebiten.Image))

	WaterImages = append(WaterImages, TilesetImage.SubImage(image.Rect(14*cellSize, 5*cellSize, 15*cellSize, 6*cellSize)).(*ebiten.Image))

	StairUpImage = TransTilesetImage.SubImage(image.Rect(2*cellSize, 6*cellSize, 3*cellSize, 7*cellSize)).(*ebiten.Image)
	StairDownImage = TransTilesetImage.SubImage(image.Rect(3*cellSize, 6*cellSize, 4*cellSize, 7*cellSize)).(*ebiten.Image)

	CursorImage = TransTilesetImage.SubImage(image.Rect(29*cellSize, 14*cellSize, 30*cellSize, 15*cellSize)).(*ebiten.Image)
}

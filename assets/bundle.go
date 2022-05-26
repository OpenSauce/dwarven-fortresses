package assets

import (
	"bytes"
	_ "embed"
	"image"
	"image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

// This package is for loading all the images and storing world information

var (
	WorldWidth  = 100
	WorldHeight = 100
	WorldLevels = 30
	CellSize    = 16

	//go:embed resources/map.png
	Tileset      []byte
	TilesetImage *ebiten.Image

	//go:embed resources/trans_map.png
	TransTileset      []byte
	TransTilesetImage *ebiten.Image

	Images = make(map[string]*ebiten.Image)
)

func init() {
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

	Images["cursor"] = TransTilesetImage.SubImage(image.Rect(29*CellSize, 14*CellSize, 30*CellSize, 15*CellSize)).(*ebiten.Image)

	Images["rock"] = TilesetImage.SubImage(image.Rect(20*CellSize, 1*CellSize, 21*CellSize, 2*CellSize)).(*ebiten.Image)

	Images["dirt0"] = TilesetImage.SubImage(image.Rect(0*CellSize, 0*CellSize, 1*CellSize, 1*CellSize)).(*ebiten.Image)
	Images["dirt1"] = TilesetImage.SubImage(image.Rect(1*CellSize, 0*CellSize, 2*CellSize, 1*CellSize)).(*ebiten.Image)

	Images["grass0"] = TilesetImage.SubImage(image.Rect(5*CellSize, 0*CellSize, 6*CellSize, 1*CellSize)).(*ebiten.Image)
	Images["grass1"] = TilesetImage.SubImage(image.Rect(6*CellSize, 0*CellSize, 7*CellSize, 1*CellSize)).(*ebiten.Image)
	Images["grass2"] = TilesetImage.SubImage(image.Rect(7*CellSize, 0*CellSize, 8*CellSize, 1*CellSize)).(*ebiten.Image)

	Images["tree0"] = TilesetImage.SubImage(image.Rect(0*CellSize, 1*CellSize, 1*CellSize, 2*CellSize)).(*ebiten.Image)
	Images["tree1"] = TilesetImage.SubImage(image.Rect(1*CellSize, 1*CellSize, 2*CellSize, 2*CellSize)).(*ebiten.Image)
	Images["tree2"] = TilesetImage.SubImage(image.Rect(2*CellSize, 1*CellSize, 3*CellSize, 2*CellSize)).(*ebiten.Image)
	Images["tree3"] = TilesetImage.SubImage(image.Rect(3*CellSize, 1*CellSize, 4*CellSize, 2*CellSize)).(*ebiten.Image)
	Images["tree4"] = TilesetImage.SubImage(image.Rect(4*CellSize, 1*CellSize, 5*CellSize, 2*CellSize)).(*ebiten.Image)
	Images["tree5"] = TilesetImage.SubImage(image.Rect(5*CellSize, 1*CellSize, 6*CellSize, 2*CellSize)).(*ebiten.Image)

	Images["water"] = TilesetImage.SubImage(image.Rect(14*CellSize, 5*CellSize, 15*CellSize, 6*CellSize)).(*ebiten.Image)

}

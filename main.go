package main

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"log"

	_ "embed"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	worldWidth  int = 200
	worldHeight int = 100
	cellWidth   int = 8
	cellHeight  int = 8
)

var (
	//go:embed resources/map001.png
	sprite_sheet []byte
	tilesImage   *ebiten.Image
)

type Game struct {
	// tileMap  []*Tile
	// tileSize int
	gameMap *GameMap
	units   []*Unit
}

func init() {
	img, err := png.Decode(bytes.NewReader(sprite_sheet))
	if err != nil {
		log.Fatal(err)
	}
	tilesImage = ebiten.NewImageFromImage(img)
}

// type Tile struct {
// 	X, Y           int
// 	Hover, Clicked bool
// }

func (g *Game) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		x = x / cellWidth
		y = y / cellHeight
		t := g.gameMap.grid.Get(x, y)
		if t != nil {
			CreateJob(t.X, t.Y)
		}
	}

	count := len(g.units)
	for i := 0; i < count; i++ {
		g.units[i].Running = true
		// g.units[i].Update()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	var i *ebiten.Image

	for x := 0; x < worldWidth; x++ {
		for y := 0; y < worldHeight; y++ {
			t := g.gameMap.grid.Get(x, y)
			if t.Walkable {
				i = tilesImage.SubImage(image.Rectangle{
					Min: image.Pt(0, 0),
					Max: image.Pt(cellWidth, cellHeight),
				}).(*ebiten.Image)
			} else {
				i = tilesImage.SubImage(image.Rectangle{
					Min: image.Pt(cellWidth, 0),
					Max: image.Pt(cellWidth+cellWidth, cellHeight),
				}).(*ebiten.Image)
			}
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(t.X*cellWidth), float64(t.Y*cellHeight))
			screen.DrawImage(i, op)
		}
	}

	count := len(g.units)
	for i := 0; i < count; i++ {
		g.units[i].Draw(screen)
	}

	msg := fmt.Sprintf(`TPS: %0.2f
FPS: %0.2f`, ebiten.CurrentTPS(), ebiten.CurrentFPS())
	ebitenutil.DebugPrint(screen, msg)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	// return 512, 384
	w, h := ebiten.WindowSize()

	return w / 2, h / 2
}

func main() {
	game := Game{
		// tileMap:  []*Tile{},
		// tileSize: 10,
		gameMap: NewGameMap(worldWidth, worldHeight, cellWidth, cellWidth),
	}

	for i := 0; i < 10; i++ {
		game.units = append(game.units, NewUnit(1, 1, game.gameMap, GetNextJob))
	}

	// count := 0
	// for x := 0; x < 100; x++ {
	// 	for y := 0; y < 100; y++ {
	// 		game.tileMap = append(game.tileMap, &Tile{
	// 			X: x * game.tileSize,
	// 			Y: y * game.tileSize,
	// 		})
	// 		// game.tileMap[count].X = x * game.tileSize
	// 		// game.tileMap[count].Y =
	// 		count++
	// 	}
	// }

	ebiten.SetWindowSize(1024, 768)
	ebiten.SetWindowTitle("Mouse Test")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}

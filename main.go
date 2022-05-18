package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	tileMap  []*Tile
	tileSize int
}

type Tile struct {
	X, Y           int
	Hover, Clicked bool
}

func (g *Game) Update() error {
	clicked := false
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		clicked = true
	}

	x, y := ebiten.CursorPosition()
	for _, t := range g.tileMap {
		if x >= t.X && x < t.X+g.tileSize && y >= t.Y && y < t.Y+g.tileSize {
			t.Hover = true
			if clicked {
				t.Clicked = !t.Clicked
			}
		} else {
			t.Hover = false
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, t := range g.tileMap {

		if t.Clicked {
			i := ebiten.NewImage(g.tileSize, g.tileSize)
			i.Fill(color.RGBA{
				R: 255,
				G: 0,
				B: 0,
				A: 255,
			})
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(t.X), float64(t.Y))
			screen.DrawImage(i, op)
		}
		if t.Hover {
			i := ebiten.NewImage(g.tileSize, g.tileSize)
			i.Fill(color.RGBA{
				R: 255,
				G: 255,
				B: 255,
				A: 255,
			})
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(t.X), float64(t.Y))
			screen.DrawImage(i, op)
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	w, h := ebiten.WindowSize() //512, 384

	return w / 2, h / 2
}

func main() {
	game := Game{
		tileMap:  []*Tile{},
		tileSize: 10,
	}

	count := 0
	for x := 0; x < 100; x++ {
		for y := 0; y < 100; y++ {
			game.tileMap = append(game.tileMap, &Tile{
				X: x * game.tileSize,
				Y: y * game.tileSize,
			})
			// game.tileMap[count].X = x * game.tileSize
			// game.tileMap[count].Y =
			count++
		}
	}

	ebiten.SetWindowSize(1024, 768)
	ebiten.SetWindowTitle("Mouse Test")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}

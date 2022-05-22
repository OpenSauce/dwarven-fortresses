package main

import (
	"bytes"
	"fmt"
	"image/png"
	"log"

	_ "embed"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	camera "github.com/melonfunction/ebiten-camera"
)

const (
	worldWidth  int = 250
	worldHeight int = 250
	cellWidth   int = 8
	cellHeight  int = 8
)

var (
	Cam                               *camera.Camera
	LastWindowWidth, LastWindowHeight int

	//go:embed resources/cursor.png
	cursor_sheet []byte
	cursorImage  *ebiten.Image
	msx, msy     int
)

type Game struct {
	gameMap *GameMap
	units   []*Unit
}

func init() {
	img, err := png.Decode(bytes.NewReader(cursor_sheet))
	if err != nil {
		log.Fatal(err)
	}
	cursorImage = ebiten.NewImageFromImage(img)
}

func (g *Game) Update() error {

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		msx, msy = getCursorCellPos()
	}

	if inpututil.MouseButtonPressDuration(ebiten.MouseButtonLeft) > 0 {

	}

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		mex, mey := getCursorCellPos()

		for x := msx; x < mex; x++ {
			for y := msy; y < mey; y++ {
				c := g.gameMap.grid.Get(x, y)
				t := g.gameMap.tiles[c]
				CreateJob(c, t)
			}
		}
	}

	// Move the camera
	if inpututil.KeyPressDuration(ebiten.KeyD) > 0 && Cam.X < float64(worldWidth*cellWidth) {
		Cam.X++
	} else if inpututil.KeyPressDuration(ebiten.KeyA) > 0 && Cam.X > 0 {
		Cam.X--
	}
	if inpututil.KeyPressDuration(ebiten.KeyS) > 0 && Cam.Y < float64(worldHeight*cellHeight) {
		Cam.Y++
	} else if inpututil.KeyPressDuration(ebiten.KeyW) > 0 && Cam.Y > 0 {
		Cam.Y--
	}

	_, wy := ebiten.Wheel()
	if wy > 0 && Cam.Scale < 10 {
		// Cam.Resize(Cam.Width+10, Cam.Height+10)
		Cam.Zoom(1.1)
	} else if wy < 0 && Cam.Scale > 1.1 {
		// Cam.Resize(Cam.Width-10, Cam.Height-10)
		Cam.Zoom(0.9)
	}

	count := len(g.units)
	for i := 0; i < count; i++ {
		g.units[i].Running = true
		// g.units[i].Update()
	}

	for _, t := range g.gameMap.tiles {
		t.Update()
	}

	return nil
}

func getCursorCellPos() (int, int) {
	x, y := Cam.GetCursorCoords()
	xi := int(x) / cellWidth
	yi := int(y) / cellHeight

	return xi, yi
}

func (g *Game) Draw(screen *ebiten.Image) {
	Cam.Surface.Clear()
	g.gameMap.Draw(screen)

	count := len(g.units)
	for i := 0; i < count; i++ {
		g.units[i].Draw(screen)
	}

	cx, cy := getCursorCellPos()
	// Draw the cursoe
	Cam.Surface.DrawImage(cursorImage, Cam.GetTranslation(float64(cx*cellWidth), float64(cy*cellHeight)))

	// Draw to screen and zoom
	Cam.Blit(screen)

	msg := fmt.Sprintf("TPS: %0.2f FPS: %0.2f\n",
		ebiten.CurrentTPS(), ebiten.CurrentFPS())

	// camXPos := int(Cam.X) / cellWidth
	// camYPos := int(Cam.Y) / cellHeight

	// msg += fmt.Sprintf("CAM SIZE: %d / %d\n", camXPos, camYPos)
	// msg += fmt.Sprintf("TILES DRAWN: %d\n", g.gameMap.DrawnTileCount())
	// msg += fmt.Sprintf("CAM SCALE: %0.2f", Cam.Scale)

	// for i, u := range g.units {
	// 	msg += fmt.Sprintf("%d: %d\n", i, u.Energy)
	// }

	ebitenutil.DebugPrint(screen, msg)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	if LastWindowWidth != outsideWidth || LastWindowHeight != outsideHeight {
		Cam.Resize(outsideWidth, outsideHeight)
		LastWindowWidth = outsideWidth
		LastWindowHeight = outsideHeight
	}
	return outsideWidth, outsideHeight
	// Cam.Resize(outsideWidth, outsideHeight)
	// return outsideWidth, outsideHeight
}

func main() {
	// Create camera
	Cam = camera.NewCamera(1024, 768, float64(worldWidth*cellWidth/2), float64(worldHeight*cellHeight/2), 0, 1)

	game := Game{
		gameMap: NewGameMap(worldWidth, worldHeight, cellWidth, cellWidth),
	}

	for i := 0; i < 50; i++ {
		game.units = append(game.units, NewUnit(worldWidth/2, worldHeight/2, game.gameMap, GetNextJob))
	}

	ebiten.SetWindowSize(1024, 768)
	ebiten.SetWindowTitle("Mouse Test")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}

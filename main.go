package main

import (
	"fmt"
	"log"
	"math/rand"

	_ "embed"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	camera "github.com/melonfunction/ebiten-camera"
	"github.com/tomknightdev/dwarven-fortresses/components"
	"github.com/tomknightdev/dwarven-fortresses/entities"
	"github.com/tomknightdev/dwarven-fortresses/managers"
)

const (
	worldWidth  int = 10
	worldHeight int = 10
	cellSize    int = 16

	units int = 1
)

var (
	Cam                               *camera.Camera
	CamZLevel                         int
	LastWindowWidth, LastWindowHeight int

	msx, msy int
)

type Game struct {
	gameMap    *entities.GameMap
	units      []*entities.Unit
	jobManager *managers.JobManager
}

func init() {

}

func (g *Game) Update() error {
	// Designate gather area
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		msx, msy = getCursorCellPos()
	}

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		mex, mey := getCursorCellPos()

		for x := msx; x <= mex; x++ {
			for y := msy; y <= mey; y++ {
				c := g.gameMap.GetGridForZLevel(CamZLevel).Get(x, y)
				t := g.gameMap.GetTileForCell(c)
				g.jobManager.CreateJob(c, t, managers.Gather)
			}
		}
	}

	// Place z travesable tile
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonRight) {
		mex, mey := getCursorCellPos()

		c := g.gameMap.GetGridForZLevel(CamZLevel).Get(mex, mey)
		t := g.gameMap.GetTileForCell(c)
		g.jobManager.CreateJob(c, t, managers.StairDown)
		c = g.gameMap.GetGridForZLevel(CamZLevel-1).Get(mex, mey)
		t = g.gameMap.GetTileForCell(c)
		g.jobManager.CreateJob(c, t, managers.StairUp)
	}

	// Move the camera
	if inpututil.KeyPressDuration(ebiten.KeyD) > 0 && Cam.X < float64(worldWidth*cellSize) {
		Cam.X += 2
	} else if inpututil.KeyPressDuration(ebiten.KeyA) > 0 && Cam.X > 0 {
		Cam.X -= 2
	}
	if inpututil.KeyPressDuration(ebiten.KeyS) > 0 && Cam.Y < float64(worldHeight*cellSize) {
		Cam.Y += 2
	} else if inpututil.KeyPressDuration(ebiten.KeyW) > 0 && Cam.Y > 0 {
		Cam.Y -= 2
	}

	// Zoom the camera
	_, wy := ebiten.Wheel()
	if wy > 0 && Cam.Scale < 10 {
		// Cam.Resize(Cam.Width+10, Cam.Height+10)
		Cam.Zoom(1.1)
	} else if wy < 0 && Cam.Scale > 1.1 {
		// Cam.Resize(Cam.Width-10, Cam.Height-10)
		Cam.Zoom(0.9)
	}

	// Adjust camera z level
	if inpututil.IsKeyJustPressed(ebiten.KeyE) && CamZLevel > 0 {
		CamZLevel--
	} else if inpututil.IsKeyJustPressed(ebiten.KeyQ) && CamZLevel < 9 {
		CamZLevel++
	}

	count := len(g.units)
	for i := 0; i < count; i++ {
		g.units[i].Running = true
		// g.units[i].Update()
	}

	g.gameMap.Update()

	return nil
}

func getCursorCellPos() (int, int) {
	x, y := Cam.GetCursorCoords()
	xi := int(x) / cellSize
	yi := int(y) / cellSize

	return xi, yi
}

func (g *Game) Draw(screen *ebiten.Image) {
	Cam.Surface.Clear()
	g.gameMap.Draw(screen, CamZLevel)

	count := len(g.units)
	for i := 0; i < count; i++ {
		g.units[i].Draw(Cam, CamZLevel)
	}

	cx, cy := getCursorCellPos()
	// Draw the cursor
	Cam.Surface.DrawImage(cursorImage, Cam.GetTranslation(float64(cx*cellSize), float64(cy*cellSize)))

	// Draw to screen and zoom
	Cam.Blit(screen)

	msg := fmt.Sprintf("TPS: %0.2f FPS: %0.2f\n",
		ebiten.CurrentTPS(), ebiten.CurrentFPS())

	// camXPos := int(Cam.X) / cellSize
	// camYPos := int(Cam.Y) / cellHeight

	// msg += fmt.Sprintf("CAM SIZE: %d / %d\n", camXPos, camYPos)
	// msg += fmt.Sprintf("TILES DRAWN: %d\n", g.gameMap.DrawnTileCount())
	// msg += fmt.Sprintf("CAM SCALE: %0.2f", Cam.Scale)
	msg += fmt.Sprintf("CAM Z LEVEL: %d", CamZLevel)

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
	components.SetupImages(cellSize)

	// Create camera
	Cam = camera.NewCamera(1024, 768, float64(worldWidth*cellSize/2), float64(worldHeight*cellSize/2), 0, 1)
	CamZLevel = 5

	entities.SetupResources(cellSize)

	jm := managers.NewJobManager()

	game := Game{
		gameMap:    entities.NewGameMap(worldWidth, worldHeight, cellSize, cellSize),
		jobManager: jm,
	}

	pf := components.NewPathfinder(game.gameMap)

	for i := 0; i < units; i++ {
		x := (worldWidth / 2) + rand.Intn(10) - 5
		y := (worldHeight / 2) + rand.Intn(10) - 5
		game.units = append(game.units, entities.NewUnit(x, y, pf, game.jobManager.GetNextJob))
	}

	ebiten.SetWindowSize(1024, 768)
	ebiten.SetWindowTitle("Mouse Test")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}

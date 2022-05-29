package systems

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/sedyh/mizu/pkg/engine"
	"github.com/tomknightdev/dwarven-fortresses/assets"
	"github.com/tomknightdev/dwarven-fortresses/components"
	"github.com/tomknightdev/dwarven-fortresses/entities"
	"github.com/tomknightdev/dwarven-fortresses/enums"
)

type Input struct {
}

func NewInput() *Input {
	return &Input{}
}

func (i *Input) Update(w engine.World) {
	// Handle camera update
	camera, found := w.View(components.Zoom{}, components.Position{}).Get()
	if !found {
		return
	}
	var zoom *components.Zoom
	var camPos *components.Position
	camera.Get(&zoom, &camPos)

	if inpututil.KeyPressDuration(ebiten.KeyD) > 0 && camPos.X < assets.WorldWidth*assets.CellSize {
		camPos.X += assets.CellSize / 2
	} else if inpututil.KeyPressDuration(ebiten.KeyA) > 0 && camPos.X > 0 {
		camPos.X -= assets.CellSize / 2
	}
	if inpututil.KeyPressDuration(ebiten.KeyS) > 0 && camPos.Y < assets.WorldHeight*assets.CellSize {
		camPos.Y += assets.CellSize / 2
	} else if inpututil.KeyPressDuration(ebiten.KeyW) > 0 && camPos.Y > 0 {
		camPos.Y -= assets.CellSize / 2
	}

	// Zoom the camera
	_, wy := ebiten.Wheel()
	if wy > 0 && zoom.Value < 10 {
		zoom.Value += 0.5
	} else if wy < 0 && zoom.Value > 1.1 {
		zoom.Value -= 0.5
	}

	// Adjust camera z level
	if inpututil.IsKeyJustPressed(ebiten.KeyE) && camPos.Z > 0 {
		camPos.Z--
	} else if inpututil.IsKeyJustPressed(ebiten.KeyQ) && camPos.Z < assets.WorldLevels {
		camPos.Z++
	}

	// Handle mouse update
	mouseInput, found := w.View(components.Input{}, components.Position{}).Get()
	if !found {
		return
	}
	var input *components.Input
	var mousePos *components.Position
	mouseInput.Get(&input, &mousePos)

	cx, cy := ebiten.CursorPosition()
	ww, wh := ebiten.WindowSize()
	cx = cx + (camPos.X - (ww / 2))
	cy = cy + (camPos.Y - (wh / 2))
	mousePos.X = int((float64(cx) / zoom.Value) / float64(assets.CellSize))
	mousePos.Y = int((float64(cy) / zoom.Value) / float64(assets.CellSize))

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		tiles := w.View(components.TileType{}, components.Position{}).Filter()
		var tp *components.Position

		for _, t := range tiles {
			t.Get(&tp)
			if tp.X == mousePos.X && tp.Y == mousePos.Y && tp.Z == camPos.Z {
				w.AddEntities(&entities.Job{
					Position: components.NewPosition(tp.X, tp.Y, tp.Z),
					Task:     components.NewTask(enums.MoveTo),
				})
				break
			}
		}
	}

}

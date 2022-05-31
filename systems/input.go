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
	cx, cy := ebiten.CursorPosition()
	mouseInput, found := w.View(components.Input{}, components.Position{}, components.Sprite{}).Get()
	if !found {
		return
	}
	var input *components.Input
	var mousePos *components.Position
	var mouseSprite *components.Sprite
	mouseInput.Get(&input, &mousePos, &mouseSprite)
	mouseMode := input.MouseMode

	// Check if mouse is over the gui
	guis := w.View(components.Gui{}, components.Sprite{}).Filter()
	for _, g := range guis {
		var gsp *components.Sprite
		var gui *components.Gui
		g.Get(&gsp, &gui)

		if gui.Within(cx, cy) {
			if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
				mouseMode = enums.MouseModeBuild
			}
		}
	}

	ww, wh := ebiten.WindowSize()
	cx = cx + (camPos.X - (ww / 2))
	cy = cy + (camPos.Y - (wh / 2))

	if inpututil.IsKeyJustReleased(ebiten.KeyEscape) {
		setMouseMode(input, mouseSprite, enums.MouseModeNone)
	} else if input.MouseMode != mouseMode {
		setMouseMode(input, mouseSprite, mouseMode)
	}

	mousePos.X = int((float64(cx) / zoom.Value) / float64(assets.CellSize))
	mousePos.Y = int((float64(cy) / zoom.Value) / float64(assets.CellSize))

	if mouseMode != enums.MouseModeNone && inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		tiles := w.View(components.TileType{}, components.Position{}).Filter()
		var tp *components.Position

		for _, t := range tiles {
			t.Get(&tp)
			if tp.X == mousePos.X && tp.Y == mousePos.Y && tp.Z == camPos.Z {
				w.AddEntities(&entities.Job{
					Position: components.NewPosition(tp.X, tp.Y, tp.Z),
					Task:     components.NewTask(enums.TaskTypeMoveTo),
				})
				break
			}
		}
	}

}

func setMouseMode(i *components.Input, s *components.Sprite, mouseMode enums.MouseModeEnum) {
	i.MouseMode = mouseMode

	switch mouseMode {
	case enums.MouseModeNone:
		s.Image = assets.Images["empty"]
	case enums.MouseModeGather:
		s.Image = assets.Images["cursor"]
	case enums.MouseModeBuild:
		s.Image = assets.Images["stairdown"]
	}
}

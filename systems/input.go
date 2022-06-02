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
		zoom.Value += 0.1
	} else if wy < 0 && zoom.Value > 0.2 {
		zoom.Value -= 0.1
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
	inputMode := input.InputMode

	// Check if mouse is over the gui
	guis := w.View(components.Gui{}, components.Sprite{}).Filter()
	for _, g := range guis {
		var gsp *components.Sprite
		var gui *components.Gui
		g.Get(&gsp, &gui)

		if gui.Within(cx, cy) {
			if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
				switch gui.Action {
				case enums.GuiActionStair:
					inputMode = enums.InputModeBuild
				case enums.GuiActionChop:
					inputMode = enums.InputModeChop
				}
			}
		}
	}

	ww, wh := ebiten.WindowSize()
	cx = cx + (camPos.X - (ww / 2))
	cy = cy + (camPos.Y - (wh / 2))

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonRight) {
		setMouseMode(input, mouseSprite, enums.InputModeNone)
	} else if input.InputMode != inputMode {
		setMouseMode(input, mouseSprite, inputMode)
	}

	mousePos.X = int((float64(cx) / zoom.Value) / float64(assets.CellSize))
	mousePos.Y = int((float64(cy) / zoom.Value) / float64(assets.CellSize))

	if inputMode != enums.InputModeNone && inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		var resources []engine.Entity
		switch inputMode {
		case enums.InputModeChop:
			resources = w.View(components.Choppable{}, components.Position{}).Filter()
		}

		var rPos *components.Position
		for _, r := range resources {
			r.Get(&rPos)
			if rPos.X == mousePos.X && rPos.Y == mousePos.Y && rPos.Z == camPos.Z {
				w.AddEntities(&entities.Job{
					Position: components.NewPosition(mousePos.X, mousePos.Y, camPos.Z),
					Task:     components.NewTask(inputMode, 10, r.ID()),
				})
				break
			}
		}
	}
}

func setMouseMode(i *components.Input, s *components.Sprite, mouseMode enums.InputModeEnum) {
	i.InputMode = mouseMode

	switch mouseMode {
	case enums.InputModeNone:
		s.Image = assets.Images["empty"]
	case enums.InputModeGather:
		s.Image = assets.Images["cursor"]
	case enums.InputModeBuild:
		s.Image = assets.Images["stairdown"]
	case enums.InputModeChop:
		s.Image = assets.Images["cursor"]
	}
}

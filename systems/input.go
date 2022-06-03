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
	MouseStart components.Position
	GameMap    GameMap
}

func NewInput(gameMap GameMap) *Input {
	return &Input{
		GameMap: gameMap,
	}
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
				case enums.GuiActionMine:
					inputMode = enums.InputModeMine
				}
			}
		}
	}

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonRight) {
		setMouseMode(input, mouseSprite, enums.InputModeNone)
	} else if input.InputMode != inputMode {
		setMouseMode(input, mouseSprite, inputMode)
		return
	}

	ww, wh := ebiten.WindowSize()
	cx = cx + (camPos.X - (ww / 2))
	cy = cy + (camPos.Y - (wh / 2))
	mousePos.X = int((float64(cx) / zoom.Value) / float64(assets.CellSize))
	mousePos.Y = int((float64(cy) / zoom.Value) / float64(assets.CellSize))
	mousePos.Z = camPos.Z

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		if mousePos.X < 0 || mousePos.Y < 0 {
			return
		}

		i.MouseStart = components.NewPosition(mousePos.X, mousePos.Y, camPos.Z)
	}

	if inputMode != enums.InputModeNone && inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		startX := i.MouseStart.X
		startY := i.MouseStart.Y
		endX := mousePos.X
		endY := mousePos.Y

		if endX < startX {
			startX = endX + startX
			endX = startX - endX
			startX = startX - endX
		}

		if endY < startY {
			startY = endY + startY
			endY = startY - endY
			startY = startY - endY
		}

		// var resources []engine.Entity
		switch inputMode {
		case enums.InputModeChop:
			resources := w.View(components.Choppable{}, components.Position{}).Filter()
			var rPos *components.Position
			for _, r := range resources {
				r.Get(&rPos)
				for mx := startX; mx <= endX; mx++ {
					for my := startY; my <= endY; my++ {
						if rPos.X == mx && rPos.Y == my && rPos.Z == camPos.Z {
							w.AddEntities(&entities.Job{
								Position: components.NewPosition(mx, my, camPos.Z),
								Task:     components.NewTask(inputMode, 10, r.ID()),
							})
						}
					}
				}

			}
		case enums.InputModeMine:
			for mx := startX; mx <= endX; mx++ {
				for my := startY; my <= endY; my++ {
					index := i.GameMap.GetTileByTypeIndexFromPos(enums.TileTypeRock, components.NewPosition(mx, my, camPos.Z))
					if index < 0 {
						continue
					}

					w.AddEntities(&entities.Job{
						Position: components.NewPosition(mx, my, camPos.Z),
						Task:     components.NewTask(inputMode, 10, 0),
					})
				}
			}
		case enums.InputModeBuild:
			w.AddEntities(&entities.Job{
				Position: components.NewPosition(endX, endY, camPos.Z),
				Task:     components.NewTask(inputMode, 10, 0),
			})
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
	case enums.InputModeMine:
		s.Image = assets.Images["pickaxe"]
	}
}

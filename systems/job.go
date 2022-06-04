package systems

import (
	"log"

	"github.com/sedyh/mizu/pkg/engine"
	"github.com/tomknightdev/dwarven-fortresses/components"
	"github.com/tomknightdev/dwarven-fortresses/entities"
	"github.com/tomknightdev/dwarven-fortresses/enums"
	"github.com/tomknightdev/dwarven-fortresses/helpers"
)

type Job struct {
}

func NewJob() *Job {
	return &Job{}
}

func (j *Job) Update(w engine.World) {
	var inputSingleton *components.InputSingleton
	is, found := w.View(components.InputSingleton{}).Get()
	if !found {
		panic("input singleton was not found")
	}
	is.Get(&inputSingleton)

	if inputSingleton.InputMode == enums.InputModeNone {
		return
	}

	if len(inputSingleton.SelectedTiles) > 0 {
		switch inputSingleton.InputMode {
		case enums.InputModeChop:
			resources := w.View(components.Choppable{}, components.Position{}).Filter()
			var rPos *components.Position
			for _, r := range resources {
				r.Get(&rPos)
				for _, st := range inputSingleton.SelectedTiles {
					if rPos.X == st.X && rPos.Y == st.Y && rPos.Z == st.Z {
						w.AddEntities(&entities.Job{
							Position: components.NewPosition(st.X, st.Y, st.Z),
							Task:     components.NewTask(enums.InputModeChop, 10, r.ID()),
						})
					}
				}

			}
		case enums.InputModeMine:
			gms, found := w.View(components.GameMapSingleton{}).Get()
			if !found {
				panic("game map singleton not found")
			}
			var gmComp *components.GameMapSingleton
			gms.Get(&gmComp)

			for _, st := range inputSingleton.SelectedTiles {
				index, err := helpers.GetTileByTypeIndexFromPos(components.NewPosition(st.X, st.Y, st.Z), gmComp.TilesByType[enums.TileTypeRock])
				if err != nil {
					log.Println(err)
					continue
				}

				if index < 0 {
					continue
				}

				w.AddEntities(&entities.Job{
					Position: components.NewPosition(st.X, st.Y, st.Z),
					Task:     components.NewTask(enums.InputModeMine, 10, 0),
				})
			}

		case enums.InputModeBuild:
			w.AddEntities(&entities.Job{
				Position: components.NewPosition(inputSingleton.SelectedTiles[0].X, inputSingleton.SelectedTiles[0].Y, inputSingleton.SelectedTiles[0].Z),
				Task:     components.NewTask(enums.InputModeBuild, 10, 0),
			})
		}
	}
}

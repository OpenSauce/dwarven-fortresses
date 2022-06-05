package systems

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sedyh/mizu/pkg/engine"
	"github.com/tomknightdev/dwarven-fortresses/components"
	"github.com/tomknightdev/dwarven-fortresses/enums"
	"github.com/tomknightdev/dwarven-fortresses/helpers"
)

type Actor struct {
}

func NewActor() *Actor {
	return &Actor{}
}

func (a *Actor) Update(w engine.World) {
	jobs := w.View(components.Task{}, components.Position{}).Filter()

	actors := w.View(components.Worker{}, components.Move{}, components.Position{})
	actors.Each(func(e engine.Entity) {
		var worker *components.Worker
		var move *components.Move
		var pos *components.Position

		e.Get(&worker, &move, &pos)

		if !worker.HasJob {
			if len(jobs) == 0 {
				return
			}

			var jp *components.Position
			var task *components.Task

			for _, job := range jobs {
				job.Get(&jp, &task)
				if task.Claimed {
					continue
				}

				worker.HasJob = true
				worker.JobID = job.ID()
				worker.InputModeEnum = task.InputModeEnum

				move.Adjacent = true

				move.X = jp.X
				move.Y = jp.Y
				move.Z = jp.Z
				task.Claimed = true
				break
			}
		} else if move.Arrived {
			if move.CurrentEnergy < move.TotalEnergy {
				move.CurrentEnergy++
				return
			}

			job, found := w.GetEntity(worker.JobID)
			if !found {
				log.Println("arrived at location but job not found")
				worker.HasJob = false
				worker.InputModeEnum = enums.InputModeNone
				return
			}

			var task *components.Task
			job.Get(&task)
			if task.ActionsComplete < task.RequiredActions {
				task.ActionsComplete++
				move.CurrentEnergy = 0
				return
			}

			task.CompleteTask()
			worker.HasJob = false
			worker.JobID = 0
			worker.InputModeEnum = enums.InputModeNone
		} else {
			// Check for job cancellation
			_, found := w.GetEntity(worker.JobID)
			if !found {
				worker.HasJob = false
				worker.JobID = 0
				worker.InputModeEnum = enums.InputModeBuild
				move.X = pos.X
				move.Y = pos.Y
				move.Z = pos.Z

				move.Arrived = true
			}

		}
	})
}

func (a *Actor) Draw(w engine.World, screen *ebiten.Image) {
	ents := w.View(components.Move{}, components.Position{}, components.Sprite{})
	var p *components.Position
	var s *components.Sprite

	ents.Each(func(e engine.Entity) {
		e.Get(&p, &s)

		helpers.DrawImage(w, screen, *p, s.Image)
	})
}

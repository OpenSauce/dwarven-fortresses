package systems

import (
	"github.com/sedyh/mizu/pkg/engine"
	"github.com/tomknightdev/dwarven-fortresses/components"
	"github.com/tomknightdev/dwarven-fortresses/enums"
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
				panic("arrived at location but job not found")
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
			worker.InputModeEnum = enums.InputModeNone
		}
	})

}

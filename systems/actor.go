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
	var jobsToRemove []engine.Entity
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
				worker.TaskTypeEnum = task.TaskTypeEnum
				move.X = jp.X
				move.Y = jp.Y
				move.Z = jp.Z
				task.Claimed = true
				break
			}
		} else if move.Arrived {
			job, found := w.GetEntity(worker.JobID)
			if !found {
				panic("arrived at location but job not found")
			}

			worker.HasJob = false
			worker.TaskTypeEnum = enums.TaskTypeNone
			jobsToRemove = append(jobsToRemove, job)
		}
	})

	for _, job := range jobsToRemove {
		w.RemoveEntity(job)
	}
}

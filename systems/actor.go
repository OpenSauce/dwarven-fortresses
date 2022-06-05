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
	jobs := w.View(components.Job{}).Filter()

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

			var tasks *components.Job

			for _, job := range jobs {
				job.Get(&tasks)
				if tasks.ClaimedByID > 0 {
					continue
				}

				worker.HasJob = true
				worker.JobID = job.ID()

				currentTask := tasks.Tasks[0]

				move.Adjacent = true
				if currentTask.TaskTypeEnum == enums.TaskTypePickUp || currentTask.TaskTypeEnum == enums.TaskTypeDrop {
					move.Adjacent = false
				}

				move.X = currentTask.Position.X
				move.Y = currentTask.Position.Y
				move.Z = currentTask.Position.Z
				tasks.ClaimedByID = e.ID()
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
				return
			}

			var tasks *components.Job
			job.Get(&tasks)
			currentTask := tasks.Tasks[0]

			if currentTask.ActionsComplete < currentTask.RequiredActions {
				currentTask.ActionsComplete++
				move.CurrentEnergy = 0
				return
			}

			currentTask.CompleteTask()
			if len(tasks.Tasks) > 1 {
				task := tasks.Tasks[1]
				move.Adjacent = true
				if task.TaskTypeEnum == enums.TaskTypePickUp || task.TaskTypeEnum == enums.TaskTypeDrop || task.TaskTypeEnum == enums.TaskTypeAddToStockpile {
					move.Adjacent = false
				}

				move.X = task.Position.X
				move.Y = task.Position.Y
				move.Z = task.Position.Z
				tasks.ClaimedByID = e.ID()
				move.Arrived = false
				return
			}

			worker.HasJob = false
			worker.JobID = 0
		} else {
			// Check for job cancellation
			_, found := w.GetEntity(worker.JobID)
			if !found {
				worker.HasJob = false
				worker.JobID = 0
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

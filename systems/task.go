package systems

import (
	"github.com/sedyh/mizu/pkg/engine"
	"github.com/tomknightdev/dwarven-fortresses/assets"
	"github.com/tomknightdev/dwarven-fortresses/components"
	"github.com/tomknightdev/dwarven-fortresses/entities"
	"github.com/tomknightdev/dwarven-fortresses/enums"
)

type Task struct {
}

func NewTask() *Task {
	return &Task{}
}

func (t *Task) Update(w engine.World) {
	jobs := w.View(components.Task{}).Filter()
	var entitiesToRemove []engine.Entity

	var task *components.Task
	for _, job := range jobs {
		job.Get(&task)

		if task.Completed {
			ent, ok := w.GetEntity(task.EntityID)
			if !ok {
				panic("entity not found")
			}

			switch task.InputModeEnum {
			case enums.InputModeChop:
				var drop *components.Drops
				var pos *components.Position
				ent.Get(&drop, &pos)

				for i := 0; i < drop.DropCount; i++ {
					w.AddEntities(&entities.Resource{
						Position: *pos,
						Sprite:   components.NewSprite(assets.Images["log0"]),
						Resource: components.NewResource(),
					})
				}

				entitiesToRemove = append(entitiesToRemove, ent)
			}

			entitiesToRemove = append(entitiesToRemove, job)
		}
	}

	for _, job := range entitiesToRemove {
		w.RemoveEntity(job)
	}
}

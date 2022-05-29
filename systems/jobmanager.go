package systems

import (
	"github.com/sedyh/mizu/pkg/engine"
	"github.com/tomknightdev/dwarven-fortresses/components"
)

type JobManager struct {
}

func (jm *JobManager) Update(w engine.World) {
	view := w.View(components.Task{})
	view.Each(func(e engine.Entity) {
		var job *components.Task

		e.Get(&job)

	})
}

package components

import "github.com/tomknightdev/dwarven-fortresses/enums"

type Worker struct {
	HasJob bool
	enums.TaskTypeEnum
	JobID int
}

func NewWorker() Worker {
	return Worker{}
}

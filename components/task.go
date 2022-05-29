package components

import "github.com/tomknightdev/dwarven-fortresses/enums"

type Task struct {
	enums.TaskTypeEnum
	Claimed bool
}

func NewTask(tt enums.TaskTypeEnum) Task {
	return Task{
		TaskTypeEnum: tt,
	}
}

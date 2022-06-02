package components

import "github.com/tomknightdev/dwarven-fortresses/enums"

type Task struct {
	enums.InputModeEnum
	Claimed         bool
	RequiredActions int
	ActionsComplete int
	EntityID        int
	Completed       bool
}

func NewTask(inputModeEnum enums.InputModeEnum, requiredActions, entityId int) Task {
	return Task{
		InputModeEnum:   inputModeEnum,
		RequiredActions: requiredActions,
		EntityID:        entityId,
	}
}

func (t *Task) CompleteTask() {
	t.Completed = true
}

package components

type Job struct {
	Tasks       []*Task
	ClaimedByID int
	EntityID    int
}

func NewJob(id int, tasks ...*Task) Job {
	return Job{
		EntityID: id,
		Tasks:    tasks,
	}
}

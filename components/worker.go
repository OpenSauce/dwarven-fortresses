package components

type Worker struct {
	HasJob bool
	JobID  int
}

func NewWorker() Worker {
	return Worker{}
}

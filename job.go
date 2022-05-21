package main

type Job struct {
	XPos, YPos int
}

var (
	jobs = []*Job{}
)

func CreateJob(XPos, YPos int) {
	jobs = append(jobs, &Job{
		XPos,
		YPos,
	})
}

func GetNextJob() *Job {
	if len(jobs) == 0 {
		return nil
	}

	j := jobs[0]
	jobs = jobs[1:]
	return j
}

package managers

import (
	"github.com/OpenSauce/paths"
)

type JobType int

const (
	Gather JobType = iota
	StairDown
	StairUp
)

var (
	jobs = []*Job{}
)

type JobManager struct {
}

func NewJobManager() *JobManager {
	j := JobManager{}

	return &j
}

func (j *Job) CompleteJob() {
	if j.JobType == Gather {
		j.tile.Gathered()
		// j.tile.cell.Cost = float64(Dirt)
	} else if j.JobType == StairDown {
		j.tile.SetType("stairDown")
	} else if j.JobType == StairUp {
		j.tile.SetType("stairUp")
	}
}

func (jm *JobManager) CreateJob(c *paths.Cell, t *Tile, jt JobType) {
	j := Job{
		c,
		t,
		jt,
	}

	j.tile.resource.queued = true
	jobs = append(jobs, &j)
}

func (jm *JobManager) GetNextJob() *Job {
	if len(jobs) == 0 {
		return nil
	}

	j := jobs[0]
	jobs = jobs[1:]
	return j
}
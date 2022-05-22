package main

import (
	"github.com/tomknightdev/paths"
)

type JobType int

const (
	Gather JobType = iota
)

var (
	jobs = []*Job{}
)

type Job struct {
	cell *paths.Cell
	tile *Tile
	JobType
}

func (j *Job) CompleteJob() {
	if j.JobType == Gather {
		j.tile.Gethered()
	}
}

func CreateJob(c *paths.Cell, t *Tile) {
	j := Job{
		c,
		t,
		Gather,
	}

	j.tile.resource.queued = true
	jobs = append(jobs, &j)
}

func GetNextJob() *Job {
	if len(jobs) == 0 {
		return nil
	}

	j := jobs[0]
	jobs = jobs[1:]
	return j
}

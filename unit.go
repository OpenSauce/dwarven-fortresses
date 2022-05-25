package main

import (
	"image"
	"time"

	_ "embed"

	"github.com/OpenSauce/paths"
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	unitImage *ebiten.Image
)

type Pathfinder interface {
	GetPath(int, int, int, int, int, int) []struct {
		*paths.Path
		ZTraversable
	}
	GetMapDimensions() (int, int)
	IsWalkable(int, int, int) bool
	GetCellCost(int, int, int) int
	SwitchWalkable(int, int, int)
}

// type Jobfinder interface {
// 	GetNextJob() *Job
// }

type Unit struct {
	Running    bool
	XPos, YPos int
	Pathfinder
	jobfinder                  func() *Job
	TurnSpeed, CurrentTurnTime int
	image                      *ebiten.Image
	currentPaths               []struct {
		*paths.Path
		ZTraversable
	}
	currentJob        *Job
	maxEnergy, Energy int
	zLevel            int
}

func init() {
	unitImage = TilesetImage.SubImage(image.Rect(25*cellWidth, 0*cellHeight, 26*cellWidth, 1*cellHeight)).(*ebiten.Image)
}

func NewUnit(startX, startY int, pf Pathfinder, jf func() *Job) *Unit {
	u := Unit{
		XPos:       startX,
		YPos:       startY,
		Pathfinder: pf,
		jobfinder:  jf,
		TurnSpeed:  200,
		image:      unitImage,
		maxEnergy:  1000,
		Energy:     1000,
		zLevel:     5,
	}

	go u.gameLoop()

	return &u
}

func (u *Unit) gameLoop() {
	tick := time.Tick(time.Duration(u.TurnSpeed) * time.Millisecond)

	for range tick {
		if !u.Running {
			continue
		}
		u.Update()
	}
}

func (u *Unit) Update() error {
	if u.Energy < u.maxEnergy {
		u.Energy += 5
	}

	if u.currentJob == nil {
		if u.Energy > u.maxEnergy/2 {
			u.currentJob = GetNextJob()
		}
	}

	if u.currentJob != nil {
		if u.Move() {
			u.Work()
		}
	}

	return nil
}

func (u *Unit) Work() {
	time.Sleep(time.Second * 2)
	u.Energy -= 50
	u.currentJob.CompleteJob()
	u.currentJob = nil
}

// Move returns true if at target cell
func (u *Unit) Move() bool {
	// Check if at destination cell on correct z level
	if u.zLevel == u.currentJob.tile.zLevel && u.XPos == u.currentJob.cell.X && u.YPos == u.currentJob.cell.Y {
		return true
	}

	// If no paths in place, or current path has no next cell, find new paths
	if len(u.currentPaths) == 0 || u.currentPaths[0].Next() == nil {
		u.currentPaths = u.Pathfinder.GetPath(u.XPos, u.YPos, u.zLevel, u.currentJob.cell.X, u.currentJob.cell.Y, u.currentJob.tile.zLevel)

		// No paths found
		// TODO: return error as job cannot be reached
		if len(u.currentPaths) == 0 {
			return false
		}
	}

	// If reached end of current path and no more paths, reached destination
	if u.currentPaths[0].AtEnd() && u.atPathEnd() {
		return true
	}

	next := u.currentPaths[0].Next()

	if next != nil {
		// If the next cell is no longer walkable, find new paths
		if !next.Walkable {
			u.currentPaths = u.Pathfinder.GetPath(u.XPos, u.YPos, u.zLevel, u.currentJob.cell.X, u.currentJob.cell.Y, u.currentJob.tile.zLevel)

			// No paths found
			// TODO: return error as job cannot be reached
			if len(u.currentPaths) == 0 {
				return false
			}

			next = u.currentPaths[0].Next()
		}

		// Set unit position
		u.XPos = next.X
		u.YPos = next.Y
		u.Energy -= 10

		u.currentPaths[0].Advance()

		// If reached end of current path and no more paths, reached destination
		if u.currentPaths[0].AtEnd() && u.atPathEnd() {
			return true
		}
	}

	return false
}

func (u *Unit) atPathEnd() bool {
	if u.currentPaths[0].ZTraversable == UP {
		u.zLevel++
	} else if u.currentPaths[0].ZTraversable == DOWN {
		u.zLevel--
	}

	if len(u.currentPaths) == 1 {
		u.currentPaths = []struct {
			*paths.Path
			ZTraversable
		}{}
		return true
	} else {
		u.currentPaths = u.currentPaths[1:]
		return false
	}
}

func (u *Unit) Draw(screen *ebiten.Image) {
	// Draw the unit
	Cam.Surface.DrawImage(u.image, Cam.GetTranslation(float64(u.XPos*cellWidth), float64(u.YPos*cellHeight)))
}

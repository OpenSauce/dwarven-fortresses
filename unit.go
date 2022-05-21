package main

import (
	"image"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tomknightdev/paths"
)

type Pathfinder interface {
	GetPath(int, int, int, int) *paths.Path
	GetMapDimensions() (int, int)
	IsWalkable(int, int) bool
	GetCellCost(int, int) int
	SwitchWalkable(int, int)
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
	currentPath                *paths.Path
	currentJob                 *Job
}

func NewUnit(startX, startY int, pf Pathfinder, jf func() *Job) *Unit {
	u := Unit{
		XPos:       startX,
		YPos:       startY,
		Pathfinder: pf,
		jobfinder:  jf,
		TurnSpeed:  10,
		image: tilesImage.SubImage(image.Rectangle{
			Min: image.Pt(2*cellWidth, 3*cellHeight),
			Max: image.Pt(3*cellWidth, 4*cellHeight),
		}).(*ebiten.Image),
	}

	go u.gameLoop()

	return &u
}

func (u *Unit) gameLoop() {
	tick := time.Tick(100 * time.Millisecond)

	for range tick {
		if !u.Running {
			continue
		}
		u.Update()
	}
}

func (u *Unit) Update() error {
	// if u.CurrentTurnTime < u.TurnSpeed+u.GetCellCost(u.XPos, u.YPos) {
	// 	u.CurrentTurnTime++
	// 	return nil
	// }
	// u.CurrentTurnTime = 0

	if u.currentJob == nil {
		u.currentJob = GetNextJob()
	}

	if u.currentJob != nil {
		if u.Move() {
			u.Work()
		}
	}

	return nil
}

func (u *Unit) Work() {
	time.Sleep(time.Second * 5)

	u.SwitchWalkable(u.currentJob.XPos, u.currentJob.YPos)
	u.currentJob = nil
}

// Move returns true if at target cell
func (u *Unit) Move() bool {
	if (u.XPos == u.currentJob.XPos || u.XPos == u.currentJob.XPos-1 || u.XPos == u.currentJob.XPos+1) &&
		(u.YPos == u.currentJob.YPos || u.YPos == u.currentJob.YPos-1 || u.YPos == u.currentJob.YPos+1) {
		return true
	}

	if u.currentPath == nil || u.currentPath.Next() == nil {
		u.currentPath = u.Pathfinder.GetPath(u.XPos, u.YPos, u.currentJob.XPos, u.currentJob.YPos)

		if u.currentPath == nil {
			return false
		}
	}

	next := u.currentPath.Next()

	if next != nil {
		if !next.Walkable {
			u.currentPath = u.Pathfinder.GetPath(u.XPos, u.YPos, u.currentJob.XPos, u.currentJob.YPos)
			if u.currentPath == nil {
				return false
			}

			next = u.currentPath.Next()
		}

		u.XPos = next.X
		u.YPos = next.Y

		u.currentPath.Advance()
	}

	return false
}

func (u *Unit) Draw(screen *ebiten.Image) {

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(u.XPos*cellWidth), float64(u.YPos*cellHeight))
	screen.DrawImage(u.image, op)
}

// func (u *Unit) getNewPos() {
// 	width, height := u.Pathfinder.GetMapDimensions()
// 	x := rand.Intn(width - 1)
// 	y := rand.Intn(height - 1)

// 	// if u.Pathfinder.IsWalkable(x, y) {
// 	u.XTar = x
// 	u.YTar = y
// 	// }
// }

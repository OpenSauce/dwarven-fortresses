package main

import (
	"bytes"
	"image/png"
	"log"
	"time"

	_ "embed"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tomknightdev/paths"
)

var (
	//go:embed resources/creatures/player.png
	player_sheet []byte
	playerImage  *ebiten.Image
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
	maxEnergy, Energy          int
}

func init() {
	img, err := png.Decode(bytes.NewReader(player_sheet))
	if err != nil {
		log.Fatal(err)
	}
	playerImage = ebiten.NewImageFromImage(img)
}

func NewUnit(startX, startY int, pf Pathfinder, jf func() *Job) *Unit {
	u := Unit{
		XPos:       startX,
		YPos:       startY,
		Pathfinder: pf,
		jobfinder:  jf,
		TurnSpeed:  200,
		image:      playerImage,
		maxEnergy:  1000,
		Energy:     1000,
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
	u.Energy -= 100
	u.currentJob.CompleteJob()
	u.currentJob = nil
}

// Move returns true if at target cell
func (u *Unit) Move() bool {
	if (u.XPos == u.currentJob.cell.X || u.XPos == u.currentJob.cell.X-1 || u.XPos == u.currentJob.cell.X+1) &&
		(u.YPos == u.currentJob.cell.Y || u.YPos == u.currentJob.cell.Y-1 || u.YPos == u.currentJob.cell.Y+1) {
		return true
	}

	if u.currentPath == nil || u.currentPath.Next() == nil {
		u.currentPath = u.Pathfinder.GetPath(u.XPos, u.YPos, u.currentJob.cell.X, u.currentJob.cell.Y)

		if u.currentPath == nil {
			return false
		}
	}

	next := u.currentPath.Next()

	if next != nil {
		if !next.Walkable {
			u.currentPath = u.Pathfinder.GetPath(u.XPos, u.YPos, u.currentJob.cell.X, u.currentJob.cell.Y)
			if u.currentPath == nil {
				return false
			}

			next = u.currentPath.Next()
		}

		u.XPos = next.X
		u.YPos = next.Y
		u.Energy -= 10

		u.currentPath.Advance()
	}

	return false
}

func (u *Unit) Draw(screen *ebiten.Image) {
	// Draw the unit
	Cam.Surface.DrawImage(u.image, Cam.GetTranslation(float64(u.XPos*cellWidth), float64(u.YPos*cellHeight)))
}

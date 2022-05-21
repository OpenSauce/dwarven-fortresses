package main

import (
	"image"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/paths"
)

type Pathfinder interface {
	GetPath(int, int, int, int) *paths.Path
	GetMapDimensions() (int, int)
	IsWalkable(int, int) bool
	GetCellCost(int, int) int
}

type Unit struct {
	XPos, YPos int
	XTar, YTar int
	Pathfinder
	TurnSpeed, CurrentTurnTime int
	image                      *ebiten.Image
	currentPath                *paths.Path
}

func NewUnit(startX, startY int, pf Pathfinder) *Unit {
	u := Unit{
		startX,
		startY,
		startX,
		startY,
		pf,
		10,
		0,
		tilesImage.SubImage(image.Rectangle{
			Min: image.Pt(2*cellWidth, 3*cellHeight),
			Max: image.Pt(3*cellWidth, 4*cellHeight),
		}).(*ebiten.Image),
		nil,
	}

	return &u
}

func (u *Unit) Update() error {
	if u.CurrentTurnTime < u.TurnSpeed+u.GetCellCost(u.XPos, u.YPos) {
		u.CurrentTurnTime++
		return nil
	}
	u.CurrentTurnTime = 0

	if (u.XTar == u.XPos && u.YTar == u.YPos) || u.currentPath == nil {
		u.CurrentTurnTime -= 50
		u.getNewPos()
		u.currentPath = u.Pathfinder.GetPath(u.XPos, u.YPos, u.XTar, u.YTar)

		if u.currentPath == nil {
			return nil
		}
	}

	next := u.currentPath.Next()

	if next != nil {
		if !next.Walkable {
			u.currentPath = u.Pathfinder.GetPath(u.XPos, u.YPos, u.XTar, u.YTar)
			if u.currentPath == nil {
				return nil
			}

			next = u.currentPath.Next()
		}

		u.XPos = next.X
		u.YPos = next.Y

		u.currentPath.Advance()
	}

	return nil
}

func (u *Unit) Draw(screen *ebiten.Image) {

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(u.XPos*cellWidth), float64(u.YPos*cellHeight))
	screen.DrawImage(u.image, op)
}

func (u *Unit) getNewPos() {
	width, height := u.Pathfinder.GetMapDimensions()
	x := rand.Intn(width - 1)
	y := rand.Intn(height - 1)

	if u.Pathfinder.IsWalkable(x, y) {
		u.XTar = x
		u.YTar = y
	}
}

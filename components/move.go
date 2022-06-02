package components

import "github.com/OpenSauce/paths"

type Move struct {
	X, Y, Z       int
	Adjacent      bool // Defines whether to move to an adjacent tile or the actual x,y tile
	Arrived       bool
	CurrentEnergy int
	TotalEnergy   int
	GettingRoute  bool
	CurrentPaths  []struct {
		*paths.Path
		Level int
	}
}

func NewMove(x, y, z int) Move {
	return Move{
		X:             x,
		Y:             y,
		Z:             z,
		CurrentEnergy: 10,
		TotalEnergy:   10,
	}
}

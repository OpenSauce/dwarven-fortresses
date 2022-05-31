package components

type Move struct {
	X, Y, Z       int
	Arrived       bool
	CurrentEnergy int
	TotalEnergy   int
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

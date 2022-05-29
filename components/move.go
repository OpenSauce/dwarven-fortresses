package components

type Move struct {
	X, Y, Z int
	Arrived bool
}

func NewMove(x, y, z int) Move {
	return Move{
		X: x,
		Y: y,
		Z: z,
	}
}

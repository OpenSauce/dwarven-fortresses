package components

type Position struct {
	X, Y, Z int
}

func NewPosition(x, y, z int) Position {
	return Position{x, y, z}
}

package components

type Mouse struct {
	MouseStart        Position
	LeftClickedTiles  []Position
	RightClickedTiles []Position
}

func NewMouse() Mouse {
	return Mouse{}
}

package components

type Mouse struct {
	MouseStart    Position
	SelectedTiles []Position
}

func NewMouse() Mouse {
	return Mouse{}
}

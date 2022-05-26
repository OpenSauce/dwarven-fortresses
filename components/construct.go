package components

type Construct struct {
	Width, Height, Levels int
}

func NewConstruct(width, height, levels int) Construct {
	return Construct{width, height, levels}
}

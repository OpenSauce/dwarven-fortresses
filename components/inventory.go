package components

type Inventory struct {
	Items  []int
	Weight int
}

func NewInventory() Inventory {
	return Inventory{}
}

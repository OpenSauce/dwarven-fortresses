package components

import "github.com/tomknightdev/dwarven-fortresses/enums"

type Item struct {
	Haulable    bool
	InStockpile bool
	Claimed     bool
	Weight      int
	ItemType    enums.ItemTypeEnum
}

func NewItem(haulable bool, weight int, itemType enums.ItemTypeEnum) Item {
	return Item{
		Haulable: haulable,
		Weight:   weight,
		ItemType: itemType,
	}
}

package components

import "github.com/tomknightdev/dwarven-fortresses/enums"

type Drops struct {
	DropType  enums.DropTypeEnum
	DropCount int
}

func NewDrops(dropType enums.DropTypeEnum, dropCount int) Drops {
	return Drops{
		DropType:  dropType,
		DropCount: dropCount,
	}
}

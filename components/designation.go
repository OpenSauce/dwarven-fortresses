package components

import "github.com/tomknightdev/dwarven-fortresses/enums"

type Designation struct {
	DesignationType enums.DesignationTypeEnum
	ItemType        enums.ItemTypeEnum
	MaxItems        int
}

func NewDesignation(designationType enums.DesignationTypeEnum) Designation {
	return Designation{
		DesignationType: designationType,
		MaxItems:        10,
	}
}

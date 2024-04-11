package entity

import "trekkstay/core"

type DestinationEntity struct {
	core.BaseEntity `json:",inline"`
	Name            string `json:"name" gorm:"not null;"`
	Code            string `json:"code" gorm:"uniqueIndex;not null;"`
	Unit            string `json:"unit" gorm:"not null;"`
}

func (DestinationEntity) TableName() string {
	return "destinations"
}

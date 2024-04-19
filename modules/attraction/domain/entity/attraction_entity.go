package entity

import "trekkstay/core"

type AttractionEntity struct {
	core.BaseEntity `json:",inline"`
	Name            string `json:"name" gorm:"not null;"`
	ProvinceCode    string `json:"-" gorm:"not null;"`
	DistrictCode    string `json:"-" gorm:"not null;"`
	WardCode        string `json:"-" gorm:"not null;"`
	// Relations
	Province ProvinceEntity `json:"province" gorm:"foreignKey:ProvinceCode;references:Code"`
	District DistrictEntity `json:"district" gorm:"foreignKey:DistrictCode;references:Code"`
	Ward     WardEntity     `json:"ward" gorm:"foreignKey:WardCode;references:Code"`
}

func (AttractionEntity) TableName() string {
	return "attractions"
}

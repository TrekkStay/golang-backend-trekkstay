package entity

type DistrictEntity struct {
	Code         string `json:"code" gorm:"column:code;primaryKey"`
	NameVi       string `json:"name_vi" gorm:"column:name_vi"`
	NameEn       string `json:"name_en" gorm:"column:name_en"`
	FullNameVi   string `json:"full_name_vi" gorm:"column:full_name_vi"`
	FullNameEn   string `json:"full_name_en" gorm:"column:full_name_en"`
	ProvinceCode string `json:"-" gorm:"column:province_code"`
}

func (DistrictEntity) TableName() string {
	return "districts"
}

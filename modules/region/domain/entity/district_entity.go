package entity

type DistrictEntity struct {
	Code       string `json:"code" gorm:"column:code;primaryKey"`
	FullNameVi string `json:"full_name_vi" gorm:"column:full_name_vi"`
	FullNameEn string `json:"full_name_en" gorm:"column:full_name_en"`
}

func (DistrictEntity) TableName() string {
	return "districts"
}

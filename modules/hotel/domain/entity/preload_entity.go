package entity

type ProvinceEntity struct {
	Code       string `json:"-" gorm:"column:code;primaryKey"`
	FullNameVi string `json:"full_name_vi" gorm:"column:full_name_vi"`
	FullNameEn string `json:"full_name_en" gorm:"column:full_name_en"`
}

func (ProvinceEntity) TableName() string {
	return "provinces"
}

type DistrictEntity struct {
	Code       string `json:"-" gorm:"column:code;primaryKey"`
	FullNameVi string `json:"full_name_vi" gorm:"column:full_name_vi"`
	FullNameEn string `json:"full_name_en" gorm:"column:full_name_en"`
}

func (DistrictEntity) TableName() string {
	return "districts"
}

type WardEntity struct {
	Code       string `json:"-" gorm:"column:code;primaryKey"`
	FullNameVi string `json:"full_name_vi" gorm:"column:full_name_vi"`
	FullNameEn string `json:"full_name_en" gorm:"column:full_name_en"`
}

func (WardEntity) TableName() string {
	return "wards"
}

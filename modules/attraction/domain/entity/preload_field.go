package entity

type ProvinceEntity struct {
	Code       string `json:"code" gorm:"column:code;primary_key:true"`
	NameVi     string `json:"name_vi" gorm:"column:name_vi"`
	NameEn     string `json:"name_en" gorm:"column:name_en"`
	FullNameVi string `json:"full_name_vi" gorm:"column:full_name_vi"`
	FullNameEn string `json:"full_name_en" gorm:"column:full_name_en"`
}

type DistrictEntity struct {
	Code       string `json:"code" gorm:"column:code;primary_key:true"`
	NameVi     string `json:"name_vi" gorm:"column:name_vi"`
	NameEn     string `json:"name_en" gorm:"column:name_en"`
	FullNameVi string `json:"full_name_vi" gorm:"column:full_name_vi"`
	FullNameEn string `json:"full_name_en" gorm:"column:full_name_en"`
}

type WardEntity struct {
	Code       string `json:"code" gorm:"column:code;primary_key:true"`
	NameVi     string `json:"name_vi" gorm:"column:name_vi"`
	NameEn     string `json:"name_en" gorm:"column:name_en"`
	FullNameVi string `json:"full_name_vi" gorm:"column:full_name_vi"`
	FullNameEn string `json:"full_name_en" gorm:"column:full_name_en"`
}

func (WardEntity) TableName() string {
	return "wards"
}

func (DistrictEntity) TableName() string {
	return "districts"
}

func (ProvinceEntity) TableName() string {
	return "provinces"
}

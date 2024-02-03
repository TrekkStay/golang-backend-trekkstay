package entity

type ProvinceEntity struct {
	Code       string `json:"-" gorm:"column:code;primary_key:true"`
	NameVi     string `json:"name_vi" gorm:"column:name_vi"`
	NameEn     string `json:"name_en" gorm:"column:name_en"`
	FullNameVi string `json:"full_name_vi" gorm:"column:full_name_vi"`
	FullNameEn string `json:"full_name_en" gorm:"column:full_name_en"`
}

type DistrictEntity struct {
	Code       string `json:"-" gorm:"column:code;primary_key:true"`
	NameVi     string `json:"name_vi" gorm:"column:name_vi"`
	NameEn     string `json:"name_en" gorm:"column:name_en"`
	FullNameVi string `json:"full_name_vi" gorm:"column:full_name_vi"`
	FullNameEn string `json:"full_name_en" gorm:"column:full_name_en"`
}

type WardEntity struct {
	Code       string `json:"-" gorm:"column:code;primary_key:true"`
	NameVi     string `json:"name_vi" gorm:"column:name_vi"`
	NameEn     string `json:"name_en" gorm:"column:name_en"`
	FullNameVi string `json:"full_name_vi" gorm:"column:full_name_vi"`
	FullNameEn string `json:"full_name_en" gorm:"column:full_name_en"`
}

type HotelEmployeeEntity struct {
	ID         string `json:"id" gorm:"column:id;"`
	HotelID    string `json:"hotel_id" gorm:"column:hotel_id;"`
	FullName   string `json:"full_name" gorm:"column:full_name;"`
	Email      string `json:"email" gorm:"column:email;"`
	Phone      string `json:"phone" gorm:"column:phone;"`
	Contract   string `json:"contract" gorm:"column:contract;"`
	BaseSalary int    `json:"base_salary" gorm:"column:base_salary;"`
	Role       string `json:"role" gorm:"column:role;"`
	Status     string `json:"status" gorm:"column:status;"`
}

type HotelRoomEntity struct {
	ID            string    `json:"id" gorm:"column:id;"`
	HotelID       string    `json:"hotel_id" gorm:"column:hotel_id;"`
	Type          string    `json:"type" gorm:"column:type;"`
	Description   string    `json:"description" gorm:"column:description;"`
	Quantity      int       `json:"quantity" gorm:"column:quantity;"`
	DiscountRate  int       `json:"discount_rate" gorm:"column:discount_rate;"`
	OriginalPrice int       `json:"original_price" gorm:"column:original_price;"`
	Videos        MediaJSON `json:"videos" gorm:"type:jsonb;"`
	Images        MediaJSON `json:"images" gorm:"type:jsonb;"`
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

func (HotelEmployeeEntity) TableName() string {
	return "hotel_employees"
}

func (HotelRoomEntity) TableName() string {
	return "hotel_rooms"
}

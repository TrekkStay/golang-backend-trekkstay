package entity

type RoomFacilityEntity struct {
	RoomID  string `json:"room_id" gorm:"column:room_id"`
	Balcony bool   `json:"balcony" gorm:"column:balcony"`
	BathTub bool   `json:"bath_tub" gorm:"column:bath_tub"`
	View    string `json:"view" gorm:"column:view;default:city_view"` // city_view, sea_view, mountain_view
}

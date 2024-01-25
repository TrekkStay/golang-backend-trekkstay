package entity

type RoomFacilityEntity struct {
	RoomID      string `json:"room_id" gorm:"not null;"`
	RoomSize    int    `json:"room_size" gorm:"column:room_size"`
	NumberOfBed int    `json:"number_of_bed" gorm:"column:number_of_bed;default:1"`
	View        string `json:"view" gorm:"column:view;default:none"` // none, city_view, sea_view, mountain_view
	// Default: false
	Balcony bool `json:"balcony" gorm:"column:balcony;default:false"`
	BathTub bool `json:"bath_tub" gorm:"column:bath_tub;default:false"`
	Kitchen bool `json:"kitchen" gorm:"column:kitchen;default:false"`
	// Default: true
	Television     bool `json:"television" gorm:"column:television;default:true"`
	Shower         bool `json:"shower" gorm:"column:shower;default:true"`
	NonSmoking     bool `json:"non_smoking" gorm:"column:non_smoking;default:true"`
	HairDryer      bool `json:"hair_dryer" gorm:"column:hair_dryer;default:true"`
	AirConditioner bool `json:"air_conditioner" gorm:"column:air_conditioner;default:true"`
	Slippers       bool `json:"slippers" gorm:"column:slipper;default:true"`
	// Jsonb
	Sleeps SleepsEntity `json:"sleeps" gorm:"type:jsonb;default:null"`
}

func (RoomFacilityEntity) TableName() string {
	return "room_facilities"
}

package req

type SleepJSON struct {
	Adults   int `json:"adults" gorm:"default:2"`
	Children int `json:"children" gorm:"default:1"`
}

type MediaJSON struct {
	URLS []string `json:"urls"`
}

type HotelRoomFacilitiesJSON struct {
	RoomSize    int    `json:"room_size"`
	NumberOfBed int    `json:"number_of_bed"`
	View        string `json:"view"`
	// Default: false
	Balcony bool `json:"balcony" gorm:"default:false"`
	BathTub bool `json:"bath_tub" gorm:"default:false"`
	Kitchen bool `json:"kitchen" gorm:"default:false"`
	// Default: true
	Television     bool `json:"television" gorm:"default:true"`
	Shower         bool `json:"shower" gorm:"default:true"`
	NonSmoking     bool `json:"non_smoking" gorm:"default:true"`
	HairDryer      bool `json:"hair_dryer" gorm:"default:true"`
	AirConditioner bool `json:"air_conditioner" gorm:"default:true"`
	Slippers       bool `json:"slippers" gorm:"default:true"`
	// Jsonb
	Sleeps SleepJSON `json:"sleeps"`
}

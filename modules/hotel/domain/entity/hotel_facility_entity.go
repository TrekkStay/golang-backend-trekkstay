package entity

type HotelFacilityEntity struct {
	HotelID         string `json:"hotel_id" gorm:"not null;"`
	FitnessCenter   bool   `json:"fitness_center" gorm:"default:false"`
	ConferenceRoom  bool   `json:"conference_room" gorm:"default:false"`
	ParkingArea     bool   `json:"parking_area" gorm:"default:false"`
	SwimmingPool    bool   `json:"swimming_pool" gorm:"default:false"`
	FreeWifi        bool   `json:"free_wifi" gorm:"default:false"`
	AirportTransfer bool   `json:"airport_transfer" gorm:"default:false"`
	MotorBikeRental bool   `json:"motor_bike_rental" gorm:"default:false"`
	SpaService      bool   `json:"spa_service" gorm:"default:false"`
	FoodService     bool   `json:"food_service" gorm:"default:false"`
	LaundryService  bool   `json:"laundry_service" gorm:"default:false"`
}

func (HotelFacilityEntity) TableName() string {
	return "hotel_facilities"
}

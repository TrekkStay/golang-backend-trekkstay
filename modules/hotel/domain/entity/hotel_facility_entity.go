package entity

type HotelFacilityEntity struct {
	HotelID         string `json:"hotel_id" gorm:"not null;"`
	CheckInHour     string `json:"check_in_hour" gorm:"default:14:00"`
	CheckOutHour    string `json:"check_out_hour" gorm:"default:12:00"`
	HotTub          bool   `json:"hot_tub" gorm:"default:false"`
	CarPark         bool   `json:"car_park" gorm:"default:false"`
	SwimmingPool    bool   `json:"swimming_pool" gorm:"default:false"`
	FreeWifi        bool   `json:"free_wifi" gorm:"default:false"`
	FoodService     bool   `json:"food_service" gorm:"default:false"`
	LaundryService  bool   `json:"laundry_service" gorm:"default:false"`
	AirportTransfer bool   `json:"airport_transfer" gorm:"default:false"`
	MotorBikeRental bool   `json:"motor_bike_rental" gorm:"default:false"`
}

func (HotelFacilityEntity) TableName() string {
	return "hotel_facilities"
}

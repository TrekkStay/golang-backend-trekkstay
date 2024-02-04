package req

type MediaJSON struct {
	URLS []string `json:"urls"`
}

type CoordinatesJSON struct {
	Lng float64 `json:"lng"`
	Lat float64 `json:"lat"`
}

type HotelFacilitiesJSON struct {
	FitnessCenter   bool `json:"fitness_center"`
	ConferenceRoom  bool `json:"conference_room"`
	ParkingArea     bool `json:"parking_area"`
	SwimmingPool    bool `json:"swimming_pool"`
	FreeWifi        bool `json:"free_wifi"`
	AirportTransfer bool `json:"airport_transfer"`
	MotorBikeRental bool `json:"motor_bike_rental"`
	SpaService      bool `json:"spa_service"`
	FoodService     bool `json:"food_service"`
	LaundryService  bool `json:"laundry_service"`
}

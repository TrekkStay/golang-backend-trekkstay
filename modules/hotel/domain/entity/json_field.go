package entity

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type MediaJSON struct {
	URLS []string `json:"urls"`
}

type CoordinatesJSON struct {
	Lng float64 `json:"lng"`
	Lat float64 `json:"lat"`
}

type HotelFacilitiesJSON struct {
	FitnessCenter   bool `json:"fitness_center" gorm:"default:true"`
	ConferenceRoom  bool `json:"conference_room" gorm:"default:true"`
	ParkingArea     bool `json:"parking_area" gorm:"default:true"`
	SwimmingPool    bool `json:"swimming_pool" gorm:"default:true"`
	FreeWifi        bool `json:"free_wifi" gorm:"default:true"`
	AirportTransfer bool `json:"airport_transfer" gorm:"default:true"`
	MotorBikeRental bool `json:"motor_bike_rental" gorm:"default:false"`
	SpaService      bool `json:"spa_service" gorm:"default:false"`
	FoodService     bool `json:"food_service" gorm:"default:true"`
	LaundryService  bool `json:"laundry_service" gorm:"default:true"`
}

// ---------------------------- Media ----------------------------

func (media *MediaJSON) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf(fmt.Sprint("Failed to unmarshal JSONB value", value))
	}

	var m MediaJSON
	if err := json.Unmarshal(bytes, &m); err != nil {
		return err
	}

	*media = m
	return nil
}

func (media *MediaJSON) Value() (driver.Value, error) {
	if media == nil {
		return nil, nil
	}

	return json.Marshal(media)
}

// ---------------------------- Coordinates ----------------------------

func (coords *CoordinatesJSON) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf(fmt.Sprint("Failed to unmarshal JSONB value", value))
	}

	var c CoordinatesJSON
	if err := json.Unmarshal(bytes, &c); err != nil {
		return err
	}

	*coords = c
	return nil
}

func (coords *CoordinatesJSON) Value() (driver.Value, error) {
	if coords == nil {
		return nil, nil
	}

	return json.Marshal(coords)
}

// ---------------------------- Hotel Facilities ----------------------------

func (hotelFacilities *HotelFacilitiesJSON) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf(fmt.Sprint("Failed to unmarshal JSONB value", value))
	}

	var h HotelFacilitiesJSON
	if err := json.Unmarshal(bytes, &h); err != nil {
		return err
	}

	*hotelFacilities = h
	return nil
}

func (hotelFacilities *HotelFacilitiesJSON) Value() (driver.Value, error) {
	if hotelFacilities == nil {
		return nil, nil
	}

	return json.Marshal(hotelFacilities)
}

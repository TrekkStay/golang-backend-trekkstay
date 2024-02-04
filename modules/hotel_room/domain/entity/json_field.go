package entity

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

// ---------------------------- JSON Object ----------------------------

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

// ---------------------------- Sleep ----------------------------

func (sleep *SleepJSON) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf(fmt.Sprint("Failed to unmarshal JSONB value", value))
	}

	var s SleepJSON
	if err := json.Unmarshal(bytes, &s); err != nil {
		return err
	}

	*sleep = s
	return nil
}

func (sleep *SleepJSON) Value() (driver.Value, error) {
	if sleep == nil {
		return nil, nil
	}

	return json.Marshal(sleep)
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

// ---------------------------- Hotel Room Facilities ----------------------------

func (hotelRoomFacilities *HotelRoomFacilitiesJSON) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf(fmt.Sprint("Failed to unmarshal JSONB value", value))
	}

	var h HotelRoomFacilitiesJSON
	if err := json.Unmarshal(bytes, &h); err != nil {
		return err
	}

	*hotelRoomFacilities = h
	return nil
}

func (hotelRoomFacilities *HotelRoomFacilitiesJSON) Value() (driver.Value, error) {
	if hotelRoomFacilities == nil {
		return nil, nil
	}

	return json.Marshal(hotelRoomFacilities)
}

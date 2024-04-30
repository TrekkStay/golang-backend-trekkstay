package entity

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type MediaJSON struct {
	URLS []string `json:"urls"`
}

type RoomJSON struct {
	HotelID       string    `json:"hotel_id"`
	HotelName     string    `json:"hotel_name"`
	Location      string    `json:"location"`
	Type          string    `json:"type"`
	OriginalPrice int       `json:"original_price"`
	BookingPrice  int       `json:"booking_price"`
	Images        MediaJSON `json:"images"`
}

type GuestInfoJSON struct {
	FullName string `json:"full_name"`
	Contact  string `json:"contact"`
	Adults   int    `json:"adults"`
	Children int    `json:"children"`
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

// ---------------------------- Room ----------------------------

func (room *RoomJSON) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf(fmt.Sprint("Failed to unmarshal JSONB value", value))
	}

	var r RoomJSON
	if err := json.Unmarshal(bytes, &r); err != nil {
		return err
	}

	*room = r
	return nil
}

func (room *RoomJSON) Value() (driver.Value, error) {
	if room == nil {
		return nil, nil
	}

	return json.Marshal(room)
}

// ---------------------------- GuestInfo ----------------------------

func (guestInfo *GuestInfoJSON) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf(fmt.Sprint("Failed to unmarshal JSONB value", value))
	}

	var g GuestInfoJSON
	if err := json.Unmarshal(bytes, &g); err != nil {
		return err
	}

	*guestInfo = g
	return nil
}

func (guestInfo *GuestInfoJSON) Value() (driver.Value, error) {
	if guestInfo == nil {
		return nil, nil
	}

	return json.Marshal(guestInfo)
}

package entity

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type SleepsEntity struct {
	Adults   int `json:"adults"`
	Children int `json:"children"`
}

type MediaObject struct {
	URL []string `json:"url"`
}

type CoordinatesEntity struct {
	Lng float64 `json:"lng"`
	Lat float64 `json:"lat"`
}

func (sleep *SleepsEntity) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf(fmt.Sprint("Failed to unmarshal JSONB value", value))
	}

	var sleepEntity SleepsEntity
	if err := json.Unmarshal(bytes, &sleepEntity); err != nil {
		return err
	}

	*sleep = sleepEntity
	return nil
}

func (sleep *SleepsEntity) Value() (driver.Value, error) {
	if sleep == nil {
		return nil, nil
	}

	return json.Marshal(sleep)
}

func (img *MediaObject) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf(fmt.Sprint("Failed to unmarshal JSONB value", value))
	}

	var image MediaObject
	if err := json.Unmarshal(bytes, &image); err != nil {
		return err
	}

	*img = image
	return nil
}

func (img *MediaObject) Value() (driver.Value, error) {
	if img == nil {
		return nil, nil
	}

	return json.Marshal(img)
}

func (coords *CoordinatesEntity) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf(fmt.Sprint("Failed to unmarshal JSONB value", value))
	}

	var coordinatesEntity CoordinatesEntity
	if err := json.Unmarshal(bytes, &coordinatesEntity); err != nil {
		return err
	}

	*coords = coordinatesEntity
	return nil
}

func (coords *CoordinatesEntity) Value() (driver.Value, error) {
	if coords == nil {
		return nil, nil
	}

	return json.Marshal(coords)
}

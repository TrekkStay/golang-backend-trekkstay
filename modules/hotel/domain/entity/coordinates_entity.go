package entity

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type CoordinatesEntity struct {
	Lng float64 `json:"lng"`
	Lat float64 `json:"lat"`
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

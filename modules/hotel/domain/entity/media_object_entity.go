package entity

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type MediaObject struct {
	URL []string `json:"url"`
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

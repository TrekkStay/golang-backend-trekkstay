package core

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Entity struct with common fields for almost all entities
type Entity struct {
	ID        string         `json:"id"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoCreateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`
}

// BeforeCreate is a hook that generates a UUID before creating an entity.
// It ensures that the ID field is populated with a unique identifier.
func (record *Entity) BeforeCreate(*gorm.DB) (err error) {
	// Check id field is empty
	if record.ID == "" {
		// Assign the generated UUID to the ID field
		record.ID = uuid.New().String()
	}

	return nil
}

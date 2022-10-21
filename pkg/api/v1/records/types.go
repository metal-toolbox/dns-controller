package record

import (
	"time"

	"github.com/google/uuid"
)

// Record is the API model for a record
type Record struct {
	Name      string `json:"record"`
	Type      string `json:"record_type"`
	path      string
	UUID      uuid.UUID `json:"uuid"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

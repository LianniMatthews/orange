package data

import (
	"time"
)

// School represents one row of data in database
type School struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Level     string    `json:"level"`
	Contact   string    `json:"contact"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	Website   string    `json:"website,omitempty"`
	Address   string    `json:"address"`
	Mode      []string  `json:"mode"`
	CreatedAt time.Time `json:"-"`
	Version   int32     `json:"version"`
}

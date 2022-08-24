package element

import "time"

type Element struct {
	ID        uint      `json:"id,omitempty"`
	Key       string    `json:"key,omitempty"`
	Value     string    `json:"value,omitempty"`
	UserID    uint      `json:"user_id,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

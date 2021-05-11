package model

import "time"

// Purchase represents purchase model.
type Purchase struct {
	ID       int       `json:"id,omitempty"`
	UserId   int       `json:"user_id"`
	Date     time.Time `json:"date"`
	FileName string    `json:"file_name"`
}

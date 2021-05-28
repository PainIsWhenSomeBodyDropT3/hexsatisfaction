package model

import "time"

// Purchase represents purchase model.
type Purchase struct {
	ID     int       `json:"id,omitempty"`
	UserID int       `json:"userID"`
	Date   time.Time `json:"date"`
	FileID int       `json:"fileID"`
}

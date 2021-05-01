package model

import "time"

type Purchase struct {
	ID       int       `json:"id,omitempty"`
	UserId   int       `json:"user_id"`
	Date     time.Time `json:"date"`
	FileName string    `json:"file_name"`
}

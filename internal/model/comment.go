package model

import "time"

// Comment represents comment model.
type Comment struct {
	ID         int       `json:"id,omitempty"`
	UserID     int       `json:"userID"`
	PurchaseID int       `json:"purchaseID"`
	Date       time.Time `json:"date"`
	Text       string    `json:"text"`
}

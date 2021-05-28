package model

// Author represents author model.
type Author struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name"`
	Age         int    `json:"age"`
	Description string `json:"description"`
	UserID      int    `json:"userID"`
}

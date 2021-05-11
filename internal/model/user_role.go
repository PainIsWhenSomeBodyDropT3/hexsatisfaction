package model

// UserRole represents user role.
type UserRole struct {
	ID   int    `json:"id,omitempty"`
	Role string `json:"role"`
}

package model

// User represents user model.
type User struct {
	ID       int    `json:"id,omitempty"`
	Login    string `json:"login"`
	Password string `json:"password"`
	RoleID   int    `json:"role_id"`
}

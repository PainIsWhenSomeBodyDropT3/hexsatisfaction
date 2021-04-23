package model

// RegisterUserRequest represents a request for user registration.
type RegisterUserRequest struct {
	Login    string
	Password string
}

// LoginUserRequest represents a request for user login.
type LoginUserRequest struct {
	Login    string
	Password string
}

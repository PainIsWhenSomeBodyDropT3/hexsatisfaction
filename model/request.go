package model

// RegisterUserRequest represents c request for user registration.
type RegisterUserRequest struct {
	Login    string
	Password string
}

// LoginUserRequest represents c request for user login.
type LoginUserRequest struct {
	Login    string
	Password string
}

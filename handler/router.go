package handler

import (
	"github.com/gorilla/mux"
)

const userPath = "/user"

// API represents a structure with APIs.
type API struct {
	*mux.Router
}

// NewRouter creates and serves endpoints of API.
func NewRouter(userService UserService) *API {
	api := API{
		mux.NewRouter(),
	}
	api.PathPrefix(userPath).Handler(newUser(userService))

	return &api
}

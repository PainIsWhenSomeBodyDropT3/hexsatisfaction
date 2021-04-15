package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

type userRouter struct {
	*mux.Router
	service UserService
}

func newUser(userService UserService) userRouter {
	router := mux.NewRouter().PathPrefix(userPath).Subrouter()

	handler := userRouter{
		router,
		userService,
	}

	router.Path("/login").
		Methods(http.MethodPost).
		HandlerFunc(loginUser)

	router.Path("/registration").
		Methods(http.MethodPost).
		HandlerFunc(registerUser)

	return handler

}

func loginUser(w http.ResponseWriter, r *http.Request) {

}

func registerUser(w http.ResponseWriter, r *http.Request) {

}

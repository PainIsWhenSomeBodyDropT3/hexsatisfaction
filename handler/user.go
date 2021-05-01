package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/JesusG2000/hexsatisfaction/jwt"
	"github.com/JesusG2000/hexsatisfaction/middleware"
	"github.com/JesusG2000/hexsatisfaction/model"
	"github.com/gorilla/mux"
)

type userRouter struct {
	*mux.Router
	service      UserService
	tokenManager jwt.TokenManager
}

func newUser(userService UserService, tokenManager *jwt.Manager) userRouter {
	router := mux.NewRouter().PathPrefix(userPath).Subrouter()

	handler := userRouter{
		router,
		userService,
		tokenManager,
	}

	router.Path("/login").
		Methods(http.MethodPost).
		HandlerFunc(handler.loginUser)

	router.Path("/registration").
		Methods(http.MethodPost).
		HandlerFunc(handler.registerUser)

	return handler

}

type loginRequest struct {
	model.LoginUserRequest
}

// Build builds request for user login.
func (req *loginRequest) Build(r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(&req.LoginUserRequest)
	defer r.Body.Close()
	if err != nil {
		return err
	}

	return nil
}

// Validate validates request for user login.
func (req *loginRequest) Validate() error {
	switch {
	case req.Login == "":
		return fmt.Errorf("login is required")
	case req.Password == "":
		return fmt.Errorf("password is required")
	default:
		return nil
	}
}

func (u *userRouter) loginUser(w http.ResponseWriter, r *http.Request) {
	var req loginRequest
	err := middleware.ParseRequest(r, &req)
	if err != nil {
		middleware.JSONError(w, err, http.StatusBadRequest)
		return
	}
	token, err := u.service.FindByCredentials(req.LoginUserRequest)
	if err != nil {
		middleware.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	if len(token) == 0 {
		middleware.Empty(w, http.StatusNotFound)
		return
	}

	middleware.JSONReturn(w, http.StatusOK, token)

}

type registerRequest struct {
	model.RegisterUserRequest
}

// Build builds request for user registration.
func (req *registerRequest) Build(r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(&req.RegisterUserRequest)
	defer r.Body.Close()
	if err != nil {
		return err
	}

	return nil
}

// Validate validates request for user registration.
func (req *registerRequest) Validate() error {
	switch {
	case req.Login == "":
		return fmt.Errorf("login is required")
	case req.Password == "":
		return fmt.Errorf("password is required")
	default:
		return nil
	}
}

func (u *userRouter) registerUser(w http.ResponseWriter, r *http.Request) {
	var req registerRequest
	err := middleware.ParseRequest(r, &req)
	if err != nil {
		middleware.JSONError(w, err, http.StatusBadRequest)
		return
	}

	exist, err := u.service.IsExist(req.Login)
	if err != nil {
		middleware.JSONError(w, err, http.StatusInternalServerError)
		return
	}
	if exist {
		middleware.JSONReturn(w, http.StatusFound, "this user already exist")
		return
	}

	err = u.service.Create(req.RegisterUserRequest)
	if err != nil {
		middleware.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	middleware.JSONReturn(w, http.StatusOK, "user created")
}

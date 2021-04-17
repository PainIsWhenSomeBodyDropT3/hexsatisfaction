package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/JesusG2000/hexsatisfaction/middleware"
	"github.com/JesusG2000/hexsatisfaction/model"
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
		HandlerFunc(handler.loginUser)

	router.Path("/registration").
		Methods(http.MethodPost).
		HandlerFunc(handler.registerUser)

	return handler

}

type loginRequest struct {
	model.LoginUserRequest
}

func (req *loginRequest) Build(r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(&req.LoginUserRequest)
	if err != nil {
		return err
	}

	return nil
}

func (req *loginRequest) Validate() error {
	if req.Login == "" {
		return fmt.Errorf("login is required")
	}
	if req.Password == "" {
		return fmt.Errorf("password is required")
	}
	return nil
}

func (u *userRouter) loginUser(w http.ResponseWriter, r *http.Request) {
	var req loginRequest
	err := middleware.ParseRequest(r, &req)
	if err != nil {
		middleware.JSONError(w, err, http.StatusBadRequest)
		return
	}

	user, err := u.service.FindByCredentials(req.LoginUserRequest)
	if err != nil {
		middleware.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	if user.ID == 0 {
		middleware.Empty(w, http.StatusNotFound)
		return
	}

	middleware.JSONReturn(w, http.StatusOK, user)

}

type registerRequest struct {
	model.RegisterUserRequest
}

func (req *registerRequest) Build(r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(&req.RegisterUserRequest)
	if err != nil {
		return err
	}

	return nil
}

func (req *registerRequest) Validate() error {
	if req.Login == "" {
		return fmt.Errorf("login is required")
	}
	if req.Password == "" {
		return fmt.Errorf("password is required")
	}
	return nil
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

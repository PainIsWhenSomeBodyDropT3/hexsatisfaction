package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/JesusG2000/hexsatisfaction/internal/model"
	"github.com/JesusG2000/hexsatisfaction/internal/service"
	"github.com/JesusG2000/hexsatisfaction/pkg/auth"
	"github.com/JesusG2000/hexsatisfaction/pkg/middleware"
	"github.com/gorilla/mux"
)

type userRouter struct {
	*mux.Router
	services     *service.Services
	tokenManager auth.TokenManager
}

func newUser(services *service.Services, tokenManager auth.TokenManager) userRouter {
	router := mux.NewRouter().PathPrefix(userPath).Subrouter()
	handler := userRouter{
		router,
		services,
		tokenManager,
	}

	router.Path("/login").
		Methods(http.MethodPost).
		HandlerFunc(handler.loginUser)

	router.Path("/registration").
		Methods(http.MethodPost).
		HandlerFunc(handler.registerUser)

	secure := router.PathPrefix("/api").Subrouter()
	secure.Use(handler.tokenManager.UserIdentity)

	secure.Path("/getAll").
		Methods(http.MethodGet).
		HandlerFunc(handler.getAllUser)

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

	token, err := u.services.User.FindByCredentials(req.LoginUserRequest)
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

	exist, err := u.services.User.IsExist(req.Login)
	if err != nil {
		middleware.JSONError(w, err, http.StatusInternalServerError)
		return
	}
	if exist {
		middleware.JSONReturn(w, http.StatusFound, "this user already exist")
		return
	}

	id, err := u.services.User.Create(req.RegisterUserRequest)
	if err != nil {
		middleware.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	middleware.JSONReturn(w, http.StatusOK, strconv.Itoa(id))
}

func (u *userRouter) getAllUser(w http.ResponseWriter, r *http.Request) {
	users, err := u.services.UserRole.FindAllUser()
	if err != nil {
		middleware.JSONError(w, err, http.StatusInternalServerError)
		return
	}
	middleware.JSONReturn(w, http.StatusOK, users)
}

package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/JesusG2000/hexsatisfaction/internal/model"
	"github.com/JesusG2000/hexsatisfaction/internal/service"
	"github.com/JesusG2000/hexsatisfaction/pkg/auth"
	"github.com/JesusG2000/hexsatisfaction/pkg/middleware"
	"github.com/gorilla/mux"
)

type authorRouter struct {
	*mux.Router
	services     *service.Services
	tokenManager auth.TokenManager
}

func newAuthor(services *service.Services, tokenManager auth.TokenManager) authorRouter {
	router := mux.NewRouter().PathPrefix(authorPath).Subrouter()
	handler := authorRouter{
		router,
		services,
		tokenManager,
	}

	router.Path("/{name}").
		Methods(http.MethodGet).
		HandlerFunc(handler.findByNameAuthor)

	router.Path("/").
		Methods(http.MethodGet).
		HandlerFunc(handler.findAllAuthor)

	secure := router.PathPrefix("/api").Subrouter()
	secure.Use(handler.tokenManager.UserIdentity)

	secure.Path("/").
		Methods(http.MethodPost).
		HandlerFunc(handler.createAuthor)

	secure.Path("/{id}").
		Methods(http.MethodPut).
		HandlerFunc(handler.updateAuthor)

	secure.Path("/{id}").
		Methods(http.MethodDelete).
		HandlerFunc(handler.deleteAuthor)

	secure.Path("/{id}").
		Methods(http.MethodGet).
		HandlerFunc(handler.findByIDAuthor)

	secure.Path("/user/{id}").
		Methods(http.MethodGet).
		HandlerFunc(handler.findByUserIDAuthor)

	return handler
}

type createAuthorRequest struct {
	model.CreateAuthorRequest
}

// Build builds request for create author.
func (req *createAuthorRequest) Build(r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(&req.CreateAuthorRequest)
	if err != nil {
		return err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("%v", err)
		}
	}(r.Body)

	return nil
}

// Validate validates request for create author.
func (req *createAuthorRequest) Validate() error {
	switch {
	case req.UserID < 1:
		return fmt.Errorf("not correct user id")
	case req.Age < 1:
		return fmt.Errorf("not correct age")
	case req.Name == "":
		return fmt.Errorf("name is required")
	case req.Description == "":
		return fmt.Errorf("description is required")
	default:
		return nil
	}
}

func (a *authorRouter) createAuthor(w http.ResponseWriter, r *http.Request) {
	var req createAuthorRequest
	err := middleware.ParseRequest(r, &req)
	if err != nil {
		middleware.JSONError(w, err, http.StatusBadRequest)
		return
	}

	id, err := a.services.Author.Create(req.CreateAuthorRequest)
	if err != nil {
		middleware.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	middleware.JSONReturn(w, http.StatusOK, strconv.Itoa(id))
}

type updateAuthorRequest struct {
	model.UpdateAuthorRequest
}

// Build builds request for update author.
func (req *updateAuthorRequest) Build(r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(&req.UpdateAuthorRequest)
	if err != nil {
		return err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("%v", err)
		}
	}(r.Body)

	vID, ok := mux.Vars(r)["id"]
	if !ok {
		return fmt.Errorf("no id")
	}

	id, err := strconv.Atoi(vID)
	if err != nil {
		return err
	}

	req.ID = id

	return nil
}

// Validate validates request for update author.
func (req *updateAuthorRequest) Validate() error {
	switch {
	case req.ID < 1:
		return fmt.Errorf("not correct id")
	case req.UserID < 1:
		return fmt.Errorf("not correct user id")
	case req.Age < 1:
		return fmt.Errorf("not correct age")
	case req.Name == "":
		return fmt.Errorf("name is required")
	case req.Description == "":
		return fmt.Errorf("description is required")
	default:
		return nil
	}
}

func (a *authorRouter) updateAuthor(w http.ResponseWriter, r *http.Request) {
	var req updateAuthorRequest
	err := middleware.ParseRequest(r, &req)
	if err != nil {
		middleware.JSONError(w, err, http.StatusBadRequest)
		return
	}

	id, err := a.services.Author.Update(req.UpdateAuthorRequest)
	if err != nil {
		middleware.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	if id < 1 {
		middleware.Empty(w, http.StatusNotFound)
		return
	}

	middleware.JSONReturn(w, http.StatusOK, strconv.Itoa(id))
}

type deleteAuthorRequest struct {
	model.DeleteAuthorRequest
}

// Build builds request for delete author.
func (req *deleteAuthorRequest) Build(r *http.Request) error {
	vID, ok := mux.Vars(r)["id"]
	if !ok {
		return fmt.Errorf("no id")
	}

	id, err := strconv.Atoi(vID)
	if err != nil {
		return err
	}

	req.ID = id

	return nil
}

// Validate validates request for delete author.
func (req *deleteAuthorRequest) Validate() error {
	switch {
	case req.ID < 1:
		return fmt.Errorf("not correct id")
	default:
		return nil
	}
}

func (a *authorRouter) deleteAuthor(w http.ResponseWriter, r *http.Request) {
	var req deleteAuthorRequest
	err := middleware.ParseRequest(r, &req)
	if err != nil {
		middleware.JSONError(w, err, http.StatusBadRequest)
		return
	}

	id, err := a.services.Author.Delete(req.DeleteAuthorRequest)
	if err != nil {
		middleware.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	if id < 1 {
		middleware.Empty(w, http.StatusNotFound)
		return
	}

	middleware.JSONReturn(w, http.StatusOK, strconv.Itoa(id))
}

type idAuthorRequest struct {
	model.IDAuthorRequest
}

// Build builds request to find author by id.
func (req *idAuthorRequest) Build(r *http.Request) error {
	vID, ok := mux.Vars(r)["id"]
	if !ok {
		return fmt.Errorf("no id")
	}

	id, err := strconv.Atoi(vID)
	if err != nil {
		return err
	}

	req.ID = id

	return nil
}

// Validate validates request to find author by id.
func (req *idAuthorRequest) Validate() error {
	switch {
	case req.ID < 1:
		return fmt.Errorf("not correct id")
	default:
		return nil
	}
}

func (a *authorRouter) findByIDAuthor(w http.ResponseWriter, r *http.Request) {
	var req idAuthorRequest
	err := middleware.ParseRequest(r, &req)
	if err != nil {
		middleware.JSONError(w, err, http.StatusBadRequest)
		return
	}

	author, err := a.services.Author.FindByID(req.IDAuthorRequest)
	if err != nil {
		middleware.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	if author.ID < 1 {
		middleware.Empty(w, http.StatusNotFound)
		return
	}

	middleware.JSONReturn(w, http.StatusOK, author)
}

type userIDAuthorRequest struct {
	model.UserIDAuthorRequest
}

// Build builds request to find author by user id.
func (req *userIDAuthorRequest) Build(r *http.Request) error {
	vID, ok := mux.Vars(r)["id"]
	if !ok {
		return fmt.Errorf("no id")
	}

	id, err := strconv.Atoi(vID)
	if err != nil {
		return err
	}

	req.ID = id

	return nil
}

// Validate validates request to find author by user id.
func (req *userIDAuthorRequest) Validate() error {
	switch {
	case req.ID < 1:
		return fmt.Errorf("not correct id")
	default:
		return nil
	}
}

func (a *authorRouter) findByUserIDAuthor(w http.ResponseWriter, r *http.Request) {
	var req userIDAuthorRequest
	err := middleware.ParseRequest(r, &req)
	if err != nil {
		middleware.JSONError(w, err, http.StatusBadRequest)
		return
	}

	author, err := a.services.Author.FindByUserID(req.UserIDAuthorRequest)
	if err != nil {
		middleware.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	if author.ID < 1 {
		middleware.Empty(w, http.StatusNotFound)
		return
	}

	middleware.JSONReturn(w, http.StatusOK, author)
}

type nameAuthorRequest struct {
	model.NameAuthorRequest
}

// Build builds request to find authors by name.
func (req *nameAuthorRequest) Build(r *http.Request) error {
	name, ok := mux.Vars(r)["name"]
	if !ok {
		return fmt.Errorf("no id")
	}

	req.Name = name

	return nil
}

// Validate validates request to find authors by name.
func (req *nameAuthorRequest) Validate() error {
	switch {
	case req.Name == "":
		return fmt.Errorf("name is required")
	default:
		return nil
	}
}

func (a *authorRouter) findByNameAuthor(w http.ResponseWriter, r *http.Request) {
	var req nameAuthorRequest
	err := middleware.ParseRequest(r, &req)
	if err != nil {
		middleware.JSONError(w, err, http.StatusBadRequest)
		return
	}

	authors, err := a.services.Author.FindByName(req.NameAuthorRequest)
	if err != nil {
		middleware.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	if len(authors) == 0 {
		middleware.Empty(w, http.StatusNotFound)
		return
	}

	middleware.JSONReturn(w, http.StatusOK, authors)
}

func (a *authorRouter) findAllAuthor(w http.ResponseWriter, r *http.Request) {
	authors, err := a.services.Author.FindAll()
	if err != nil {
		middleware.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	if len(authors) == 0 {
		middleware.Empty(w, http.StatusNotFound)
		return
	}

	middleware.JSONReturn(w, http.StatusOK, authors)
}

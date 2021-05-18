package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/JesusG2000/hexsatisfaction/internal/model"
	"github.com/JesusG2000/hexsatisfaction/internal/service"
	"github.com/JesusG2000/hexsatisfaction/pkg/auth"
	"github.com/JesusG2000/hexsatisfaction/pkg/middleware"
	"github.com/gorilla/mux"
)

type commentRouter struct {
	*mux.Router
	services     *service.Services
	tokenManager auth.TokenManager
}

func newComment(services *service.Services, tokenManager auth.TokenManager) commentRouter {
	router := mux.NewRouter().PathPrefix(commentPath).Subrouter()
	handler := commentRouter{
		router,
		services,
		tokenManager,
	}

	router.Path("/user/{id}").
		Methods(http.MethodGet).
		HandlerFunc(handler.findByUserIDComment)

	router.Path("/purchase/{id}").
		Methods(http.MethodGet).
		HandlerFunc(handler.findByPurchaseIDComment)

	router.Path("/user/{userID}/purchase/{purchaseID}").
		Methods(http.MethodGet).
		HandlerFunc(handler.findByUserIDAndPurchaseIDComment)

	router.Path("/text").
		Methods(http.MethodPost).
		HandlerFunc(handler.findByTextComment)

	router.Path("/period").
		Methods(http.MethodPost).
		HandlerFunc(handler.findByPeriodComment)

	secure := router.PathPrefix("/api").Subrouter()
	secure.Use(handler.tokenManager.UserIdentity)

	secure.Path("/").
		Methods(http.MethodGet).
		HandlerFunc(handler.findAllComment)

	secure.Path("/").
		Methods(http.MethodPost).
		HandlerFunc(handler.createComment)

	secure.Path("/{id}").
		Methods(http.MethodPut).
		HandlerFunc(handler.updateComment)

	secure.Path("/{id}").
		Methods(http.MethodDelete).
		HandlerFunc(handler.deleteComment)

	secure.Path("/{id}").
		Methods(http.MethodGet).
		HandlerFunc(handler.findByIDComment)

	return handler
}

type createCommentRequest struct {
	model.CreateCommentRequest
}

// Build builds request for create comment.
func (req *createCommentRequest) Build(r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(&req.CreateCommentRequest)
	if err != nil {
		return err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(r.Body)

	return nil
}

// Validate validates request for create comment.
func (req *createCommentRequest) Validate() error {
	switch {
	case req.UserID < 1:
		return fmt.Errorf("not correct user id")
	case req.PurchaseID < 1:
		return fmt.Errorf("not correct purchase id")
	case req.Date == time.Time{}:
		return fmt.Errorf("date is required")
	case req.Text == "":
		return fmt.Errorf("text is required")
	default:
		return nil
	}
}

func (c *commentRouter) createComment(w http.ResponseWriter, r *http.Request) {
	var req createCommentRequest
	err := middleware.ParseRequest(r, &req)
	if err != nil {
		middleware.JSONError(w, err, http.StatusBadRequest)
		return
	}

	id, err := c.services.Comment.Create(req.CreateCommentRequest)
	if err != nil {
		middleware.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	middleware.JSONReturn(w, http.StatusOK, strconv.Itoa(id))
}

type updateCommentRequest struct {
	model.UpdateCommentRequest
}

// Build builds request for update comment.
func (req *updateCommentRequest) Build(r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(&req.UpdateCommentRequest)
	if err != nil {
		return err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
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

// Validate validates request for update comment.
func (req *updateCommentRequest) Validate() error {
	switch {
	case req.ID < 1:
		return fmt.Errorf("not correct id")
	case req.UserID < 1:
		return fmt.Errorf("not correct user id")
	case req.PurchaseID < 1:
		return fmt.Errorf("not correct purchase id")
	case req.Date == time.Time{}:
		return fmt.Errorf("date is required")
	case req.Text == "":
		return fmt.Errorf("text is required")
	default:
		return nil
	}
}

func (c *commentRouter) updateComment(w http.ResponseWriter, r *http.Request) {
	var req updateCommentRequest
	err := middleware.ParseRequest(r, &req)
	if err != nil {
		middleware.JSONError(w, err, http.StatusBadRequest)
		return
	}

	id, err := c.services.Comment.Update(req.UpdateCommentRequest)
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

type deleteCommentRequest struct {
	model.DeleteCommentRequest
}

// Build builds request for delete comment.
func (req *deleteCommentRequest) Build(r *http.Request) error {
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

// Validate validates request for delete comment.
func (req *deleteCommentRequest) Validate() error {
	switch {
	case req.ID < 1:
		return fmt.Errorf("not correct id")
	default:
		return nil
	}
}

func (c *commentRouter) deleteComment(w http.ResponseWriter, r *http.Request) {
	var req deleteCommentRequest
	err := middleware.ParseRequest(r, &req)
	if err != nil {
		middleware.JSONError(w, err, http.StatusBadRequest)
		return
	}

	id, err := c.services.Comment.Delete(req.DeleteCommentRequest)
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

type idCommentRequest struct {
	model.IDCommentRequest
}

// Build builds request to find comment by id.
func (req *idCommentRequest) Build(r *http.Request) error {
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

// Validate validates request to find comment by id.
func (req *idCommentRequest) Validate() error {
	switch {
	case req.ID < 1:
		return fmt.Errorf("not correct id")
	default:
		return nil
	}
}

func (c *commentRouter) findByIDComment(w http.ResponseWriter, r *http.Request) {
	var req idCommentRequest
	err := middleware.ParseRequest(r, &req)
	if err != nil {
		middleware.JSONError(w, err, http.StatusBadRequest)
		return
	}

	comment, err := c.services.Comment.FindByID(req.IDCommentRequest)
	if err != nil {
		middleware.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	if comment.ID < 1 {
		middleware.Empty(w, http.StatusNotFound)
		return
	}

	middleware.JSONReturn(w, http.StatusOK, comment)
}

type userIDCommentRequest struct {
	model.UserIDCommentRequest
}

// Build builds request to find comment by user id.
func (req *userIDCommentRequest) Build(r *http.Request) error {
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

// Validate validates request to find comment by user id.
func (req *userIDCommentRequest) Validate() error {
	switch {
	case req.ID < 1:
		return fmt.Errorf("not correct id")
	default:
		return nil
	}
}

func (c *commentRouter) findByUserIDComment(w http.ResponseWriter, r *http.Request) {
	var req userIDCommentRequest
	err := middleware.ParseRequest(r, &req)
	if err != nil {
		middleware.JSONError(w, err, http.StatusBadRequest)
		return
	}

	comments, err := c.services.Comment.FindAllByUserID(req.UserIDCommentRequest)
	if err != nil {
		middleware.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	if len(comments) == 0 {
		middleware.Empty(w, http.StatusNotFound)
		return
	}

	middleware.JSONReturn(w, http.StatusOK, comments)
}

type purchaseIDCommentRequest struct {
	model.PurchaseIDCommentRequest
}

// Build builds request to find comment by user id.
func (req *purchaseIDCommentRequest) Build(r *http.Request) error {
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

// Validate validates request to find comment by user id.
func (req *purchaseIDCommentRequest) Validate() error {
	switch {
	case req.ID < 1:
		return fmt.Errorf("not correct id")
	default:
		return nil
	}
}

func (c *commentRouter) findByPurchaseIDComment(w http.ResponseWriter, r *http.Request) {
	var req purchaseIDCommentRequest
	err := middleware.ParseRequest(r, &req)
	if err != nil {
		middleware.JSONError(w, err, http.StatusBadRequest)
		return
	}

	comments, err := c.services.Comment.FindByPurchaseID(req.PurchaseIDCommentRequest)
	if err != nil {
		middleware.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	if len(comments) == 0 {
		middleware.Empty(w, http.StatusNotFound)
		return
	}

	middleware.JSONReturn(w, http.StatusOK, comments)
}

type userPurchaseIDCommentRequest struct {
	model.UserPurchaseIDCommentRequest
}

// Build builds request to find comment by user id and purchase id.
func (req *userPurchaseIDCommentRequest) Build(r *http.Request) error {
	userID, ok := mux.Vars(r)["userID"]
	if !ok {
		return fmt.Errorf("no user id")
	}
	purchaseID, ok := mux.Vars(r)["purchaseID"]
	if !ok {
		return fmt.Errorf("no purchase id")
	}

	uID, err := strconv.Atoi(userID)
	if err != nil {
		return err
	}

	pID, err := strconv.Atoi(purchaseID)
	if err != nil {
		return err
	}

	req.UserID = uID
	req.PurchaseID = pID

	return nil
}

// Validate validates request to find comment by user id and purchase id.
func (req *userPurchaseIDCommentRequest) Validate() error {
	switch {
	case req.UserID < 1:
		return fmt.Errorf("not correct user id")
	case req.PurchaseID < 1:
		return fmt.Errorf("not correct purcahse id")
	default:
		return nil
	}
}

func (c *commentRouter) findByUserIDAndPurchaseIDComment(w http.ResponseWriter, r *http.Request) {
	var req userPurchaseIDCommentRequest
	err := middleware.ParseRequest(r, &req)
	if err != nil {
		middleware.JSONError(w, err, http.StatusBadRequest)
		return
	}

	comments, err := c.services.Comment.FindByUserIDAndPurchaseID(req.UserPurchaseIDCommentRequest)
	if err != nil {
		middleware.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	if len(comments) == 0 {
		middleware.Empty(w, http.StatusNotFound)
		return
	}

	middleware.JSONReturn(w, http.StatusOK, comments)
}

func (c *commentRouter) findAllComment(w http.ResponseWriter, r *http.Request) {
	comments, err := c.services.Comment.FindAll()
	if err != nil {
		middleware.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	if len(comments) == 0 {
		middleware.Empty(w, http.StatusNotFound)
		return
	}

	middleware.JSONReturn(w, http.StatusOK, comments)
}

type textCommentRequest struct {
	model.TextCommentRequest
}

// Build builds request to find comment by text.
func (req *textCommentRequest) Build(r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(&req.TextCommentRequest)
	if err != nil {
		return err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(r.Body)

	return nil
}

// Validate validates request to find comment by text.
func (req *textCommentRequest) Validate() error {
	switch {
	case req.Text == "":
		return fmt.Errorf("text is required")
	default:
		return nil
	}
}

func (c *commentRouter) findByTextComment(w http.ResponseWriter, r *http.Request) {
	var req textCommentRequest
	err := middleware.ParseRequest(r, &req)
	if err != nil {
		middleware.JSONError(w, err, http.StatusBadRequest)
		return
	}

	comments, err := c.services.Comment.FindByText(req.TextCommentRequest)
	if err != nil {
		middleware.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	if len(comments) == 0 {
		middleware.Empty(w, http.StatusNotFound)
		return
	}

	middleware.JSONReturn(w, http.StatusOK, comments)
}

type periodCommentRequest struct {
	model.PeriodCommentRequest
}

// Build builds request to find comment by date period.
func (req *periodCommentRequest) Build(r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(&req.PeriodCommentRequest)
	if err != nil {
		return err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(r.Body)

	return nil
}

// Validate validates request to find comment by date period.
func (req *periodCommentRequest) Validate() error {
	switch {
	case req.Start == time.Time{}:
		return fmt.Errorf("invalid start")
	case req.End == time.Time{}:
		return fmt.Errorf("invalid end")
	default:
		return nil
	}
}

func (c *commentRouter) findByPeriodComment(w http.ResponseWriter, r *http.Request) {
	var req periodCommentRequest
	err := middleware.ParseRequest(r, &req)
	if err != nil {
		middleware.JSONError(w, err, http.StatusBadRequest)
		return
	}

	comments, err := c.services.Comment.FindByPeriod(req.PeriodCommentRequest)
	if err != nil {
		middleware.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	if len(comments) == 0 {
		middleware.Empty(w, http.StatusNotFound)
		return
	}

	middleware.JSONReturn(w, http.StatusOK, comments)
}
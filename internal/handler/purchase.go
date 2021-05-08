package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/JesusG2000/hexsatisfaction/internal/model"
	"github.com/JesusG2000/hexsatisfaction/internal/service"
	"github.com/JesusG2000/hexsatisfaction/pkg/auth"
	"github.com/JesusG2000/hexsatisfaction/pkg/middleware"
	"github.com/gorilla/mux"
)

type purchaseRouter struct {
	*mux.Router
	services     *service.Services
	tokenManager auth.TokenManager
}

func newPurchase(services *service.Services, tokenManager auth.TokenManager) purchaseRouter {
	router := mux.NewRouter().PathPrefix(purchasePath).Subrouter()
	handler := purchaseRouter{
		router,
		services,
		tokenManager,
	}
	secure := router.PathPrefix("/api").Subrouter()
	secure.Use(handler.tokenManager.UserIdentity)

	secure.Path("/{id}").
		Methods(http.MethodGet).
		HandlerFunc(handler.findByIdPurchase)

	secure.Path("/last/user/{id}").
		Methods(http.MethodGet).
		HandlerFunc(handler.findLastByUserIdPurchase)

	secure.Path("/user/{id}").
		Methods(http.MethodGet).
		HandlerFunc(handler.findAllByUserIdPurchase)

	secure.Path("/last/").
		Methods(http.MethodGet).
		HandlerFunc(handler.findLast)

	secure.Path("/").
		Methods(http.MethodGet).
		HandlerFunc(handler.findAll)

	secure.Path("/user/{id}/file/{file}").
		Methods(http.MethodGet).
		HandlerFunc(handler.findByUserIdAndFileNamePurchase)

	secure.Path("/file/{file}").
		Methods(http.MethodGet).
		HandlerFunc(handler.findByFileNamePurchase)

	secure.Path("/").
		Methods(http.MethodPost).
		HandlerFunc(handler.createPurchase)

	secure.Path("/period/user/{id}").
		Methods(http.MethodPost).
		HandlerFunc(handler.findByUserIdAndPeriodPurchase)

	secure.Path("/after/user/{id}").
		Methods(http.MethodPost).
		HandlerFunc(handler.findByUserIdAfterDatePurchase)

	secure.Path("/before/user/{id}").
		Methods(http.MethodPost).
		HandlerFunc(handler.findByUserIdBeforeDatePurchase)

	secure.Path("/period").
		Methods(http.MethodPost).
		HandlerFunc(handler.findByPeriodPurchase)

	secure.Path("/after").
		Methods(http.MethodPost).
		HandlerFunc(handler.findAfterDatePurchase)

	secure.Path("/before").
		Methods(http.MethodPost).
		HandlerFunc(handler.findBeforeDatePurchase)

	secure.Path("/{id}").
		Methods(http.MethodDelete).
		HandlerFunc(handler.deletePurchase)

	return handler
}

type createPurchaseRequest struct {
	model.CreatePurchaseRequest
}

// Build builds request for create purchase.
func (req *createPurchaseRequest) Build(r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(&req.CreatePurchaseRequest)
	defer r.Body.Close()
	if err != nil {
		return err
	}

	return nil
}

// Validate validates request for create purchase.
func (req *createPurchaseRequest) Validate() error {
	switch {
	case req.UserId < 1:
		return fmt.Errorf("not correct user id")
	case req.Date == time.Time{}:
		return fmt.Errorf("date is required")
	case req.FileName == "":
		return fmt.Errorf("file name is required")
	default:
		return nil
	}
}

func (p *purchaseRouter) createPurchase(w http.ResponseWriter, r *http.Request) {
	var req createPurchaseRequest
	err := middleware.ParseRequest(r, &req)
	if err != nil {
		middleware.JSONError(w, err, http.StatusBadRequest)
		return
	}

	id, err := p.services.Purchase.Create(req.CreatePurchaseRequest)
	if err != nil {
		middleware.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	middleware.JSONReturn(w, http.StatusOK, strconv.Itoa(id))
}

type deletePurchaseRequest struct {
	model.DeletePurchaseRequest
}

// Build builds request to delete purchase.
func (req *deletePurchaseRequest) Build(r *http.Request) error {
	vId, ok := mux.Vars(r)["id"]
	if !ok {
		return fmt.Errorf("no id")
	}

	id, err := strconv.Atoi(vId)
	if err != nil {
		return err
	}

	req.Id = id

	return nil
}

// Validate validates request to delete purchase.
func (req *deletePurchaseRequest) Validate() error {
	switch {
	case req.Id < 1:
		return fmt.Errorf("not correct id")
	default:
		return nil
	}
}

func (p *purchaseRouter) deletePurchase(w http.ResponseWriter, r *http.Request) {
	var req deletePurchaseRequest
	err := middleware.ParseRequest(r, &req)
	if err != nil {
		middleware.JSONError(w, err, http.StatusBadRequest)
		return
	}

	id, err := p.services.Purchase.Delete(req.DeletePurchaseRequest)
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

type idPurchaseRequest struct {
	model.IdPurchaseRequest
}

// Build builds request to find purchase by id.
func (req *idPurchaseRequest) Build(r *http.Request) error {
	vId, ok := mux.Vars(r)["id"]
	if !ok {
		return fmt.Errorf("no id")
	}

	id, err := strconv.Atoi(vId)
	if err != nil {
		return err
	}

	req.Id = id

	return nil
}

// Validate validates request to find purchase by id.
func (req *idPurchaseRequest) Validate() error {
	switch {
	case req.Id < 1:
		return fmt.Errorf("not correct id")
	default:
		return nil
	}
}

func (p *purchaseRouter) findByIdPurchase(w http.ResponseWriter, r *http.Request) {
	var req idPurchaseRequest
	err := middleware.ParseRequest(r, &req)
	if err != nil {
		middleware.JSONError(w, err, http.StatusBadRequest)
		return
	}

	purchase, err := p.services.Purchase.FindById(req.IdPurchaseRequest)
	if err != nil {
		middleware.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	if purchase.ID < 1 {
		middleware.Empty(w, http.StatusNotFound)
		return
	}

	middleware.JSONReturn(w, http.StatusOK, purchase)
}

type lastUserIdPurchaseRequest struct {
	model.UserIdPurchaseRequest
}

// Build builds request to find last purchase by user id.
func (req *lastUserIdPurchaseRequest) Build(r *http.Request) error {
	vId, ok := mux.Vars(r)["id"]
	if !ok {
		return fmt.Errorf("no id")
	}

	id, err := strconv.Atoi(vId)
	if err != nil {
		return err
	}

	req.Id = id

	return nil
}

// Validate validates request to find last purchase by user id.
func (req *lastUserIdPurchaseRequest) Validate() error {
	switch {
	case req.Id < 1:
		return fmt.Errorf("not correct id")
	default:
		return nil
	}
}

func (p *purchaseRouter) findLastByUserIdPurchase(w http.ResponseWriter, r *http.Request) {
	var req lastUserIdPurchaseRequest
	err := middleware.ParseRequest(r, &req)
	if err != nil {
		middleware.JSONError(w, err, http.StatusBadRequest)
		return
	}

	purchase, err := p.services.Purchase.FindLastByUserId(req.UserIdPurchaseRequest)
	if err != nil {
		middleware.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	if purchase.ID < 1 {
		middleware.Empty(w, http.StatusNotFound)
		return
	}

	middleware.JSONReturn(w, http.StatusOK, purchase)
}

type userIdPurchaseRequest struct {
	model.UserIdPurchaseRequest
}

// Build builds request to find all purchases by user id.
func (req *userIdPurchaseRequest) Build(r *http.Request) error {
	vId, ok := mux.Vars(r)["id"]
	if !ok {
		return fmt.Errorf("no id")
	}

	id, err := strconv.Atoi(vId)
	if err != nil {
		return err
	}

	req.Id = id

	return nil
}

// Validate validates request to find all purchases by user id.
func (req *userIdPurchaseRequest) Validate() error {
	switch {
	case req.Id < 1:
		return fmt.Errorf("not correct id")
	default:
		return nil
	}
}

func (p *purchaseRouter) findAllByUserIdPurchase(w http.ResponseWriter, r *http.Request) {
	var req userIdPurchaseRequest
	err := middleware.ParseRequest(r, &req)
	if err != nil {
		middleware.JSONError(w, err, http.StatusBadRequest)
		return
	}

	purchases, err := p.services.Purchase.FindAllByUserId(req.UserIdPurchaseRequest)
	if err != nil {
		middleware.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	if len(purchases) == 0 {
		middleware.Empty(w, http.StatusNotFound)
		return
	}

	middleware.JSONReturn(w, http.StatusOK, purchases)
}

type userIdPeriodPurchaseRequest struct {
	model.UserIdPeriodPurchaseRequest
}

// Build builds request to find all purchases by user id and date period.
func (req *userIdPeriodPurchaseRequest) Build(r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(&req.UserIdPeriodPurchaseRequest)
	defer r.Body.Close()
	if err != nil {
		return err
	}

	vId, ok := mux.Vars(r)["id"]
	if !ok {
		return fmt.Errorf("no id")
	}

	id, err := strconv.Atoi(vId)
	if err != nil {
		return err
	}

	req.Id = id

	return nil
}

// Validate validates request to find all purchases by user id and date period.
func (req *userIdPeriodPurchaseRequest) Validate() error {
	switch {
	case req.Id < 1:
		return fmt.Errorf("not correct id")
	case req.Start == time.Time{}:
		return fmt.Errorf("start date is required")
	case req.End == time.Time{}:
		return fmt.Errorf("end date is required")
	default:
		return nil
	}
}

func (p *purchaseRouter) findByUserIdAndPeriodPurchase(w http.ResponseWriter, r *http.Request) {
	var req userIdPeriodPurchaseRequest
	err := middleware.ParseRequest(r, &req)
	if err != nil {
		middleware.JSONError(w, err, http.StatusBadRequest)
		return
	}

	purchases, err := p.services.Purchase.FindByUserIdAndPeriod(req.UserIdPeriodPurchaseRequest)
	if err != nil {
		middleware.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	if len(purchases) == 0 {
		middleware.Empty(w, http.StatusNotFound)
		return
	}

	middleware.JSONReturn(w, http.StatusOK, purchases)
}

type userIdAfterDatePurchaseRequest struct {
	model.UserIdAfterDatePurchaseRequest
}

// Build builds request to find all purchases by user id after date.
func (req *userIdAfterDatePurchaseRequest) Build(r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(&req.UserIdAfterDatePurchaseRequest)
	defer r.Body.Close()
	if err != nil {
		return err
	}

	vId, ok := mux.Vars(r)["id"]
	if !ok {
		return fmt.Errorf("no id")
	}

	id, err := strconv.Atoi(vId)
	if err != nil {
		return err
	}

	req.Id = id

	return nil
}

// Validate validates request to find all purchases by user id after date.
func (req *userIdAfterDatePurchaseRequest) Validate() error {
	switch {
	case req.Id < 1:
		return fmt.Errorf("not correct id")
	case req.Start == time.Time{}:
		return fmt.Errorf("start date is required")
	default:
		return nil
	}
}

func (p *purchaseRouter) findByUserIdAfterDatePurchase(w http.ResponseWriter, r *http.Request) {
	var req userIdAfterDatePurchaseRequest
	err := middleware.ParseRequest(r, &req)
	if err != nil {
		middleware.JSONError(w, err, http.StatusBadRequest)
		return
	}

	purchases, err := p.services.Purchase.FindByUserIdAfterDate(req.UserIdAfterDatePurchaseRequest)
	if err != nil {
		middleware.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	if len(purchases) == 0 {
		middleware.Empty(w, http.StatusNotFound)
		return
	}

	middleware.JSONReturn(w, http.StatusOK, purchases)
}

type userIdBeforeDatePurchaseRequest struct {
	model.UserIdBeforeDatePurchaseRequest
}

// Build builds request to find all purchases by user id before date.
func (req *userIdBeforeDatePurchaseRequest) Build(r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(&req.UserIdBeforeDatePurchaseRequest)
	defer r.Body.Close()
	if err != nil {
		return err
	}

	vId, ok := mux.Vars(r)["id"]
	if !ok {
		return fmt.Errorf("no id")
	}

	id, err := strconv.Atoi(vId)
	if err != nil {
		return err
	}

	req.Id = id

	return nil
}

// Validate validates request to find all purchases by user id before date.
func (req *userIdBeforeDatePurchaseRequest) Validate() error {
	switch {
	case req.Id < 1:
		return fmt.Errorf("not correct id")
	case req.End == time.Time{}:
		return fmt.Errorf("end date is required")
	default:
		return nil
	}
}

func (p *purchaseRouter) findByUserIdBeforeDatePurchase(w http.ResponseWriter, r *http.Request) {
	var req userIdBeforeDatePurchaseRequest
	err := middleware.ParseRequest(r, &req)
	if err != nil {
		middleware.JSONError(w, err, http.StatusBadRequest)
		return
	}

	purchases, err := p.services.Purchase.FindByUserIdBeforeDate(req.UserIdBeforeDatePurchaseRequest)
	if err != nil {
		middleware.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	if len(purchases) == 0 {
		middleware.Empty(w, http.StatusNotFound)
		return
	}

	middleware.JSONReturn(w, http.StatusOK, purchases)
}

type userIdFileNamePurchaseRequest struct {
	model.UserIdFileNamePurchaseRequest
}

// Build builds request to find all purchases by user id and file name.
func (req *userIdFileNamePurchaseRequest) Build(r *http.Request) error {

	vId, ok := mux.Vars(r)["id"]
	if !ok {
		return fmt.Errorf("no id")
	}

	id, err := strconv.Atoi(vId)
	if err != nil {
		return err
	}
	name, ok := mux.Vars(r)["file"]
	if !ok {
		return fmt.Errorf("no file name")
	}

	req.Id = id
	req.FileName = name

	return nil
}

// Validate validates request to find all purchases by user id and file name.
func (req *userIdFileNamePurchaseRequest) Validate() error {
	switch {
	case req.Id < 1:
		return fmt.Errorf("not correct id")
	case req.FileName == "":
		return fmt.Errorf("file name is required")
	default:
		return nil
	}
}

func (p *purchaseRouter) findByUserIdAndFileNamePurchase(w http.ResponseWriter, r *http.Request) {
	var req userIdFileNamePurchaseRequest
	err := middleware.ParseRequest(r, &req)
	if err != nil {
		middleware.JSONError(w, err, http.StatusBadRequest)
		return
	}

	purchases, err := p.services.Purchase.FindByUserIdAndFileName(req.UserIdFileNamePurchaseRequest)
	if err != nil {
		middleware.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	if len(purchases) == 0 {
		middleware.Empty(w, http.StatusNotFound)
		return
	}

	middleware.JSONReturn(w, http.StatusOK, purchases)
}

func (p *purchaseRouter) findLast(w http.ResponseWriter, r *http.Request) {
	purchase, err := p.services.Purchase.FindLast()
	if err != nil {
		middleware.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	if purchase.ID < 1 {
		middleware.Empty(w, http.StatusNotFound)
		return
	}

	middleware.JSONReturn(w, http.StatusOK, purchase)
}

func (p *purchaseRouter) findAll(w http.ResponseWriter, r *http.Request) {
	purchases, err := p.services.Purchase.FindAll()
	if err != nil {
		middleware.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	if len(purchases) == 0 {
		middleware.Empty(w, http.StatusNotFound)
		return
	}

	middleware.JSONReturn(w, http.StatusOK, purchases)
}

type periodPurchaseRequest struct {
	model.PeriodPurchaseRequest
}

// Build builds request to find all purchases by date period.
func (req *periodPurchaseRequest) Build(r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(&req.PeriodPurchaseRequest)
	defer r.Body.Close()
	if err != nil {
		return err
	}

	return nil
}

// Validate validates request to find all purchases by date period.
func (req *periodPurchaseRequest) Validate() error {
	switch {
	case req.Start == time.Time{}:
		return fmt.Errorf("start date is required")
	case req.End == time.Time{}:
		return fmt.Errorf("end date is required")
	default:
		return nil
	}
}

func (p *purchaseRouter) findByPeriodPurchase(w http.ResponseWriter, r *http.Request) {
	var req periodPurchaseRequest
	err := middleware.ParseRequest(r, &req)
	if err != nil {
		middleware.JSONError(w, err, http.StatusBadRequest)
		return
	}

	purchases, err := p.services.Purchase.FindByPeriod(req.PeriodPurchaseRequest)
	if err != nil {
		middleware.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	if len(purchases) == 0 {
		middleware.Empty(w, http.StatusNotFound)
		return
	}

	middleware.JSONReturn(w, http.StatusOK, purchases)
}

type afterDatePurchaseRequest struct {
	model.AfterDatePurchaseRequest
}

// Build builds request to find all purchases after date.
func (req *afterDatePurchaseRequest) Build(r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(&req.AfterDatePurchaseRequest)
	defer r.Body.Close()
	if err != nil {
		return err
	}

	return nil
}

// Validate validates request to find all purchases after date.
func (req *afterDatePurchaseRequest) Validate() error {
	switch {
	case req.Start == time.Time{}:
		return fmt.Errorf("start date is required")
	default:
		return nil
	}
}

func (p *purchaseRouter) findAfterDatePurchase(w http.ResponseWriter, r *http.Request) {
	var req afterDatePurchaseRequest
	err := middleware.ParseRequest(r, &req)
	if err != nil {
		middleware.JSONError(w, err, http.StatusBadRequest)
		return
	}

	purchases, err := p.services.Purchase.FindAfterDate(req.AfterDatePurchaseRequest)
	if err != nil {
		middleware.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	if len(purchases) == 0 {
		middleware.Empty(w, http.StatusNotFound)
		return
	}

	middleware.JSONReturn(w, http.StatusOK, purchases)
}

type beforeDatePurchaseRequest struct {
	model.BeforeDatePurchaseRequest
}

// Build builds request to find all purchases before date.
func (req *beforeDatePurchaseRequest) Build(r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(&req.BeforeDatePurchaseRequest)
	defer r.Body.Close()
	if err != nil {
		return err
	}

	return nil
}

// Validate validates request to find all purchases before date.
func (req *beforeDatePurchaseRequest) Validate() error {
	switch {
	case req.End == time.Time{}:
		return fmt.Errorf("end date is required")
	default:
		return nil
	}
}

func (p *purchaseRouter) findBeforeDatePurchase(w http.ResponseWriter, r *http.Request) {
	var req beforeDatePurchaseRequest
	err := middleware.ParseRequest(r, &req)
	if err != nil {
		middleware.JSONError(w, err, http.StatusBadRequest)
		return
	}

	purchases, err := p.services.Purchase.FindBeforeDate(req.BeforeDatePurchaseRequest)
	if err != nil {
		middleware.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	if len(purchases) == 0 {
		middleware.Empty(w, http.StatusIMUsed)
		return
	}

	middleware.JSONReturn(w, http.StatusOK, purchases)
}

type fileNamePurchaseRequest struct {
	model.FileNamePurchaseRequest
}

// Build builds request to find all purchases by file name.
func (req *fileNamePurchaseRequest) Build(r *http.Request) error {
	name, ok := mux.Vars(r)["file"]
	if !ok {
		return fmt.Errorf("no file name")
	}

	req.FileName = name

	return nil
}

// Validate validates request to find all purchases by file name.
func (req *fileNamePurchaseRequest) Validate() error {
	switch {
	case req.FileName == "":
		return fmt.Errorf("file name is required")
	default:
		return nil
	}
}

func (p *purchaseRouter) findByFileNamePurchase(w http.ResponseWriter, r *http.Request) {
	var req fileNamePurchaseRequest
	err := middleware.ParseRequest(r, &req)
	if err != nil {
		middleware.JSONError(w, err, http.StatusBadRequest)
		return
	}

	purchases, err := p.services.Purchase.FindByFileName(req.FileNamePurchaseRequest)
	if err != nil {
		middleware.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	if len(purchases) == 0 {
		middleware.Empty(w, http.StatusNotFound)
		return
	}

	middleware.JSONReturn(w, http.StatusOK, purchases)
}

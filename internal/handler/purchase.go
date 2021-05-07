package handler

import (
	"net/http"

	"github.com/JesusG2000/hexsatisfaction/internal/model"
	"github.com/JesusG2000/hexsatisfaction/internal/service"
	"github.com/JesusG2000/hexsatisfaction/pkg/auth"
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

	return handler
}

type createPurchaseRequest struct {
	model.CreatePurchaseRequest
}

// Build builds request for create purchase.
func (req *createPurchaseRequest) Build(r *http.Request) error {

}

// Validate validates request for create purchase.
func (req *createPurchaseRequest) Validate() error {

}

func (p *purchaseRouter) createPurchase(w http.ResponseWriter, r *http.Request) {

}

type deletePurchaseRequest struct {
	model.DeletePurchaseRequest
}

// Build builds request to delete purchase.
func (req *deletePurchaseRequest) Build(r *http.Request) error {

}

// Validate validates request to delete purchase.
func (req *deletePurchaseRequest) Validate() error {

}

func (p *purchaseRouter) deletePurchase(w http.ResponseWriter, r *http.Request) {

}

type idPurchaseRequest struct {
	model.IdPurchaseRequest
}

// Build builds request to find purchase by id.
func (req *idPurchaseRequest) Build(r *http.Request) error {

}

// Validate validates request to find purchase by id.
func (req *idPurchaseRequest) Validate() error {

}

func (p *purchaseRouter) findByIdPurchase(w http.ResponseWriter, r *http.Request) {

}

type UserIdPurchaseRequest struct {
	model.UserIdPurchaseRequest
}

// Build builds request to find last purchase by user id.
func (req *UserIdPurchaseRequest) Build(r *http.Request) error {

}

// Validate validates request to find last purchase by user id.
func (req *UserIdPurchaseRequest) Validate() error {

}

func (p *purchaseRouter) findLastByUserIdPurchase(w http.ResponseWriter, r *http.Request) {

}

type userIdPurchaseRequest struct {
	model.UserIdPurchaseRequest
}

// Build builds request to find all purchases by user id.
func (req *userIdPurchaseRequest) Build(r *http.Request) error {

}

// Validate validates request to find all purchases by user id.
func (req *userIdPurchaseRequest) Validate() error {

}

func (p *purchaseRouter) findAllByUserIdPurchase(w http.ResponseWriter, r *http.Request) {

}

type userIdPeriodPurchaseRequest struct {
	model.UserIdPeriodPurchaseRequest
}

// Build builds request to find all purchases by user id and date period.
func (req *userIdPeriodPurchaseRequest) Build(r *http.Request) error {

}

// Validate validates request to find all purchases by user id and date period.
func (req *userIdPeriodPurchaseRequest) Validate() error {

}

func (p *purchaseRouter) findByUserIdAndPeriodPurchase(w http.ResponseWriter, r *http.Request) {

}

type userIdAfterDatePurchaseRequest struct {
	model.UserIdAfterDatePurchaseRequest
}

// Build builds request to find all purchases by user id after date.
func (req *userIdAfterDatePurchaseRequest) Build(r *http.Request) error {

}

// Validate validates request to find all purchases by user id after date.
func (req *userIdAfterDatePurchaseRequest) Validate() error {

}

func (p *purchaseRouter) findByUserIdAfterDatePurchase(w http.ResponseWriter, r *http.Request) {

}

type userIdBeforeDatePurchaseRequest struct {
	model.UserIdBeforeDatePurchaseRequest
}

// Build builds request to find all purchases by user id before date.
func (req *userIdBeforeDatePurchaseRequest) Build(r *http.Request) error {

}

// Validate validates request to find all purchases by user id before date.
func (req *userIdBeforeDatePurchaseRequest) Validate() error {

}

func (p *purchaseRouter) findByUserIdBeforeDatePurchase(w http.ResponseWriter, r *http.Request) {

}

type userIdFileNamePurchaseRequest struct {
	model.UserIdFileNamePurchaseRequest
}

// Build builds request to find all purchases by user id and file name.
func (req *userIdFileNamePurchaseRequest) Build(r *http.Request) error {

}

// Validate validates request to find all purchases by user id and file name.
func (req *userIdFileNamePurchaseRequest) Validate() error {

}

func (p *purchaseRouter) findByUserIdAndFileNamePurchase(w http.ResponseWriter, r *http.Request) {

}

func (p *purchaseRouter) findLast(w http.ResponseWriter, r *http.Request) (*model.Purchase, error) {

}

func (p *purchaseRouter) findAll(w http.ResponseWriter, r *http.Request) ([]model.Purchase, error) {

}

type periodPurchaseRequest struct {
	model.PeriodPurchaseRequest
}

// Build builds request to find all purchases by date period.
func (req *periodPurchaseRequest) Build(r *http.Request) error {

}

// Validate validates request to find all purchases by date period.
func (req *periodPurchaseRequest) Validate() error {

}

func (p *purchaseRouter) findByPeriodPurchase(w http.ResponseWriter, r *http.Request) {

}

type afterDatePurchaseRequest struct {
	model.AfterDatePurchaseRequest
}

// Build builds request to find all purchases after date.
func (req *afterDatePurchaseRequest) Build(r *http.Request) error {

}

// Validate validates request to find all purchases after date.
func (req *afterDatePurchaseRequest) Validate() error {

}

func (p *purchaseRouter) findAfterDatePurchase(w http.ResponseWriter, r *http.Request) {

}

type beforeDatePurchaseRequest struct {
	model.BeforeDatePurchaseRequest
}

// Build builds request to find all purchases before date.
func (req *beforeDatePurchaseRequest) Build(r *http.Request) error {

}

// Validate validates request to find all purchases before date.
func (req *beforeDatePurchaseRequest) Validate() error {

}

func (p *purchaseRouter) findBeforeDatePurchase(w http.ResponseWriter, r *http.Request) {

}

type fileNamePurchaseRequest struct {
	model.FileNamePurchaseRequest
}

// Build builds request to find all purchases by file name.
func (req *fileNamePurchaseRequest) Build(r *http.Request) error {

}

// Validate validates request to find all purchases by file name.
func (req *fileNamePurchaseRequest) Validate() error {

}

func (p *purchaseRouter) findByFileNamePurchase(w http.ResponseWriter, r *http.Request) {

}

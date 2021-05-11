package handler

import (
	"github.com/JesusG2000/hexsatisfaction/internal/service"
	"github.com/JesusG2000/hexsatisfaction/pkg/auth"
	"github.com/gorilla/mux"
)

const userPath = "/user"
const purchasePath = "/purchase"

// API represents a structure with APIs.
type API struct {
	*mux.Router
}

// NewHandler creates and serves endpoints of API.
func NewHandler(services *service.Services, tokenManager auth.TokenManager) *API {
	api := API{
		mux.NewRouter(),
	}
	api.PathPrefix(userPath).Handler(newUser(services, tokenManager))
	api.PathPrefix(purchasePath).Handler(newPurchase(services, tokenManager))

	return &api
}

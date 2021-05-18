package model

import "time"

type (

	// RegisterUserRequest represents a request for user registration.
	RegisterUserRequest struct {
		// required: true
		Login string `json:"login"`
		// required: true
		Password string `json:"password"`
	}

	// LoginUserRequest represents a request for user login.
	LoginUserRequest struct {
		// required: true
		Login string `json:"login"`
		// required: true
		Password string `json:"password"`
	}
)

type (

	// CreatePurchaseRequest represents a request to create purchase.
	CreatePurchaseRequest struct {
		// required: true
		UserID int `json:"userID"`
		// required: true
		Date time.Time `json:"date"`
		// required: true
		FileName string `json:"fileName"`
	}

	// IDPurchaseRequest represents a request to find the purchase by id.
	IDPurchaseRequest struct {
		// required: true
		ID int `json:"-"`
	}

	// DeletePurchaseRequest represents a request to delete purchase.
	DeletePurchaseRequest = struct {
		// required: true
		ID int `json:"-"`
	}

	// UserIDPurchaseRequest represents a request to find last added purchase by user id.
	UserIDPurchaseRequest = struct {
		// required: true
		ID int `json:"-"`
	}

	// UserIDPeriodPurchaseRequest represents a request to find all purchases by user id and date period.
	UserIDPeriodPurchaseRequest struct {
		// required: true
		ID int `json:"-"`
		// required: true
		Start time.Time `json:"start"`
		// required: true
		End time.Time `json:"end"`
	}

	// UserIDAfterDatePurchaseRequest represents a request to find all purchases by user id after date.
	UserIDAfterDatePurchaseRequest struct {
		// required: true
		ID int `json:"-"`
		// required: true
		Start time.Time `json:"start"`
	}

	// UserIDBeforeDatePurchaseRequest represents a request to find all purchases by user id before date.
	UserIDBeforeDatePurchaseRequest struct {
		// required: true
		ID int `json:"-"`
		// required: true
		End time.Time `json:"end"`
	}

	// UserIDFileNamePurchaseRequest represents a request to find all purchases by user id and file name.
	UserIDFileNamePurchaseRequest struct {
		// required: true
		ID int `json:"-"`
		// required: true
		FileName string `json:"fileName"`
	}

	// PeriodPurchaseRequest represents a request to find all purchases by date period.
	PeriodPurchaseRequest struct {
		// required: true
		Start time.Time `json:"start"`
		// required: true
		End time.Time `json:"end"`
	}

	// AfterDatePurchaseRequest represents a request to find all purchases after date.
	AfterDatePurchaseRequest struct {
		// required: true
		Start time.Time `json:"start"`
	}

	// BeforeDatePurchaseRequest represents a request to find all purchases before date.
	BeforeDatePurchaseRequest struct {
		// required: true
		End time.Time `json:"end"`
	}

	// FileNamePurchaseRequest represents a request to find all purchases by file name.
	FileNamePurchaseRequest struct {
		// required: true
		FileName string `json:"-"`
	}
)

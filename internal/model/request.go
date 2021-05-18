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

type (

	// CreateCommentRequest represents a request to create comment.
	CreateCommentRequest struct {
		// required: true
		UserID int `json:"userID"`
		// required: true
		PurchaseID int `json:"purchaseID"`
		// required: true
		Date time.Time `json:"Date"`
		// required: true
		Text string `json:"Text"`
	}

	// UpdateCommentRequest represents a request to update comment.
	UpdateCommentRequest struct {
		// required: true
		ID int `json:"-"`
		// required: true
		UserID int `json:"userID"`
		// required: true
		PurchaseID int `json:"purchaseID"`
		// required: true
		Date time.Time `json:"date"`
		// required: true
		Text string `json:"text"`
	}

	// DeleteCommentRequest represents a request to delete comment.
	DeleteCommentRequest struct {
		// required: true
		ID int `json:"-"`
	}

	// IDCommentRequest represents a request to find comment by id.
	IDCommentRequest struct {
		// required: true
		ID int `json:"-"`
	}

	// UserIDCommentRequest represents a request to find comments by user id.
	UserIDCommentRequest struct {
		// required: true
		ID int `json:"-"`
	}

	// PurchaseIDCommentRequest represents a request to find comments by purchase id.
	PurchaseIDCommentRequest struct {
		// required: true
		ID int `json:"-"`
	}

	// UserPurchaseIDCommentRequest represents a request to find comments by purchase and user ids.
	UserPurchaseIDCommentRequest struct {
		// required: true
		UserID     int `json:"-"`
		PurchaseID int `json:"-"`
	}

	// TextCommentRequest represents a request to find comments by text.
	TextCommentRequest struct {
		// required: true
		Text string `json:"text"`
	}

	// PeriodCommentRequest represents a request to find comments by date period.
	PeriodCommentRequest struct {
		// required: true
		Start time.Time `json:"start"`
		// required: true
		End time.Time `json:"end"`
	}
)

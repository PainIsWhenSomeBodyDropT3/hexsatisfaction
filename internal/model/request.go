package model

import "time"

type (

	// RegisterUserRequest represents a request for user registration.
	RegisterUserRequest struct {
		Login    string
		Password string
	}

	// LoginUserRequest represents a request for user login.
	LoginUserRequest struct {
		Login    string
		Password string
	}
)

type (

	// CreatePurchaseRequest represents a request to create purchase.
	CreatePurchaseRequest struct {
		UserId   int
		Date     time.Time
		FileName string
	}

	// IdPurchaseRequest represents a request to find the purchase by id.
	IdPurchaseRequest struct {
		Id int
	}

	// DeletePurchaseRequest represents a request to delete purchase.
	DeletePurchaseRequest = struct {
		Id int
	}

	// UserIdPurchaseRequest represents a request to find last added purchase by user id.
	UserIdPurchaseRequest = struct {
		Id int
	}

	// UserIdPeriodPurchaseRequest represents a request to find all purchases by user id and date period.
	UserIdPeriodPurchaseRequest struct {
		Id    int
		Start time.Time
		End   time.Time
	}

	// UserIdAfterDatePurchaseRequest represents a request to find all purchases by user id after date.
	UserIdAfterDatePurchaseRequest struct {
		Id    int
		Start time.Time
	}

	// UserIdBeforeDatePurchaseRequest represents a request to find all purchases by user id before date.
	UserIdBeforeDatePurchaseRequest struct {
		Id  int
		End time.Time
	}

	// UserIdFileNamePurchaseRequest represents a request to find all purchases by user id and file name.
	UserIdFileNamePurchaseRequest struct {
		Id       int
		FileName string
	}

	// PeriodPurchaseRequest represents a request to find all purchases by date period.
	PeriodPurchaseRequest struct {
		Start time.Time
		End   time.Time
	}

	// AfterDatePurchaseRequest represents a request to find all purchases after date.
	AfterDatePurchaseRequest struct {
		Start time.Time
	}

	// BeforeDatePurchaseRequest represents a request to find all purchases before date.
	BeforeDatePurchaseRequest struct {
		End time.Time
	}

	// FileNamePurchaseRequest represents a request to find all purchases by file name.
	FileNamePurchaseRequest struct {
		FileName string
	}
)

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

	// DeletePurchaseRequest represents a request to delete purchase.
	DeletePurchaseRequest struct {
		Id int
	}

	// FindByIdPurchaseRequest represents a request to find the purchase by id.
	FindByIdPurchaseRequest struct {
		Id int
	}

	// FindLastByUserIdPurchaseRequest represents a request to find last added purchase by user id.
	FindLastByUserIdPurchaseRequest struct {
		Id int
	}

	// FindAllByUserIdPurchaseRequest represents a request to find all purchases by user id.
	FindAllByUserIdPurchaseRequest struct {
		Id int
	}

	// FindByUserIdAndPeriodPurchaseRequest represents a request to find all purchases by user id and date period.
	FindByUserIdAndPeriodPurchaseRequest struct {
		Id    int
		Start time.Time
		End   time.Time
	}

	// FindByUserIdAfterDatePurchaseRequest represents a request to find all purchases by user id after date.
	FindByUserIdAfterDatePurchaseRequest struct {
		Id    int
		Start time.Time
	}

	// FindByUserIdBeforeDatePurchaseRequest represents a request to find all purchases by user id before date.
	FindByUserIdBeforeDatePurchaseRequest struct {
		Id  int
		End time.Time
	}

	// FindByUserIdAndFileNamePurchaseRequest represents a request to find all purchases by user id and file name.
	FindByUserIdAndFileNamePurchaseRequest struct {
		Id       int
		FileName string
	}

	// FindByPeriodPurchaseRequest represents a request to find all purchases by date period.
	FindByPeriodPurchaseRequest struct {
		Start time.Time
		End   time.Time
	}

	// FindAfterDatePurchaseRequest represents a request to find all purchases after date.
	FindAfterDatePurchaseRequest struct {
		Start time.Time
	}

	// FindBeforeDatePurchaseRequest represents a request to find all purchases before date.
	FindBeforeDatePurchaseRequest struct {
		End time.Time
	}

	// FindByFileNamePurchaseRequest represents a request to find all purchases by file name.
	FindByFileNamePurchaseRequest struct {
		FileName string
	}
)

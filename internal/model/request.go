package model

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
	// CreateAuthorRequest represents a request to create author.
	CreateAuthorRequest struct {
		// required: true
		Name string `json:"name"`
		// required: true
		Age int `json:"age"`
		// required: true
		Description string `json:"description"`
		// required: true
		UserID int `json:"userID"`
	}

	// UpdateAuthorRequest represents a request to update author.
	UpdateAuthorRequest struct {
		// required: true
		ID int `json:"-"`
		// required: true
		Name string `json:"name"`
		// required: true
		Age int `json:"age"`
		// required: true
		Description string `json:"description"`
		// required: true
		UserID int `json:"userID"`
	}

	// DeleteAuthorRequest represents a request to delete author.
	DeleteAuthorRequest struct {
		// required: true
		ID int `json:"-"`
	}

	// IDAuthorRequest represents a request to find author by id.
	IDAuthorRequest struct {
		// required: true
		ID int `json:"-"`
	}

	// UserIDAuthorRequest represents a request to find author by user id.
	UserIDAuthorRequest struct {
		// required: true
		ID int `json:"-"`
	}

	// NameAuthorRequest represents a request to find authors by name.
	NameAuthorRequest struct {
		// required: true
		Name string `json:"-"`
	}
)

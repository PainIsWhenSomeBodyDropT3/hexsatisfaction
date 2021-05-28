package model

import "time"

// File represents file model.
type File struct {
	ID          int       `json:"id,omitempty"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Size        int       `json:"size"`
	Path        string    `json:"path"`
	AddDate     time.Time `json:"addDate"`
	UpdateDate  time.Time `json:"updateDate"`
	Actual      bool      `json:"actual"`
	AuthorID    int       `json:"authorID"`
}

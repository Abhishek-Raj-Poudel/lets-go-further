package data

import (
	"database/sql"
	"errors"
)

//Define a custom ErrRecordNotFound error. We'll return this from our GET() method when
//looking pu a movie that doesn't exit in our database.

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

// here we will add other models
type Models struct {
	Movies MovieModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Movies: MovieModel{DB: db},
	}
}

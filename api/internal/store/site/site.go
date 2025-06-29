package site

import (
	"database/sql"
	"time"
)

type Store struct {
	db           *sql.DB
	queryTimeout time.Duration
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:           db,
		queryTimeout: 5 * time.Second,
	}
}

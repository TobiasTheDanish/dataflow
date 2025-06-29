package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type Config struct {
	Url   string `goenv:"TURSO_AUTH_URL,required"`
	Token string `goenv:"TURSO_AUTH_TOKEN,required"`
}

func NewContext(ctx context.Context, config Config) (*sql.DB, error) {
	fullPath := fmt.Sprintf("%s?authToken=%s", config.Url, config.Token)

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	db, err := sql.Open("libsql", fullPath)
	if err != nil {
		return nil, err
	}

	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}

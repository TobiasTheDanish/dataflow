package store

import (
	"api/internal/store/site"
	"context"
	"database/sql"
)

type Store struct {
	Sites SiteStore
}

func New(db *sql.DB) *Store {
	return &Store{
		Sites: site.NewStore(db),
	}
}

type SiteStore interface {
	All(ctx context.Context) ([]site.Site, error)
	CreateHttp(ctx context.Context, data site.NewHttpSite) (res site.HttpSite, err error)
	CreateFtp(ctx context.Context, data site.NewFtpSite) (res site.FtpSite, err error)
}

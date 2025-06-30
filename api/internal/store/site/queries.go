package site

import (
	"context"
	"encoding/json"
	"log/slog"
)

func (s *Store) All(ctx context.Context) ([]Site, error) {
	ctx, cancel := context.WithTimeout(ctx, s.queryTimeout)
	defer cancel()

	stmt := `
	SELECT id, name, conn_type, conn_config FROM df_site
	`

	rows, err := s.db.QueryContext(ctx, stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	sites := make([]Site, 0, 0)
	for rows.Next() {
		var site Site
		if err := rows.Scan(&site.Id, &site.Name, &site.Type, &site.Config); err != nil {
			return nil, err
		}

		sites = append(sites, site)
	}

	return sites, nil
}

func (s *Store) CreateHttp(ctx context.Context, data NewHttpSite) (res HttpSite, err error) {
	ctx, cancel := context.WithTimeout(ctx, s.queryTimeout)
	defer cancel()

	stmt := `
	INSERT INTO df_site (name, conn_type, conn_config)
	VALUES (?, ?, ?)
	RETURNING id
	`

	jsonConfig, err := json.Marshal(data.Config)
	if err != nil {
		return
	}

	slog.Default().Info("Serialized config", "jsonConfig", string(jsonConfig))

	row := s.db.QueryRowContext(
		ctx,
		stmt,
		data.Name,
		"http",
		string(jsonConfig),
	)

	if err = row.Scan(&res.Id); err != nil {
		return
	}

	res.Name = data.Name
	res.Config = data.Config

	return
}

func (s *Store) CreateFtp(ctx context.Context, data NewFtpSite) (res FtpSite, err error) {
	ctx, cancel := context.WithTimeout(ctx, s.queryTimeout)
	defer cancel()

	stmt := `
	INSERT INTO df_site (name, conn_type, conn_config)
	VALUES (?, ?, ?)
	RETURNING id
	`

	jsonConfig, err := json.Marshal(data.Config)
	if err != nil {
		return
	}

	slog.Default().Info("Serialized config", "jsonConfig", string(jsonConfig))

	row := s.db.QueryRowContext(
		ctx,
		stmt,
		data.Name,
		"ftp",
		string(jsonConfig),
	)

	if err = row.Scan(&res.Id); err != nil {
		return
	}

	res.Name = data.Name
	res.Config = data.Config

	return
}

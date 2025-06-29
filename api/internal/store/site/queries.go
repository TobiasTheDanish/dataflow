package site

import (
	"context"
	"encoding/json"
	"log/slog"
)

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

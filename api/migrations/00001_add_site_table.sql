-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS df_site (
	id INTEGER PRIMARY KEY,
	name TEXT NOT NULL,
	conn_type TEXT NOT NULL,
	conn_config TEXT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS df_site;
-- +goose StatementEnd

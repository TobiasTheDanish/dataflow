package main

import (
	"api/internal/database"
	"api/internal/server"
)

type Config struct {
	Server server.Config
	Db     database.Config
}

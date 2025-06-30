package server

import (
	"api/internal/store"
	"fmt"
	"net/http"
	"time"
)

type Config struct {
	Port int `goenv:"PORT,required"`
}

type Server struct {
	port  int
	store *store.Store
}

func NewServer(config Config, store *store.Store) *http.Server {
	NewServer := &Server{
		port:  config.Port,
		store: store,
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}

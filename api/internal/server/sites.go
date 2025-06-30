package server

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Server) getAllSites(c echo.Context) error {
	sites, err := s.store.Sites.All(c.Request().Context())
	if err != nil {
		slog.Info("Error getting all sites", "error", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	slog.Info("Sites fetched from db", "allSites", sites)

	res := map[string]any{
		"sites": sites,
	}

	return c.JSON(http.StatusOK, res)
}

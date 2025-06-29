package server

import (
	"io"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Server) testHttpConnection(c echo.Context) error {
	var config struct {
		Url     string            `json:"url"`
		Headers map[string]string `json:"headers"`
	}
	if err := c.Bind(&config); err != nil {
		slog.Error("Failed to bind http config", "error", err)

		return echo.NewHTTPError(400, err)
	}

	req, err := http.NewRequest(http.MethodGet, config.Url, nil)
	if err != nil {
		slog.Error("Failed to create new request", "error", err)
		return echo.NewHTTPError(400, err)
	}

	for k, v := range config.Headers {
		req.Header.Add(k, v)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		slog.Error("Failed to send request", "error", err)
		return echo.NewHTTPError(400, err)
	}
	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		slog.Error("Failed to read request body", "error", err)
		return echo.NewHTTPError(400, err)
	}

	return c.JSON(201, map[string]any{
		"status": res.Status,
		"body":   string(b),
	})
}

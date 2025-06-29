package server

import (
	"api/internal/store/site"
	"io"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Server) testHttpConnection(c echo.Context) error {
	var config site.HttpSiteConfig
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

func (s *Server) createHttpSite(c echo.Context) error {
	var httpSiteData site.NewHttpSite
	if err := c.Bind(&httpSiteData); err != nil {
		slog.Error("Failed to bind http site data", "error", err)
		return echo.NewHTTPError(400, err)
	}

	slog.Info("Received request", "data", httpSiteData)

	site, err := s.store.Sites.CreateHttp(c.Request().Context(), httpSiteData)
	if err != nil {
		slog.Error("Failed to create http site", "error", err)
		return echo.NewHTTPError(500, err)
	}

	res := map[string]any{
		"site": site,
	}

	return c.JSON(201, res)
}

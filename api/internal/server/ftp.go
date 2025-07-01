package server

import (
	"api/internal/store/site"
	"fmt"
	"log/slog"
	"time"

	"github.com/jlaffaye/ftp"
	"github.com/labstack/echo/v4"
)

func (s *Server) testFtpConnection(c echo.Context) error {
	var config site.FtpSiteConfig
	if err := c.Bind(&config); err != nil {
		slog.Error("Failed to bind ftp config", "error", err)
		return echo.NewHTTPError(400, err)
	}

	connStr := fmt.Sprintf("%s:%d", config.Url, config.Port)
	conn, err := ftp.Dial(
		connStr,
		ftp.DialWithContext(c.Request().Context()),
		ftp.DialWithTimeout(5*time.Second),
	)
	if err != nil {
		slog.Error("Failed to dial ftp server", "connectionStr", connStr, "error", err)
		return echo.NewHTTPError(400, err)
	}

	if config.AuthRequired {
		if err = conn.Login(config.Username, config.Password); err != nil {
			slog.Error("Failed to login on ftp server", "connectionStr", connStr, "username", config.Username, "error", err)
			return echo.NewHTTPError(400, err)
		}
	}

	// Maybe we can do something to test the connection more
	entries, err := conn.List(".")
	if err != nil {
		slog.Error("Failed to issue LIST command to ftp server", "connectionStr", connStr, "error", err)
		return echo.NewHTTPError(400, err)
	}

	if err = conn.Quit(); err != nil {
		slog.Error("Failed to quit connection to ftp server", "connectionStr", connStr, "error", err)
		return echo.NewHTTPError(500, err)
	}

	return c.JSON(201, map[string]any{
		"status":  "success",
		"entries": entries,
	})
}

func (s *Server) createFtpSite(c echo.Context) error {
	var ftpSiteData site.NewFtpSite
	if err := c.Bind(&ftpSiteData); err != nil {
		slog.Error("Failed to bind ftp site data", "error", err)
		return echo.NewHTTPError(400, err)
	}

	slog.Info("Received request", "data", ftpSiteData)

	site, err := s.store.Sites.CreateFtp(c.Request().Context(), ftpSiteData)
	if err != nil {
		slog.Error("Failed to create ftp site", "error", err)
		return echo.NewHTTPError(500, err)
	}

	res := map[string]any{
		"site": site,
	}

	return c.JSON(201, res)
}

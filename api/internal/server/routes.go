package server

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"https://*", "http://*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	httpGroup := e.Group("/http")

	httpGroup.POST("/connect", s.testHttpConnection)

	e.POST("/", s.IndexHandler)

	return e
}

func (s *Server) IndexHandler(c echo.Context) error {
	var data struct {
		Type         string `json:"type"`
		Address      string `json:"address"`
		Port         int64  `json:"port"`
		AuthRequired bool   `json:"authenticationRequired"`
		Username     string `json:"username"`
		Password     string `json:"password"`
		File         struct {
			Path string `json:"path"`
			Type string `json:"type"`
		} `json:"fileInfo"`
	}
	if err := c.Bind(&data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	slog.Info("Received data", data)

	resp := map[string]string{
		"message": "Hello World",
	}

	return c.JSON(http.StatusOK, resp)
}

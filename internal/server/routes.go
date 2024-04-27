package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", s.Health)
	e.POST("/idea", s.Idea)

	return e
}

func (s *Server) Health(c echo.Context) error {
	resp := map[string]string{
		"message": "Server is UP and running and ready to kick ass",
	}

	return c.JSON(http.StatusOK, resp)
}

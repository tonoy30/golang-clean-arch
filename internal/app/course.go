package app

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (a app) HealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "healthy")
}

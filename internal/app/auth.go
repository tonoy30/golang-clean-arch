package app

import (
	"github.com/labstack/echo/v4"
	"github.com/tonoy30/clean-arch/internal/dto"
	"net/http"
)

func (a app) SignUp(c echo.Context) error {
	u := &dto.User{}
	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	id, err := a.authService.SignUp(u)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, id)
}

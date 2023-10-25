package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Register(c echo.Context)error {
	return c.JSON(http.StatusOK, nil)
}
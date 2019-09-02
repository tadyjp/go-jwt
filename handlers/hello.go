package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Hello hoge
func Hello(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"hello": "world"})
}

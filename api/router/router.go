package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Hello return hello message
func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, IMBIZ!")
}

package router

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// Hello return hello message
func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, IMBIZ!")
}

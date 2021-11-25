package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// @Summary Get Hello
// @Description Get Hello Test
// @Accept json
// @Produce json
// @Success 200 {string} string "Hello, IMBIZ!"
// @Router / [get]
func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, IMBIZ!")
}

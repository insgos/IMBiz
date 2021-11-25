package auth

import (
	"github.com/labstack/echo/v4"
)

// @Summary Read User
// @Description Read User
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {string} string "success"
// @Router /userinfo/read [get]
func ReadUser(c echo.Context) error {
	return nil
}

// @Summary  Create User
// @Description Create User
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {string} string "success"
// @Router /userinfo/create [post]
func CreateUser(c echo.Context) error {
	return nil
}

// @Summary Update User
// @Description Update User
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {string} string "success"
// @Router /userinfo/update [put]
func UpdateUser(c echo.Context) error {
	return nil
}

// @Summary Delete User
// @Description Delete User
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {string} string "success"
// @Router /userinfo/delete [delete]
func DeleteUser(c echo.Context) error {
	return nil
}

// @Summary Login
// @Description Login
// @Tags Login
// @Accept json
// @Produce json
// @Success 200 {string} string "success"
// @Router /userinfo/login [POST]
func Login(c echo.Context) error {
	return nil
}

// @Summary Logout
// @Description Logout
// @Tags Login
// @Accept json
// @Produce json
// @Success 200 {string} string "success"
// @Router /userinfo/logout [POST]
func Logout(c echo.Context) error {
	return nil
}

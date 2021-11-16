package main

import (
	"imbiz/api/auth"
	"imbiz/api/router"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", router.Hello)

	// USER INFO CRUD
	e.POST("/userinfo/create", auth.CreateUser)
	e.POST("/userinfo/read", auth.ReadUser)
	e.POST("/userinfo/update", auth.UpdateUser)
	e.POST("/userinfo/delete", auth.DeleteUser)

	e.POST("/userinfo/login", auth.Login)
	e.POST("/userinfo/logout", auth.Logout)

	e.Logger.Fatal(e.Start(":18181"))
}

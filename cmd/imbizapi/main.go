package main

import (
	"imbiz/api/auth"
	"imbiz/api/router"
	_ "imbiz/cmd/imbizapi/docs"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	swagger "github.com/swaggo/echo-swagger"
)

// @title IMBiz Swagger API
// @version 1.0
// @host localhost:18181
// @BasePath /
func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	// SWAGGER URL
	e.GET("/swagger/*", swagger.WrapHandler)
	// TEST URL
	e.GET("/", router.Hello)
	// USER INFO CRUD

	e.GET("/userinfo/read", auth.ReadUser)
	e.POST("/userinfo/create", auth.CreateUser)
	e.PUT("/userinfo/update", auth.UpdateUser)
	e.DELETE("/userinfo/delete", auth.DeleteUser)

	e.POST("/userinfo/login", auth.Login)
	e.POST("/userinfo/logout", auth.Logout)

	e.Logger.Fatal(e.Start(":18181"))
}

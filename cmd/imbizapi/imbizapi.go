package main

import (
	"github.com/labstack/echo/v4"
	"imbiz/api/router"
)

func main() {
	e := echo.New()

	e.GET("/", router.Hello)

	e.Logger.Fatal(e.Start(":18181"))
}

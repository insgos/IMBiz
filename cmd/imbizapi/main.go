package main

import (
	"imbiz/api/auth"
	"imbiz/api/router"
	_ "imbiz/cmd/imbizapi/docs"
	"log"
	"net/http"
	"os"
	"strconv"

	cnf "imbiz/api/config"
	logger "imbiz/logger"

	mysql "imbiz/db/mysql"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	swagger "github.com/swaggo/echo-swagger"
)

type db struct {
	mysql *mysql.MYDB
}

// @title IMBiz Swagger API
// @version 1.0
// @host localhost:18181
// @BasePath /
func main() {
	conFile := "api/config/default.json"
	env := os.Getenv("GO_ENV")

	if env == "stage" {
		conFile = "api/config/stage.json"
	}

	// config setting
	_conf, _ := cnf.LoadConfiguration(conFile)
	logger.SetLogger(_conf.Logger.Path)

	// db init
	mysql.Connect(_conf)
	e := echo.New()
	/*
		echo Middleware setting
		- CORS Middleware
		- Recover (에러가 발생해도 서버가 바로 죽지 않고, 다시 살아 남)
		- Log print
		- context
	*/
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return next(c)
		}
	})

	// SWAGGER URL
	e.GET("/swagger/*", swagger.WrapHandler)

	// health check url
	e.GET("/", router.Hello)
	e.GET("/mysql/read", router.MysqlRead)
	e.POST("/mysql/create", router.MysqlCreate)
	e.PUT("/mysql/update", router.MysqlUpdate)
	e.DELETE("/mysql/delete", router.MysqlDelete)

	// USER INFO CRUD
	e.GET("/userinfo/read", auth.ReadUser)
	e.POST("/userinfo/create", auth.CreateUser)
	e.PUT("/userinfo/update", auth.UpdateUser)
	e.DELETE("/userinfo/delete", auth.DeleteUser)

	e.POST("/userinfo/login", auth.Login)
	e.POST("/userinfo/logout", auth.Logout)

	log.Println("application started on http:0.0.0.0://" + strconv.Itoa(_conf.System.Port))
	e.Logger.Fatal(e.Start(":" + strconv.Itoa(_conf.System.Port)))
}

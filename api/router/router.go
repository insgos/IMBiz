package router

import (
	"log"
	"net/http"

	mysql "imbiz/db/mysql"

	"github.com/labstack/echo/v4"
)

type Test struct {
	INDEX int
	TEXT  string
}

// @Summary Get Hello
// @Description Get Hello Test
// @Accept json
// @Produce json
// @Success 200 {string} string "Hello, IMBIZ!"
// @Router / [get]
func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, IMBIZ!")
}

// @Summary GET Mysql Create
// @Description GET Mysql Create
// @Accept json
// @Produce json
// @Success 200 {array} []map[string]string{}
// @Router /mysql/read [get]
func MysqlRead(c echo.Context) error {
	db := mysql.GetDBInstance("imbiz")
	log.Println(db)
	d, err := db.Get("SELECT * FROM TEST")
	if err != nil {
		log.Panic(err)
	}
	return c.JSONPretty(http.StatusOK, d, "  ")
}

// @Summary POST Mysql Create
// @Description POST Mysql Create
// @Accept json
// @Produce json
// @Param body body object true "TEXT"
// @Success 200 {string} string "success"
// @Router /mysql/create [post]
func MysqlCreate(c echo.Context) error {
	db := mysql.GetDBInstance("imbiz")
	params := make(map[string]string)
	_ = c.Bind(&params)

	err := db.Set("INSERT INTO `TEST` (`TEXT`) VALUES ('" + params["TEXT"] + "')")
	if err != nil {
		log.Panic(err)
	}
	return c.String(http.StatusOK, "success")
}

// @Summary PUT Mysql Update
// @Description PUT Mysql Update
// @Accept json
// @Produce json
// @Param body body object true "INDEX, TEXT"
// @Success 200 {string} string "success"
// @Router /mysql/update [put]
func MysqlUpdate(c echo.Context) error {
	db := mysql.GetDBInstance("imbiz")
	params := make(map[string]string)
	_ = c.Bind(&params)

	err := db.Set("UPDATE `TEST` SET `TEXT` = '" + params["TEXT"] + "' WHERE (`INDEX` = " + params["INDEX"] + ")")
	if err != nil {
		log.Panic(err)
	}
	return c.String(http.StatusOK, "success")
}

// @Summary DELETE Mysql Delete
// @Description DELETE Mysql Delete
// @Accept json
// @Produce json
// @Param body body object true "INDEX"
// @Success 200 {string} string "success"
// @Router /mysql/delete [delete]
func MysqlDelete(c echo.Context) error {
	db := mysql.GetDBInstance("imbiz")
	params := make(map[string]string)
	_ = c.Bind(&params)

	err := db.Set("DELETE FROM `TEST` WHERE (`INDEX` = " + params["INDEX"] + ")")
	if err != nil {
		log.Panic(err)
	}
	return c.String(http.StatusOK, "success")
}

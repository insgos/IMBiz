package imdb

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func NewDb() *sql.DB {
	db, e := sql.Open("mysql", "root:root@tcp(127.0.0.1:3336)/IMBiz?parseTime=true")
	if e != nil {
		log.Println(e)
		panic(e)
	}
	return db
}

func CloseDb(db *sql.DB) {
	db.Close()
}

package mysql

import (
	"database/sql"
	"log"
	"strconv"
	"strings"

	cnf "imbiz/api/config"

	_ "github.com/go-sql-driver/mysql"
)

type (
	CONFIG struct {
		Addr     string
		Username string
		Password string
		Database string
	}
)

type MYDB struct {
	*sql.DB
}

var DB *MYDB

func (t *CONFIG) GetUrl() string {
	tmp := []string{}
	if len(t.Username) > 0 && len(t.Password) > 0 {
		tmp = append(tmp, t.Username+":"+t.Password+"@")
	}
	tmp = append(tmp, "tcp("+t.Addr+")")
	if len(t.Database) > 0 {
		tmp = append(tmp, "/"+t.Database)
	}
	tmp = append(tmp, "?&multiStatements=true&charset=utf8&autocommit=true")
	return strings.Join(tmp, "")
}

func Connect(_conf cnf.Config) {
	log.Println(_conf)
	connstr := (&CONFIG{
		Addr:     _conf.Mysql.Host + ":" + strconv.Itoa(_conf.Mysql.Port),
		Username: _conf.Mysql.Usr,
		Password: _conf.Mysql.Pwd,
		Database: _conf.Mysql.Db,
	}).GetUrl()

	instance, err := sql.Open("mysql", connstr)
	if err != nil {
		log.Panic(err)
	}
	if err = instance.Ping(); err != nil {
		log.Panic(err)
	}

	// instance.SetConnMaxLifetime(0)
	instance.SetMaxIdleConns(10)
	instance.SetMaxOpenConns(100)
	DB = (&MYDB{DB: instance})
}

func GetDBInstance() *MYDB {
	return DB
}

func (t *MYDB) Get(query string) ([]map[string]interface{}, error) {
	rows, err := t.Query(query)
	if err != nil {
		return nil, err
	}
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	return tableData, nil
}

func (t *MYDB) Set(query string) error {
	_, err := t.Exec(query)
	return err
}

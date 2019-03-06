package database

import (
	"database/sql"
	"log"

	// mysql驱动
	_ "github.com/go-sql-driver/mysql"
)

var SqlDB *sql.DB

func init() {
	var err error
	//本地数据库
	// SqlDB, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/blog01?parseTime=true")
	//测试数据库
	SqlDB, err = sql.Open("mysql", "wallace:wallace@gpxj115566@tcp(debug.gupiaoxianji.com:3309)/gpxj_app?parseTime=true")
	if err != nil {
		log.Fatal(err.Error())
	}

	err = SqlDB.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}

	SqlDB.SetMaxIdleConns(4)
	SqlDB.SetMaxOpenConns(10)
}

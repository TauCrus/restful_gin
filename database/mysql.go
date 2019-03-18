package database

import (
	sql "database/sql"
	"log"
	"restful_gin/config"

	// mysql驱动
	_ "github.com/go-sql-driver/mysql"
)

// SQLDB 数据库连接
var SQLDB *sql.DB

func init() {
	if config.Apns {
		SQLDB = NewMysql(config.DSN, 4, 10)
	} else {
		SQLDB = NewMysql(config.TestDSN, 4, 10)
	}
}

// NewMysql 新建数据库连接
func NewMysql(dataSourceName string, maxIdleConns, maxOpenConns int) *sql.DB {
	log.Println("Data Source Name:", dataSourceName)

	db, err := sql.Open("mysql", dataSourceName)
	if nil != err {
		log.Println(err)
	}

	err = db.Ping()
	if nil != err {
		log.Println(err)
	}

	db.SetMaxIdleConns(maxIdleConns)
	db.SetMaxOpenConns(maxOpenConns)

	return db
}

package main

import (
	db "restful_gin/database"
)

func main() {
	defer db.SQLDB.Close()
	router := initRouter()
	router.Run(":9420")
}

package main

import(
	db "restful_gin/database"
)

func main() {
	defer db.SqlDB.Close()
	router := initRouter()
	router.Run(":8080")
}


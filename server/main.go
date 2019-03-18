package main

import (
	"flag"
	"fmt"
	"restful_gin/config"

	"github.com/golang/glog"
)

func main() {
	flag.Parse()
	defer glog.Flush()

	// defer db.SQLDB.Close()
	router := initRouter()
	ListenOnPort := fmt.Sprintf(":%d", config.Port)
	router.Run(ListenOnPort)
}

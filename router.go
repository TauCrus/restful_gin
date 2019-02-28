package main

import (
	"github.com/gin-gonic/gin"
	. "restful_gin/apis"
)

func initRouter() *gin.Engine{
	router := gin.Default()

	router.GET("/", IndexApi)

	router.POST("/person", AddPersonApi)

	router.GET("/person", GetPersonsApi)

	router.GET("/person/:id", GetPersonApi)

	router.PUT("/person/:id", UpdatePersonApi)
	
	router.DELETE("/person/:id", DeletePersonApi)

	return router
}
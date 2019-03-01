package main

import (
	. "restful_gin/apis"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	router := gin.Default()
	// 处理跨域问题
	router.Use(cors.Default())
	/**
	* 测试用例
	 */
	router.GET("/", IndexAPI)

	router.POST("/person", AddPersonAPI)

	router.GET("/person", GetPersonsAPI)

	router.GET("/person/:id", GetPersonAPI)

	router.PUT("/person/:id", UpdatePersonAPI)

	router.DELETE("/person/:id", DeletePersonAPI)

	/**
	*用户接口
	 */
	router.POST("/api/user/login", UserLoginAPI)

	router.POST("/api/user/logout", UserLogoutAPI)

	router.POST("/api/user/register", UserRegisterAPI)

	router.GET("/api/user/info", UserInfoAPI)

	router.POST("/api/user/changePassword", UserChangePwdAPI)

	/**
	*系统用户接口
	 */
	router.GET("/api/sys/user/query", SysUserQueryAPI)

	/**
	*系统设置接口
	 */
	router.GET("/api/sys/review/query", SysReviewsQueryAPI)

	/**
	* 订单
	 */
	router.GET("/api/count/order", GetOrdersAPI)

	return router
}

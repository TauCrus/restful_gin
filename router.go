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
	{
		router.GET("/", IndexAPI)

		router.POST("/person", AddPersonAPI)

		router.GET("/person", GetPersonsAPI)

		router.GET("/person/:id", GetPersonAPI)

		router.PUT("/person/:id", UpdatePersonAPI)

		router.DELETE("/person/:id", DeletePersonAPI)

	}

	/**
	*用户接口
	 */
	user := router.Group("api/user")
	{
		user.POST("/login", UserLoginAPI)

		user.POST("/logout", UserLogoutAPI)

		user.POST("/register", UserRegisterAPI)

		user.GET("/info", UserInfoAPI)

		user.POST("/changePassword", UserChangePwdAPI)
	}
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
	router.GET("/api/count/order", QueryOrdersAPI)

	// 服务路由
	{
		router.POST("/api/user/service", AddServiceAPI)
		router.GET("/api/user/service", GetServicesAPI)
		// router.GET("/api/user/service/:id", GetServiceAPI)
		router.PUT("/api/user/service", ModifyServiceAPI)
		router.DELETE("/api/user/service", DropServiceAPI)
	}

	//产品路由
	{
		router.GET("/api/product/product", GetProductsAPI)
	}

	return router
}

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
	*管理用户接口
	 */
	router.GET("/api/sys/user/query", GetGetAdminUsersAPI)

	/**
	*系统设置接口
	 */
	router.GET("/api/sys/review/query", SysReviewsQueryAPI)

	/**
	* 订单
	 */
	router.GET("/api/count/order", QueryOrdersAPI)

	{
		// 用户管理路由
		router.GET("/api/user/manage", GetUsersAPI)

		// 用户服务路由
		router.POST("/api/user/service", AddServiceAPI)
		router.GET("/api/user/service", GetServicesAPI)
		router.PUT("/api/user/service", ModifyServiceAPI)
		router.DELETE("/api/user/service", DropServiceAPI)

		// 用户反馈路由
		router.GET("/api/user/feedback", GetFeedbacksAPI)
		router.DELETE("/api/user/feedback", DropFeedbackAPI)
	}

	//产品路由
	{
		router.GET("/api/product/product", GetProductsAPI)
		router.GET("/api/product/productlist", GetProductListAPI)
		router.GET("/api/product/price", GetProductPricesAPI)
		router.GET("/api/product/recommend", GetProductRecommendsAPI)
	}

	// 优惠券路由
	{
		router.GET("/api/coupon/manage", GetCouponsAPI)
		router.GET("/api/coupon/record", GetCouponRecordAPI)
		router.GET("/api/coupon/activity", GetCouponActivityAPI)
	}

	// 内容路由
	{
		// 文案组
		cw := router.Group("api/content/copywriter")
		cw.GET("/banner", GetBannersAPI)
		cw.GET("/startpage", GetStartPagesAPI)
		cw.GET("/searchrecommend", GetSearchRecommendsAPI)
		cw.GET("/marketinglabel", GetMarketingLabelsAPI)

		// 首页组
		hp := router.Group("api/content/homepage")
		hp.GET("/column", GetColumnsAPI)
		hp.GET("/productcolumn", GetProductColumnsAPI)
		hp.GET("/spcolumn", GetSPColumnsAPI)
		hp.GET("/articlecolumn", GetArticleColumnsAPI)
		hp.GET("/adcolumn", GetAdColumnsAPI)
		hp.GET("/shortcutmenu", GetShortCutMenusAPI)
		hp.GET("/productclassify", GetProductClassifysAPI)
		hp.GET("/suspendad", GetSuspendAdsAPI)
		hp.GET("/popup", GetPopupsAPI)

	}

	return router
}

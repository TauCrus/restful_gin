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

	{
		/**
		*统计
		 */
		// 查询订单
		router.GET("/api/count/order", QueryOrdersAPI)
		//用户注册
		router.GET("/api/count/userregist", QueryUserRegistAPI)
	}

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
		cw.POST("/banner", AddBannerAPI)
		cw.PUT("/banner", ModifyBannerAPI)
		cw.DELETE("/banner", DropBannerAPI)

		cw.GET("/bannerplace", GetBannerPlacesAPI)

		cw.GET("/startpage", GetStartPagesAPI)
		cw.POST("/startpage", AddStartPageAPI)
		cw.PUT("/startpage", ModifyStartPageAPI)
		cw.DELETE("/startpage", DropStartPageAPI)

		cw.GET("/searchrecommend", GetSearchRecommendsAPI)
		cw.POST("/searchrecommend", AddSearchRecommendAPI)
		cw.PUT("/searchrecommend", ModifySearchRecommendAPI)
		cw.DELETE("/searchrecommend", DropSearchRecommendAPI)

		cw.GET("/marketinglabel", GetMarketingLabelsAPI)
		cw.POST("/marketinglabel", AddMarketingLabelAPI)
		cw.PUT("/marketinglabel", ModifyMarketingLabelAPI)
		cw.DELETE("/marketinglabel", DropMarketingLabelAPI)

		// 首页组
		hp := router.Group("api/content/homepage")
		hp.GET("/column", GetColumnsAPI)
		hp.POST("/column", AddColumnAPI)
		hp.PUT("/column", ModifyColumnAPI)
		hp.DELETE("/column", DropColumnAPI)

		hp.GET("/columntype", GetColumnTypesAPI)

		hp.GET("/productcolumn", GetProductColumnsAPI)
		hp.POST("/productcolumn", AddProductColumnAPI)
		hp.PUT("/productcolumn", ModifyProductColumnAPI)
		hp.DELETE("/productcolumn", DropProductColumnAPI)

		hp.GET("/spcolumn", GetSPColumnsAPI)
		hp.POST("/spcolumn", AddSPColumnAPI)
		hp.PUT("/spcolumn", ModifySPColumnAPI)
		hp.DELETE("/spcolumn", DropSPColumnAPI)

		hp.GET("/articlecolumn", GetArticleColumnsAPI)
		hp.POST("/articlecolumn", AddArticleColumnAPI)
		hp.PUT("/articlecolumn", ModifyArticleColumnAPI)
		hp.DELETE("/articlecolumn", DropArticleColumnAPI)

		hp.GET("/adcolumn", GetAdColumnsAPI)
		hp.POST("/adcolumn", AddAdColumnAPI)
		hp.PUT("/adcolumn", ModifyAdColumnAPI)
		hp.DELETE("/adcolumn", DropAdColumnAPI)

		hp.GET("/shortcutmenu", GetShortCutMenusAPI)
		hp.POST("/shortcutmenu", AddShortCutMenuAPI)
		hp.PUT("/shortcutmenu", ModifyShortCutMenuAPI)
		hp.DELETE("/shortcutmenu", DropShortCutMenuAPI)

		hp.GET("/productclassify", GetProductClassifysAPI)
		hp.POST("/productclassify", AddProductClassifyAPI)
		hp.PUT("/productclassify", ModifyProductClassifyAPI)
		hp.DELETE("/productclassify", DropProductClassifyAPI)

		hp.GET("/suspendad", GetSuspendAdsAPI)
		hp.POST("/suspendad", AddActivityMarketingAPI)
		hp.PUT("/suspendad", ModifyActivityMarketingAPI)
		hp.DELETE("/suspendad", DropActivityMarketingAPI)

		hp.GET("/popup", GetPopupsAPI)
		hp.POST("/popup", AddActivityMarketingAPI)
		hp.PUT("/popup", ModifyActivityMarketingAPI)
		hp.DELETE("/popup", DropActivityMarketingAPI)

	}

	return router
}

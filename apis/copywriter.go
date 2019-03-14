package apis

import (
	"log"
	"net/http"
	"restful_gin/models"

	"github.com/gin-gonic/gin"
)

// GetBannersResult 轮播图结果
type GetBannersResult struct {
	Data []models.Banner `json:"data"`
}

// GetBannersAPI 查询轮播图接口
func GetBannersAPI(c *gin.Context) {

	title := c.Request.FormValue("title")
	activityName := c.Request.FormValue("activity_name")
	status := c.Request.FormValue("status")

	cw := models.Copywriter{}

	banners, err := cw.GetBanners(0, title, activityName, status)
	if nil != err {
		log.Fatalln(err)
	}

	result := GetBannersResult{Data: banners}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"result":  result,
	})
}

// GetStartPagesResult 启动页结果
type GetStartPagesResult struct {
	Data []models.StartPage `json:"data"`
}

// GetStartPagesAPI 查询启动页接口
func GetStartPagesAPI(c *gin.Context) {
	cw := models.Copywriter{}

	startpages, err := cw.GetStartPages()
	if nil != err {
		log.Fatalln(err)
	}

	result := GetStartPagesResult{Data: startpages}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"result":  result,
	})
}

// GetSearchRecommendsResult 搜索推荐结果
type GetSearchRecommendsResult struct {
	Data []models.SearchRecommend `json:"data"`
}

// GetSearchRecommendsAPI 查询搜索推荐接口
func GetSearchRecommendsAPI(c *gin.Context) {
	cw := models.Copywriter{}

	srs, err := cw.GetSearchRecommends()
	if nil != err {
		log.Fatalln(err)
	}

	result := GetSearchRecommendsResult{Data: srs}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"result":  result,
	})
}

// GetMarketingLabelsResult 搜索推荐结果
type GetMarketingLabelsResult struct {
	Data []models.MarketingLabel `json:"data"`
}

// GetMarketingLabelsAPI 查询搜索推荐接口
func GetMarketingLabelsAPI(c *gin.Context) {
	cw := models.Copywriter{}

	labels, err := cw.GetMarketingLabels()
	if nil != err {
		log.Fatalln(err)
	}

	result := GetMarketingLabelsResult{Data: labels}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"result":  result,
	})
}

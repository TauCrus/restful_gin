package apis

import (
	"fmt"
	"log"
	"net/http"
	"restful_gin/models"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
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

// GetBannerPlaceResult 轮播图位置结果
type GetBannerPlaceResult struct {
	Data []models.BannerPlace `json:"data"`
}

// GetBannerPlacesAPI 查询轮播图位置结果
func GetBannerPlacesAPI(c *gin.Context) {
	cw := models.Copywriter{}

	bps, err := cw.GetBannerPlaces()
	if nil != err {
		glog.Error(err)
	}

	result := GetBannerPlaceResult{Data: bps}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"result":  result,
	})
}

// AddBannerAPI 新增轮播图接口
func AddBannerAPI(c *gin.Context) {
	banner := models.Banner{}
	err := c.Bind(&banner)
	if nil != err {
		glog.Error(err)
	}
	raRows, err := new(models.Copywriter).AddBanner(banner)
	if nil != err {
		glog.Error(err)
	}

	if raRows > 0 {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"result":  gin.H{"id": raRows},
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"result":  "error add",
		})
	}
}

// ModifyBannerAPI 修改轮播图接口
func ModifyBannerAPI(c *gin.Context) {
	banner := models.Banner{}
	err := c.Bind(&banner)
	if nil != err {
		glog.Error(err)
	}

	raRows, err := new(models.Copywriter).ModifyBanner(banner)
	if nil != err {
		glog.Error(err)
	}

	msg := fmt.Sprintf("Update Banner %d successful %d", banner.ID, raRows)
	if raRows > 0 {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"result":  msg,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"result":  "error modify",
		})
	}
}

// DropBannerAPI 删除轮播图接口
func DropBannerAPI(c *gin.Context) {
	banner := models.Banner{}
	err := c.Bind(&banner)
	if nil != err {
		glog.Error(err)
	}
	raRows, err := new(models.Copywriter).DropBanner(banner)
	if raRows > 0 {
		msg := fmt.Sprintf("Delete Banner %d successful %d", banner.ID, raRows)
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"result":  msg,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"result":  "error delete",
		})
	}
}

// GetStartPagesResult 启动页结果
type GetStartPagesResult struct {
	Data []models.StartPage `json:"data"`
}

// GetStartPagesAPI 查询启动页接口
func GetStartPagesAPI(c *gin.Context) {
	title := c.Request.FormValue("title")
	status := c.Request.FormValue("status")

	cw := models.Copywriter{}
	startpages, err := cw.GetStartPages(title, status)
	if nil != err {
		log.Fatalln(err)
	}

	result := GetStartPagesResult{Data: startpages}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"result":  result,
	})
}

// AddStartPageAPI 新增启动页接口
func AddStartPageAPI(c *gin.Context) {
	startpage := models.StartPage{}
	err := c.Bind(&startpage)
	if nil != err {
		glog.Error(err)
	}
	raRows, err := new(models.Copywriter).AddStartPage(startpage)
	if nil != err {
		glog.Error(err)
	}

	if raRows > 0 {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"result":  gin.H{"id": raRows},
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"result":  "error add",
		})
	}
}

// ModifyStartPageAPI 修改启动页接口
func ModifyStartPageAPI(c *gin.Context) {
	startpage := models.StartPage{}
	err := c.Bind(&startpage)
	if nil != err {
		glog.Error(err)
	}
	raRows, err := new(models.Copywriter).ModifyStartPage(startpage)
	if nil != err {
		glog.Error(err)
	}
	msg := fmt.Sprintf("Update StartPage %d successful %d", startpage.ID, raRows)
	if raRows > 0 {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"result":  msg,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"result":  "error modify",
		})
	}
}

// DropStartPageAPI 删除启动页接口
func DropStartPageAPI(c *gin.Context) {
	startpage := models.StartPage{}
	err := c.Bind(&startpage)
	if nil != err {
		glog.Error(err)
	}
	raRows, err := new(models.Copywriter).DropStartPage(startpage)
	if raRows > 0 {
		msg := fmt.Sprintf("Delete StartPage %d successful %d", startpage.ID, raRows)
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"result":  msg,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"result":  "error delete",
		})
	}
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

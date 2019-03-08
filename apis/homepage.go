package apis

import (
	"log"
	"net/http"
	"restful_gin/models"

	"github.com/gin-gonic/gin"
)

// GetColumnsResult 首页栏目结果
type GetColumnsResult struct {
	Data []models.Column `json:"data"`
}

// GetColumnsAPI 查询首页栏目接口
func GetColumnsAPI(c *gin.Context) {
	hp := models.Homepage{}

	columns, err := hp.GetColumns()
	if nil != err {
		log.Fatalln(err)
	}

	result := GetColumnsResult{Data: columns}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"result":  result,
	})
}

// GetProductColumnsResult 首页推荐产品栏目结果
type GetProductColumnsResult struct {
	Data []models.ProductColumn `json:"data"`
}

// GetProductColumnsAPI 查询首页推荐产品栏目接口
func GetProductColumnsAPI(c *gin.Context) {
	hp := models.Homepage{}

	pcs, err := hp.GetProductColumns()
	if nil != err {
		log.Fatalln(err)
	}

	result := GetProductColumnsResult{Data: pcs}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"result":  result,
	})
}

// GetSPColumnsResult 首页荐股产品栏目结果
type GetSPColumnsResult struct {
	Data []models.SPColumn `json:"data"`
}

// GetSPColumnsAPI 查询首页荐股产品栏目接口
func GetSPColumnsAPI(c *gin.Context) {
	hp := models.Homepage{}
	spcs, err := hp.GetSPColumns()
	if nil != err {
		log.Fatalln(err)
	}

	result := GetSPColumnsResult{Data: spcs}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"result":  result,
	})
}

// GetArticleColumnsResult 首页文章栏目结果
type GetArticleColumnsResult struct {
	Data []models.ArticleColumn `json:"data"`
}

// GetArticleColumnsAPI 查询首页文章栏目接口
func GetArticleColumnsAPI(c *gin.Context) {
	hp := models.Homepage{}
	acs, err := hp.GetArticleColumns()
	if nil != err {
		log.Fatalln(err)
	}

	result := GetArticleColumnsResult{Data: acs}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"result":  result,
	})
}

// GetAdColumnsAPI 查询首页广告栏目接口
func GetAdColumnsAPI(c *gin.Context) {
	cw := models.Copywriter{}
	banners, err := cw.GetBanners(1)
	if nil != err {
		log.Fatalln(err)
	}

	result := GetBannersResult{Data: banners}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"result":  result,
	})
}

// GetShortCutMenusResult 首页快捷菜单结果
type GetShortCutMenusResult struct {
	Data []models.ShortCutMenu `json:"data"`
}

// GetShortCutMenusAPI 查询首页快捷菜单接口
func GetShortCutMenusAPI(c *gin.Context) {
	hp := models.Homepage{}
	scms, err := hp.GetShortCutMenus()
	if nil != err {
		log.Fatalln(err)
	}

	result := GetShortCutMenusResult{Data: scms}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"result":  result,
	})
}

// GetProductClassifyResult 首页产品分类结果
type GetProductClassifyResult struct {
	Data []models.ProductClassify `json:"data"`
}

// GetProductClassifysAPI 查询首页产品分类接口
func GetProductClassifysAPI(c *gin.Context) {
	hp := models.Homepage{}
	pcs, err := hp.GetProductClassifys()
	if nil != err {
		log.Fatalln(err)
	}

	result := GetProductClassifyResult{Data: pcs}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"result":  result,
	})
}

// GetActivityMarketingResult 活动营销结果
type GetActivityMarketingResult struct {
	Data []models.ActivityMarketing `json:"data"`
}

// GetSuspendAdAPI 查询悬浮广告接口
func GetSuspendAdsAPI(c *gin.Context) {
	hp := models.Homepage{}
	ams, err := hp.GetActivityMarketings(1)
	if nil != err {
		log.Fatalln(err)
	}

	result := GetActivityMarketingResult{Data: ams}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"result":  result,
	})
}

// GetPopupsAPI 查询营销弹窗接口
func GetPopupsAPI(c *gin.Context) {
	hp := models.Homepage{}
	ams, err := hp.GetActivityMarketings(2)
	if nil != err {
		log.Fatalln(err)
	}

	result := GetActivityMarketingResult{Data: ams}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"result":  result,
	})
}

package apis

import (
	"fmt"
	"log"
	"net/http"
	"restful_gin/models"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
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

// AddColumnAPI 新增栏目接口
func AddColumnAPI(c *gin.Context) {
	column := models.Column{}
	err := c.Bind(&column)
	if nil != err {
		glog.Error(err)
	}

	raRows, err := new(models.Homepage).AddColumn(column)
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

// ModifyColumnAPI 修改栏目接口
func ModifyColumnAPI(c *gin.Context) {
	column := models.Column{}
	err := c.Bind(&column)
	if nil != err {
		glog.Error(err)
	}
	raRows, err := new(models.Homepage).ModifyColumn(column)
	if nil != err {
		glog.Error(err)
	}

	msg := fmt.Sprintf("Update Column %d successful %d", column.ID, raRows)
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

// DropColumnAPI 删除栏目接口
func DropColumnAPI(c *gin.Context) {
	column := models.Column{}
	err := c.Bind(&column)
	if nil != err {
		glog.Error(err)
	}
	raRows, err := new(models.Homepage).DropColumn(column)
	if nil != err {
		glog.Error(err)
	}

	if raRows > 0 {
		msg := fmt.Sprintf("Delete Column %d successful %d", column.ID, raRows)
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

// AddProductColumnAPI 新增推荐产品栏目接口
func AddProductColumnAPI(c *gin.Context) {
	pc := models.ProductColumn{}
	err := c.Bind(&pc)
	if nil != err {
		glog.Error(err)
	}

	raRows, err := new(models.Homepage).AddProductColumn(pc)
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

// ModifyProductColumnAPI 修改推荐产品栏目接口
func ModifyProductColumnAPI(c *gin.Context) {
	pc := models.ProductColumn{}
	err := c.Bind(&pc)
	if nil != err {
		glog.Error(err)
	}
	raRows, err := new(models.Homepage).ModifyProductColumn(pc)
	if nil != err {
		glog.Error(err)
	}

	msg := fmt.Sprintf("Update ProductColumn %d successful %d", pc.ID, raRows)
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

// DropProductColumnAPI 删除推荐产品栏目接口
func DropProductColumnAPI(c *gin.Context) {
	pc := models.ProductColumn{}
	err := c.Bind(&pc)
	if nil != err {
		glog.Error(err)
	}
	raRows, err := new(models.Homepage).DropProductColumn(pc)
	if nil != err {
		glog.Error(err)
	}

	if raRows > 0 {
		msg := fmt.Sprintf("Delete ProductColumn %d successful %d", pc.ID, raRows)
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

// AddSPColumnAPI 新增荐股产品栏目接口
func AddSPColumnAPI(c *gin.Context) {
	spc := models.SPColumn{}
	err := c.Bind(&spc)
	if nil != err {
		glog.Error(err)
	}

	raRows, err := new(models.Homepage).AddSPColumn(spc)
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

// ModifySPColumnAPI 修改荐股产品栏目接口
func ModifySPColumnAPI(c *gin.Context) {
	spc := models.SPColumn{}
	err := c.Bind(&spc)
	if nil != err {
		glog.Error(err)
	}
	raRows, err := new(models.Homepage).ModifySPColumn(spc)
	if nil != err {
		glog.Error(err)
	}

	msg := fmt.Sprintf("Update SPColumn %d successful %d", spc.ID, raRows)
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

// DropSPColumnAPI 删除荐股产品栏目接口
func DropSPColumnAPI(c *gin.Context) {
	spc := models.SPColumn{}
	err := c.Bind(&spc)
	if nil != err {
		glog.Error(err)
	}
	raRows, err := new(models.Homepage).DropSPColumn(spc)
	if nil != err {
		glog.Error(err)
	}

	if raRows > 0 {
		msg := fmt.Sprintf("Delete SPColumn %d successful %d", spc.ID, raRows)
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

// AddArticleColumnAPI 新增文章栏目接口
func AddArticleColumnAPI(c *gin.Context) {
	ac := models.ArticleColumn{}
	err := c.Bind(&ac)
	if nil != err {
		glog.Error(err)
	}

	raRows, err := new(models.Homepage).AddArticleColumn(ac)
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

// ModifyArticleColumnAPI 修改文章栏目接口
func ModifyArticleColumnAPI(c *gin.Context) {
	ac := models.ArticleColumn{}
	err := c.Bind(&ac)
	if nil != err {
		glog.Error(err)
	}
	raRows, err := new(models.Homepage).ModifyArticleColumn(ac)
	if nil != err {
		glog.Error(err)
	}

	msg := fmt.Sprintf("Update ArticleColumn %d successful %d", ac.ID, raRows)
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

// DropArticleColumnAPI 删除文章栏目接口
func DropArticleColumnAPI(c *gin.Context) {
	ac := models.ArticleColumn{}
	err := c.Bind(&ac)
	if nil != err {
		glog.Error(err)
	}
	raRows, err := new(models.Homepage).DropArticleColumn(ac)
	if nil != err {
		glog.Error(err)
	}

	if raRows > 0 {
		msg := fmt.Sprintf("Delete ArticleColumn %d successful %d", ac.ID, raRows)
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

// GetAdColumnsAPI 查询首页广告栏目接口
func GetAdColumnsAPI(c *gin.Context) {
	title := c.Request.FormValue("title")
	activityName := c.Request.FormValue("activity_name")
	status := c.Request.FormValue("status")

	cw := models.Copywriter{}
	banners, err := cw.GetBanners(1, title, activityName, status)
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

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
	columnTypeID := c.Request.FormValue("column_type_id")
	columnTypeIDs := c.Request.FormValue("column_type_ids")
	columnName := c.Request.FormValue("column_name")
	status := c.Request.FormValue("status")

	hp := models.Homepage{}
	columns, err := hp.GetColumns(columnTypeID, columnTypeIDs, columnName, status)
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

// GetColumnTypesResult 首页栏目类型结果
type GetColumnTypesResult struct {
	Data []models.ColumnType `json:"data"`
}

// GetColumnTypesAPI 查询首页栏目类型接口
func GetColumnTypesAPI(c *gin.Context) {

	hp := models.Homepage{}
	cts, err := hp.GetColumnTypes()
	if nil != err {
		log.Fatalln(err)
	}

	result := GetColumnTypesResult{Data: cts}

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
	keyword := c.Request.FormValue("keyword")
	columnID := c.Request.FormValue("column_id")
	status := c.Request.FormValue("status")

	hp := models.Homepage{}
	pcs, err := hp.GetProductColumns(keyword, columnID, status)
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
	keyword := c.Request.FormValue("keyword")
	columnID := c.Request.FormValue("column_id")
	status := c.Request.FormValue("status")

	hp := models.Homepage{}
	spcs, err := hp.GetSPColumns(keyword, columnID, status)
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
	keyword := c.Request.FormValue("keyword")
	placeID := c.Request.FormValue("place_id")
	columnID := c.Request.FormValue("column_id")
	status := c.Request.FormValue("status")

	cw := models.Copywriter{}
	banners, err := cw.GetBanners(1, keyword, placeID, columnID, status)
	if nil != err {
		log.Fatalln(err)
	}

	result := GetBannersResult{Data: banners}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"result":  result,
	})
}

// AddAdColumnAPI 新增首页广告栏目接口
func AddAdColumnAPI(c *gin.Context) {
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

// ModifyAdColumnAPI 修改首页广告栏目接口
func ModifyAdColumnAPI(c *gin.Context) {
	banner := models.Banner{}
	err := c.Bind(&banner)
	if nil != err {
		glog.Error(err)
	}

	raRows, err := new(models.Copywriter).ModifyBanner(banner)
	if nil != err {
		glog.Error(err)
	}

	msg := fmt.Sprintf("Update AdColumn %d successful %d", banner.ID, raRows)
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

// DropAdColumnAPI 删除首页广告栏目接口
func DropAdColumnAPI(c *gin.Context) {
	banner := models.Banner{}
	err := c.Bind(&banner)
	if nil != err {
		glog.Error(err)
	}
	raRows, err := new(models.Copywriter).DropBanner(banner)
	if nil != err {
		glog.Error(err)
	}

	if raRows > 0 {
		msg := fmt.Sprintf("Delete AdColumn %d successful %d", banner.ID, raRows)
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

// AddShortCutMenuAPI 新增首页快捷菜单接口
func AddShortCutMenuAPI(c *gin.Context) {
	scm := models.ShortCutMenu{}
	err := c.Bind(&scm)
	if nil != err {
		glog.Error(err)
	}

	raRows, err := new(models.Homepage).AddShortCutMenu(scm)
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

// ModifyShortCutMenuAPI 修改首页快捷菜单接口
func ModifyShortCutMenuAPI(c *gin.Context) {
	scm := models.ShortCutMenu{}
	err := c.Bind(&scm)
	if nil != err {
		glog.Error(err)
	}
	raRows, err := new(models.Homepage).ModifyShortCutMenu(scm)
	if nil != err {
		glog.Error(err)
	}

	msg := fmt.Sprintf("Update ShortCutMenu %d successful %d", scm.ID, raRows)
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

// DropShortCutMenuAPI 删除首页快捷菜单接口
func DropShortCutMenuAPI(c *gin.Context) {
	scm := models.ShortCutMenu{}
	err := c.Bind(&scm)
	if nil != err {
		glog.Error(err)
	}
	raRows, err := new(models.Homepage).DropShortCutMenu(scm)
	if nil != err {
		glog.Error(err)
	}

	if raRows > 0 {
		msg := fmt.Sprintf("Delete ShortCutMenu %d successful %d", scm.ID, raRows)
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

// AddProductClassifyAPI 新增首页产品分类接口
func AddProductClassifyAPI(c *gin.Context) {
	pc := models.ProductClassify{}
	err := c.Bind(&pc)
	if nil != err {
		glog.Error(err)
	}

	raRows, err := new(models.Homepage).AddProductClassify(pc)
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

// ModifyProductClassifyAPI 修改首页产品分类接口
func ModifyProductClassifyAPI(c *gin.Context) {
	scm := models.ProductClassify{}
	err := c.Bind(&scm)
	if nil != err {
		glog.Error(err)
	}
	raRows, err := new(models.Homepage).ModifyProductClassify(scm)
	if nil != err {
		glog.Error(err)
	}

	msg := fmt.Sprintf("Update ProductClassify %d successful %d", scm.ID, raRows)
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

// DropProductClassifyAPI 删除首页产品分类接口
func DropProductClassifyAPI(c *gin.Context) {
	pc := models.ProductClassify{}
	err := c.Bind(&pc)
	if nil != err {
		glog.Error(err)
	}
	raRows, err := new(models.Homepage).DropProductClassify(pc)
	if nil != err {
		glog.Error(err)
	}

	if raRows > 0 {
		msg := fmt.Sprintf("Delete ProductClassify %d successful %d", pc.ID, raRows)
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

// GetActivityMarketingResult 活动营销结果
type GetActivityMarketingResult struct {
	Data []models.ActivityMarketing `json:"data"`
}

// GetSuspendAdsAPI 查询悬浮广告接口
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

// AddActivityMarketingAPI 新增营销活动接口
func AddActivityMarketingAPI(c *gin.Context) {
	am := models.ActivityMarketing{}
	err := c.Bind(&am)
	if nil != err {
		glog.Error(err)
	}

	raRows, err := new(models.Homepage).AddActivityMarketing(am)
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

// ModifyActivityMarketingAPI 修改营销活动接口
func ModifyActivityMarketingAPI(c *gin.Context) {
	am := models.ActivityMarketing{}
	err := c.Bind(&am)
	if nil != err {
		glog.Error(err)
	}
	raRows, err := new(models.Homepage).ModifyActivityMarketing(am)
	if nil != err {
		glog.Error(err)
	}

	msg := fmt.Sprintf("Update ActivityMarketing %d successful %d", am.ID, raRows)
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

// DropActivityMarketingAPI 删除营销活动接口
func DropActivityMarketingAPI(c *gin.Context) {
	am := models.ActivityMarketing{}
	err := c.Bind(&am)
	if nil != err {
		glog.Error(err)
	}
	raRows, err := new(models.Homepage).DropActivityMarketing(am)
	if nil != err {
		glog.Error(err)
	}

	if raRows > 0 {
		msg := fmt.Sprintf("Delete ActivityMarketing %d successful %d", am.ID, raRows)
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

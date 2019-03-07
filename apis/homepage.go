package apis

import (
	"log"
	"net/http"
	"restful_gin/models"

	"github.com/gin-gonic/gin"
)

// GetColumnsResult 栏目结果
type GetColumnsResult struct {
	Data []models.Column `json:"data"`
}

// GetColumnsAPI 查询栏目接口
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

// GetProductColumnsResult 栏目结果
type GetProductColumnsResult struct {
	Data []models.ProductColumn `json:"data"`
}

// GetProductColumnsAPI 查询栏目接口
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

package apis

import (
	"log"
	"net/http"
	"restful_gin/models"

	"github.com/gin-gonic/gin"
)

// GetProductsResult 产品查询结果
type GetProductsResult struct {
	Data []models.Product `json:"data"`
}

// GetProductsAPI 获取产品
func GetProductsAPI(c *gin.Context) {
	products, err := new(models.Product).GetProducts()
	if nil != err {
		log.Fatalln(err)
	}

	result := GetProductsResult{Data: products}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"result":  result,
	})
}

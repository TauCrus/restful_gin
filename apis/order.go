package apis

import (
	"net/http"
	"restful_gin/models"

	"github.com/gin-gonic/gin"
)

// OrderResult 订单结果
type OrderResult struct {
	Data []models.Order `json:"data"`
}

//GetOrdersAPI 获取订单
func GetOrdersAPI(c *gin.Context) {
	order := models.Order{}
	orders, err := order.GetOrders()
	if nil != err {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   err,
		})
	}

	result := OrderResult{Data: orders}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"result":  result,
	})
}

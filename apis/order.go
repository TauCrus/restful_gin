package apis

import (
	"log"
	"net/http"
	"restful_gin/models"

	"github.com/gin-gonic/gin"
)

// OrderResult 订单结果
type OrderResult struct {
	Data []models.Order `json:"data"`
}

//QueryOrdersAPI 获取订单
func QueryOrdersAPI(c *gin.Context) {
	order := models.Order{}
	orders, err := order.QueryOrders()
	if nil != err {
		log.Fatalln(err)
	}

	result := OrderResult{Data: orders}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"result":  result,
	})

}

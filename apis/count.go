package apis

import (
	"log"
	"net/http"
	"restful_gin/models"

	"github.com/gin-gonic/gin"
)

// QueryUserRegistResult 用户注册统计结果
type QueryUserRegistResult struct {
	Data []models.UserRegist `json:"data"`
}

// QueryUserRegistAPI 查询用户注册统计接口
func QueryUserRegistAPI(c *gin.Context) {

	startDate := c.Request.FormValue("start_date")
	endDate := c.Request.FormValue("end_date")
	group := c.Request.FormValue("group")

	count := models.Count{}
	urs, err := count.QueryUserRegist(startDate, endDate, group)
	if nil != err {
		log.Fatalln(err)
	}

	result := QueryUserRegistResult{Data: urs}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"result":  result,
	})
}

// OrderResult 订单结果
type OrderResult struct {
	Data []models.Order `json:"data"`
}

//QueryOrdersAPI 获取订单
func QueryOrdersAPI(c *gin.Context) {
	startDate := c.Request.FormValue("start_date")
	endDate := c.Request.FormValue("end_date")
	orderID := c.Request.FormValue("order_id")
	payWay := c.Request.FormValue("pay_way")
	userid := c.Request.FormValue("userid")
	productClass := c.Request.FormValue("product_class")

	count := models.Count{}
	orders, err := count.QueryOrders(startDate, endDate, orderID, payWay, userid, productClass)
	if nil != err {
		log.Fatalln(err)
	}

	result := OrderResult{Data: orders}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"result":  result,
	})

}

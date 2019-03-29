package apis

import (
	"log"
	"net/http"
	"restful_gin/models"

	"github.com/gin-gonic/gin"
)

// GetCouponsResult 查询优惠券结果
type GetCouponsResult struct {
	Data []models.CouponData `json:"data"`
}

// GetCouponsAPI 查询优惠券接口
func GetCouponsAPI(c *gin.Context) {

	keyword := c.Request.FormValue("keyword")
	couponID := c.Request.FormValue("coupon_id")
	couponType := c.Request.FormValue("type")
	status := c.Request.FormValue("status")

	cdList, err := new(models.Coupon).GetCoupons(keyword, couponID, couponType, status)
	if nil != err {
		log.Fatalln(err)
	}

	result := GetCouponsResult{Data: cdList}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"result":  result,
	})
}

// GetCouponRecordResult 优惠券领取记录结果
type GetCouponRecordResult struct {
	Data []models.CouponRecord `json:"data"`
}

// GetCouponRecordAPI 查询优惠券领取记录接口
func GetCouponRecordAPI(c *gin.Context) {
	coList, err := new(models.Coupon).GetCouponRecord()
	if nil != err {
		log.Fatalln(err)
	}
	result := GetCouponRecordResult{Data: coList}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"result":  result,
	})
}

// GetCouponActivityResult 优惠券活动
type GetCouponActivityResult struct {
	Data []models.CouponActivity `json:"data"`
}

// GetCouponActivityAPI 查询优惠券领取记录接口
func GetCouponActivityAPI(c *gin.Context) {
	caList, err := new(models.Coupon).GetCouponActivity()
	if nil != err {
		log.Fatalln(err)
	}
	result := GetCouponActivityResult{Data: caList}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"result":  result,
	})
}

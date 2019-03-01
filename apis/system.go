package apis

import (
	"net/http"
	"restful_gin/models"

	"github.com/gin-gonic/gin"
)

// ReviewResult 审核结果
type ReviewResult struct {
	Data []models.Review `json:"data"`
}

// SysReviewsQueryAPI 获取审核管理数据 接口
func SysReviewsQueryAPI(c *gin.Context) {
	reviews, err := new(models.System).SysReviewsQuery()
	if nil != err {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   err,
		})
	}

	result := ReviewResult{Data: reviews}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"result":  result,
	})
}

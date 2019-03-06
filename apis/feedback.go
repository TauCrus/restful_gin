package apis

import (
	"fmt"
	"log"
	"net/http"
	"restful_gin/models"

	"github.com/gin-gonic/gin"
)

// GetFeedbacksResult 反馈查询结果
type GetFeedbacksResult struct {
	Data []models.Feedback `json:"data"`
}

// GetFeedbacksAPI 查询服务接口
func GetFeedbacksAPI(c *gin.Context) {
	phone := c.Request.FormValue("phone")
	email := c.Request.FormValue("email")

	f := models.Feedback{}

	feedbacks, err := f.GetFeedbacks(phone, email)
	if nil != err {
		log.Fatalln(err)
	}

	result := GetFeedbacksResult{Data: feedbacks}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"result":  result,
	})
}

// DropFeedbackAPI 删除服务接口
func DropFeedbackAPI(c *gin.Context) {

	f := models.Feedback{}
	err := c.Bind(&f)
	if nil != err {
		log.Fatalln(err)
	}

	raRows, err := f.DropFeedback()

	msg := fmt.Sprintf("Delete Feedback %d successful %d", f.ID, raRows)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"result":  msg,
	})
}

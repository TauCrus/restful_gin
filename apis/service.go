package apis

import (
	"fmt"
	"log"
	"net/http"
	"restful_gin/models"

	"github.com/gin-gonic/gin"
)

// AddServiceAPI 新增服务接口
func AddServiceAPI(c *gin.Context) {

	s := models.Service{}
	err := c.Bind(&s)
	if nil != err {
		log.Fatalln(err)
	}
	raRows, err := s.AddService()
	if nil != err {
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"result":  gin.H{"id": raRows},
	})
}

// GetServicesResult 服务查询结果
type GetServicesResult struct {
	Data []models.Service `json:"data"`
}

// GetServicesAPI 查询服务接口
func GetServicesAPI(c *gin.Context) {
	userid := c.Request.FormValue("userid")
	productClass := c.Request.FormValue("product_class")

	service := models.Service{}

	services, err := service.GetServices(userid, productClass)
	if nil != err {
		log.Fatalln(err)
	}

	result := GetServicesResult{Data: services}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"result":  result,
	})
}

// ModifyServiceAPI 修改服务接口
func ModifyServiceAPI(c *gin.Context) {
	service := models.Service{}
	err := c.Bind(&service)
	if nil != err {
		log.Fatalln(err)
	}
	log.Println("service:", service)

	raRows, err := service.ModifyService()
	msg := fmt.Sprintf("Update Service %d successful %d", service.ID, raRows)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"result":  msg,
	})
}

// DropServiceAPI 删除服务接口
func DropServiceAPI(c *gin.Context) {

	s := models.Service{}
	err := c.Bind(&s)
	if nil != err {
		log.Fatalln(err)
	}

	raRows, err := s.DropService()

	msg := fmt.Sprintf("Delete Service %d successful %d", s.ID, raRows)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"result":  msg,
	})
}

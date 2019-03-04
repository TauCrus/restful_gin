package apis

import (
	"log"
	"net/http"
	"restful_gin/models"

	"github.com/gin-gonic/gin"
)

// AddServiceAPI 新增服务接口
func AddServiceAPI(c *gin.Context) {

}

// GetServicesResult 服务查询结果
type GetServicesResult struct {
	Data []models.Service `json:"data"`
}

// GetServicesAPI 新增服务接口
func GetServicesAPI(c *gin.Context) {
	services, err := new(models.Service).GetServices()
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

}

// DropServiceAPI 删除服务接口
func DropServiceAPI(c *gin.Context) {

}

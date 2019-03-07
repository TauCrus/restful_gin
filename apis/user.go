package apis

import (
	"log"
	"net/http"
	"restful_gin/models"

	"github.com/gin-gonic/gin"
)

// Permission 权限
type Permission struct {
	Path string `json:"path"`
}

// UserInfo 用户信息
type UserInfo struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Roles       []string     `json:"roles"`
	Permissions []Permission `json:"permissions"`
}

// Error 错误信息
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// UserInfoAPI 用户信息接口
func UserInfoAPI(c *gin.Context) {
	var (
		roles, permissions []string
	)

	roles = []string{"admin"}
	permissions = []string{"/index", "/table", "/forms/base", "/forms/edit",
		"/user/password", "/about", "/sys/user/list", "/sys/review/list",
		"/count/order/list", "/count/regist/list",
		"/user/manage/list", "/user/service/list", "/user/feedback/list",
		"/coupon/manage/list", "/coupon/activity/list", "/coupon/record/list",
		"/product/manage/list", "/product/price/list", "/product/package/list", "/product/recommend/list",
		"/content/copywriter",
		"/content/copywriter/banner/list",
		"/content/copywriter/startpage/list",
		"/content/copywriter/searchrecommend/list",
		"/content/copywriter/marketinglabel/list",
		"/content/homepage",
		"/content/homepage/column/list"}

	var userinfo = UserInfo{ID: "12138", Name: "spring", Roles: roles}

	for _, v := range permissions {
		var p Permission
		p.Path = v
		userinfo.Permissions = append(userinfo.Permissions, p)
	}

	err := Error{Code: 100000, Message: "无效的token"}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"result":  userinfo,
		"error":   err,
	})
}

// UserLoginResult 用户登录结果
type UserLoginResult struct {
	Token string `json:"token"`
}

// UserLoginAPI 用户登录接口
func UserLoginAPI(c *gin.Context) {
	result := UserLoginResult{Token: "12138"}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"result":  result,
	})
}

//UserLogoutResult 用户注销结果
type UserLogoutResult struct {
	Token string `json:"token"`
}

// UserLogoutAPI 用户注销接口
func UserLogoutAPI(c *gin.Context) {
	result := UserLogoutResult{Token: "12138"}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"result":  result,
	})
}

// UserRegisterResult 用户注册结果
type UserRegisterResult struct {
	Token string `json:"token"`
}

// UserRegisterAPI 用户注销接口
func UserRegisterAPI(c *gin.Context) {
	result := UserRegisterResult{Token: "12138"}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"result":  result,
	})
}

// UserChangePwdResult 用户修改密码结果
type UserChangePwdResult struct {
	Token string `json:"token"`
}

// UserChangePwdAPI 用户修改密码接口
func UserChangePwdAPI(c *gin.Context) {
	result := UserChangePwdResult{Token: "12138"}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"result":  result,
	})
}

// AdminUserResult 用户查询结果
type AdminUserResult struct {
	Data []models.AdminUser `json:"data"`
}

// GetGetAdminUsersAPI 后台用户查询接口
func GetGetAdminUsersAPI(c *gin.Context) {
	u := models.User{}
	users, err := u.GetAdminUsers()
	if nil != err {
		log.Fatalln(err)
	}

	result := AdminUserResult{Data: users}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"result":  result,
	})
}

// GetUserResult 用户查询结果
type GetUserResult struct {
	Data []models.User `json:"data"`
}

// GetUsersAPI 用户查询接口
func GetUsersAPI(c *gin.Context) {
	u := models.User{}
	users, err := u.GetUsers()
	if nil != err {
		log.Fatalln(err)
	}

	result := GetUserResult{Data: users}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"result":  result,
	})
}

package apis

import (
	"encoding/json"
	"log"
	"net/http"
	"restful_gin/models"
	"restful_gin/utils"
	"time"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
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

// GetUserInfoAPI 用户信息接口
func GetUserInfoAPI(c *gin.Context) {

	username := getUsernameFromToken(c)
	glog.Info("username:", username)
	roles, permissions, err := new(models.User).GetUserInfo(username)
	if nil != err {
		glog.Info(err)
	}
	/*
		var (
			roles, permissions []string
		)

		roles = []string{"admin"}
		permissions = []string{
			"/index", "/table", "/forms/base", "/forms/edit",
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
			"/content/homepage/column/list",
			"/content/homepage/productcolumn/list",
			"/content/homepage/spcolumn/list",
			"/content/homepage/articlecolumn/list",
			"/content/homepage/adcolumn/list",
			"/content/homepage/shortcutmenu/list",
			"/content/homepage/productclassify/list",
			"/content/homepage/suspendad/list",
			"/content/homepage/popup/list"}

	*/
	var userinfo = UserInfo{ID: "12138", Name: username, Roles: roles}
	if len(permissions) == 0 {
		permissions = []string{"/index", "/table", "/forms/base", "/forms/edit", "/user/password", "/about"}
	}
	for _, v := range permissions {
		var p Permission
		p.Path = v
		userinfo.Permissions = append(userinfo.Permissions, p)
	}

	// err := Error{Code: 100000, Message: "无效的token"}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"result":  userinfo,
		// "error":   err,
	})
}

// 从token种获取用户名
func getUsernameFromToken(c *gin.Context) string {
	claims, _ := c.Get("claims")
	// glog.Info("claims:", claims)

	tmp, err := json.Marshal(claims)
	if nil != err {
		glog.Error(err)
	}

	cc := &utils.CustomClaims{}
	err = json.Unmarshal(tmp, &cc)
	if nil != err {
		glog.Error(err)
	}

	return cc.Username
}

// UserLoginParam 用户登录参数
type UserLoginParam struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// UserLoginResult 用户登录结果
type UserLoginResult struct {
	Token string `json:"token"`
}

// UserLoginAPI 用户登录接口
func UserLoginAPI(c *gin.Context) {

	param := UserLoginParam{}
	err := c.Bind(&param)
	if nil != err {
		glog.Error(err)
	}
	glog.Info("param:", param)

	cnt, err := new(models.User).UserLogin(param.Username, param.Password)

	if cnt > 0 {
		generateToken(c, param)

	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"result":  "用户名或者密码错误",
			"error":   utils.Error{Code: 10001, Message: "登录失败"},
		})
	}

}

// generateToken 生成令牌 token
func generateToken(c *gin.Context, param UserLoginParam) {
	j := &utils.JWT{
		[]byte("gpztxy"),
	}
	claims := utils.CustomClaims{
		param.Username,
		jwtgo.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000),  // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 60*60), // 过期时间 一小时
			Issuer:    "gpztxy",                         //签名的发行者
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"error":   utils.Error{Code: -1, Message: err.Error()},
		})
		return
	}

	glog.Info("token:", token)

	result := UserLoginResult{Token: token}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"result":  result,
	})

	return
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

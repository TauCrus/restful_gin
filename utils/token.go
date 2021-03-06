package utils

import (
	"errors"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

// JWTAuth 中间件 检查 token
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if "" == token {
			c.JSON(http.StatusOK, gin.H{
				"success": false,
				"error":   Error{Code: -1, Message: "请求未携带token，无权限访问"},
			})
			c.Abort()
			return
		}

		glog.Info("get token:", token)

		j := NewJWT()

		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if nil != err {
			if ErrTokenExpired == err {
				c.JSON(http.StatusOK, gin.H{
					"success": false,
					"error":   Error{Code: 100000, Message: "授权已过期"},
				})
				c.Abort()
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"success": false,
				"error":   Error{Code: -1, Message: err.Error()},
			})
			c.Abort()
			return
		}
		// 继续交由下一个路由处理,并将解析出的信息传递下去
		c.Set("claims", claims)
	}
}

// JWT 签名结构
type JWT struct {
	SigningKey []byte
}

//一些常量
var (
	ErrTokenExpired     error  = errors.New("Token is expired")
	ErrTokenNotValidYet error  = errors.New("Token not active yet")
	ErrTokenMalformed   error  = errors.New("That's not even a token")
	ErrTokenInvalid     error  = errors.New("Couldn't handle this token:")
	SignKey             string = "gpztxy"
)

// CustomClaims 载荷，可以加一些自己需要的信息
type CustomClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// NewJWT 新建一个jwt实例
func NewJWT() *JWT {
	return &JWT{
		[]byte(GetSignKey()),
	}
}

// GetSignKey 获取signKey
func GetSignKey() string {
	return SignKey
}

// SetSignKey 设置SignKey
func SetSignKey(key string) string {
	SignKey = key
	return SignKey
}

// CreateToken 生成一个token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// ParseToken 解析Tokne
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, ErrTokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, ErrTokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, ErrTokenNotValidYet
			} else {
				return nil, ErrTokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, ErrTokenInvalid
}

// RefreshToken 更新token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}

	return "", ErrTokenInvalid
}



package services

import (
	"deploy/config"
	"deploy/models"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// IdentityKey 认证绑定Key; models.User.ID
var IdentityKey = "user_id"

type login struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

// AuthMiddleware Jwt授权认证中间件结构体
var AuthMiddleware, err = jwt.New(&jwt.GinJWTMiddleware{
	Realm:           "Authentication",
	Key:             []byte(config.App.AppKey), // 秘钥（从配置文件获取）
	Timeout:         time.Duration(config.App.Jwt.Timeout) * time.Minute,
	MaxRefresh:      time.Duration(config.App.Jwt.Refresh) * time.Minute,
	IdentityKey:     IdentityKey,
	IdentityHandler: IdentityHandler,
	Authenticator:   Authenticator,
	PayloadFunc:     PayloadFunc,
	Authorizator:    Authorizator,
	Unauthorized:    Unauthorized,
	TokenLookup:     "header: Authorization",
	TokenHeadName:   "Bearer",
	TimeFunc:        time.Now,
})

// Authenticator 认证
var Authenticator = func(c *gin.Context) (interface{}, error) {
	var loginParams login
	if err := c.ShouldBind(&loginParams); nil != err {
		return "", jwt.ErrMissingLoginValues
	}

	username := loginParams.Username
	password := loginParams.Password

	user, err := models.GetUserByUsername(username)
	if nil != err {
		return nil, err
	}

	// 密码正确，那么将用户对象返回，到时候记录到payload中
	if user.IsPasswordRight(password) {
		return user, nil
	}

	return nil, jwt.ErrFailedAuthentication
}

// PayloadFunc 将会在Login的时候调用
var PayloadFunc = func(data interface{}) jwt.MapClaims {
	if v, ok := data.(*models.User); ok {
		return jwt.MapClaims{
			IdentityKey: v.ID,
		}
	}

	return jwt.MapClaims{}
}

// IdentityHandler 解析认证结果放到IdentityKey 上
var IdentityHandler = func(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	userid := claims[IdentityKey].(string)
	user, _ := models.GetUserByID(userid)

	return user
}

// Authorizator 授权，登录成功后，判断用户是否有权限访问
var Authorizator = func(user interface{}, c *gin.Context) bool {
	return true
}

// Unauthorized 授权失败;没有访问权限的情况调用该函数
var Unauthorized = func(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}

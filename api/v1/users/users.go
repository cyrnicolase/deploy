package users

import (
	"deploy/models"
	"deploy/services"
	"fmt"

	"github.com/gin-gonic/gin"
)

// GetUsers fetch all users
func GetUsers(c *gin.Context) {
	users, err := models.GetAllUsers()
	if nil != users {
		c.JSON(200, users)
		return
	}

	c.String(500, "查询用户列表出错: "+err.Error())
}

// PostUsers add a new user
func PostUsers(c *gin.Context) {
	var user models.User

	if err := c.ShouldBind(&user); nil != err {
		c.String(400, fmt.Sprintf("Bad request: %v", err))
		return
	}

	affected, err := user.Insert()
	if nil != err {
		c.String(500, fmt.Sprintf("新增用户操作失败: %v", err))
		return
	}

	c.String(200, fmt.Sprintf("新增用户操作成功，影响行数: %d", affected))
}

// ModifyPassword 修改登录用户密码
func ModifyPassword(c *gin.Context) {
	user, _ := c.Get(services.IdentityKey)
	passwd := c.PostForm("password")

	_, err := user.(*models.User).ResetPassword(passwd)
	if nil != err {
		c.String(500, fmt.Sprintf("更新用户密码错误:%v", err))
		return
	}

	c.String(200, "更新用户密码成功")
	return
}

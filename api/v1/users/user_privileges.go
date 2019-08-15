package users

import (
	"deploy/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

// GetUserPrivileges 查看用户所拥有的权限
func GetUserPrivileges(c *gin.Context) {
	userID := c.Query("user_id")
	userPrivileges := models.UserPrivilegesByUserID(userID)
	result := make([]map[string]string, 0)
	for _, v := range userPrivileges {
		row := make(map[string]string)
		row["id"] = v.UserPrivilege.ID
		row["username"] = v.Username
		row["user_id"] = v.User.ID
		row["privilege"] = v.Privilege

		result = append(result, row)
	}

	c.JSON(200, result)
}

// PostUserPrivileges 新增/修改用户权限
func PostUserPrivileges(c *gin.Context) {
	var userPrivilege models.UserPrivilege
	if err := c.ShouldBind(&userPrivilege); nil != err {
		c.String(400, "参数格式错误:"+err.Error())
		return
	}

	var affect int64
	has, err := userPrivilege.UnscopedGet()
	if !has {
		affect, err = userPrivilege.Insert()
	} else {
		result, _ := userPrivilege.Restore()
		affect, _ = result.RowsAffected()
	}

	if nil != err {
		c.String(500, "新增用户权限操作失败:"+err.Error())
		return
	}

	c.String(200, fmt.Sprintf("新增用户权限操作成功,新增记录数:%d", affect))
	return
}

// DeleteUserPrivilege 删除用户权限
func DeleteUserPrivilege(c *gin.Context) {
	id := c.Param("id")

	affected, err := models.DeleteByID(id)
	if nil != err {
		c.String(500, "删除用户权限操作失败:"+err.Error())
		return
	}

	c.String(200, fmt.Sprintf("删除用户权限操作成功,删除记录数:%d", affected))
}

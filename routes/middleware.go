package routes

import (
	"deploy/models"
	"deploy/services"

	"github.com/gin-gonic/gin"
)

func mwAuthPrivilege(privilege string) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, ok := c.Get(services.IdentityKey)
		if !ok {
			c.AbortWithStatus(403)

			return
		}

		if user.(*models.User).IsSuper {
			return // 如果是超级管理员，那么不用进行权限验证
		}

		userPrivilegs := models.UserPrivilegesByUserID(user.(*models.User).ID)
		for _, up := range userPrivilegs {
			if privilege == up.Privilege {
				return // 如果权限正确，可以访问
			}
		}

		c.AbortWithStatus(403)

		c.Next()
	}
}

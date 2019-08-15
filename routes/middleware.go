package routes

import (
	"deploy/models"
	"deploy/services"

	"github.com/gin-gonic/gin"
)

func mwAuthPrivilege(privilege string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userid, ok := c.Get(services.IdentityKey)
		if !ok {
			c.AbortWithStatus(403)

			return
		}

		userPrivilegs := models.UserPrivilegesByUserID(userid.(string))
		for _, up := range userPrivilegs {
			if privilege == up.Privilege {
				return
			}
		}

		c.AbortWithStatus(403)

		c.Next()
	}
}

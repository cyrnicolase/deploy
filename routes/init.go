package routes

import (
	"github.com/gin-gonic/gin"

	"deploy/api/v1/meta"
	"deploy/api/v1/users"
)

// Boot 是启动
func Boot() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		metaRouter := v1.Group("/meta")
		{
			metaRouter.GET("/ping", meta.GetPing)
		}

		usersRouter := v1.Group("/users")
		{
			usersRouter.GET("/users", users.GetUsers)
			usersRouter.POST("/users", users.PostUsers)

			usersRouter.POST("/privileges", users.PostUserPrivileges)
			usersRouter.GET("/privileges", users.GetUserPrivileges)
			usersRouter.DELETE("/privileges/:id", users.DeleteUserPrivilege)
		}
	}

	return router
}

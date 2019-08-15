package routes

import (
	"github.com/gin-gonic/gin"

	"deploy/api/v1/meta"
	"deploy/api/v1/users"
	"deploy/services"
)

// Boot 是启动
func Boot() *gin.Engine {
	router := gin.Default()

	// 认证
	router.POST("/api/login", services.AuthMiddleware.LoginHandler)

	v1 := router.Group("/api/v1")
	v1.Use(services.AuthMiddleware.MiddlewareFunc())
	{

		metaRouter := v1.Group("/meta")
		{
			metaRouter.GET("/ping", meta.GetPing)
		}

		usersRouter := v1.Group("/users")
		{
			usersRouter.POST("/password", users.ModifyPassword)

			usersRouter.GET("/users", mwAuthPrivilege("查看用户"), users.GetUsers)
			usersRouter.POST("/users", mwAuthPrivilege("新增用户"), users.PostUsers)

			// usersRouter.POST("/privileges", users.PostUserPrivileges)
			usersRouter.POST("/privileges", mwAuthPrivilege("关联用户权限"), users.PostUserPrivileges)
			usersRouter.GET("/privileges", mwAuthPrivilege("查看用户权限"), users.GetUserPrivileges)
			usersRouter.DELETE("/privileges/:id", mwAuthPrivilege("删除用户权限"), users.DeleteUserPrivilege)
		}
	}

	return router
}

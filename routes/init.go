package routes

import (
	"github.com/gin-gonic/gin"

	"deploy/api/v1/meta"
	"deploy/api/v1/project"
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

			usersRouter.GET("/privileges-list", users.GetAllPrivileges)
			usersRouter.POST("/privileges", mwAuthPrivilege("关联用户权限"), users.PostUserPrivileges)
			usersRouter.GET("/privileges", mwAuthPrivilege("查看用户权限"), users.GetUserPrivileges)
			usersRouter.DELETE("/privileges/:id", mwAuthPrivilege("删除用户权限"), users.DeleteUserPrivilege)
		}

		projectRouter := v1.Group("/project")
		{
			projectRouter.GET("/projects", mwAuthPrivilege("查看项目列表"), project.Projects)
			projectRouter.POST("/projects", mwAuthPrivilege("新增项目"), project.PostProject)
			projectRouter.PUT("/projects/:id", mwAuthPrivilege("修改项目"), project.PutProject)
			projectRouter.DELETE("/projects/:id", mwAuthPrivilege("删除项目"), project.DeleteProject)

			projectRouter.POST("/servers", mwAuthPrivilege("新增服务器"), project.PostServer)
			projectRouter.PUT("/servers/:id", mwAuthPrivilege("修改服务器"), project.PutServer)
		}
	}

	return router
}

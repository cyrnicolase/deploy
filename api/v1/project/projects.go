package project

import (
	"strconv"

	"deploy/models"

	"github.com/gin-gonic/gin"
)

// Projects return list of project
func Projects(c *gin.Context) {
	projects, err := models.GetProjectList()
	if nil != err {
		c.String(500, "获取项目列表失败:"+err.Error())
		return
	}

	c.JSON(200, projects)
	return
}

// PostProject build a new project
func PostProject(c *gin.Context) {
	var project models.Project
	if err := c.ShouldBindJSON(&project); nil != err {
		c.String(400, "新增项目参数格式错误:"+err.Error())
		return
	}

	affected, err := project.Create()
	if nil != err {
		c.String(500, "新增项目失败:"+err.Error())
		return
	}

	c.String(200, "新增项目成功:"+strconv.Itoa(int(affected)))
	return
}

// PutProject modify a project
func PutProject(c *gin.Context) {
	var project models.Project
	if err := c.ShouldBind(&project); nil != err {
		c.String(400, "修改项目失败:"+err.Error())
		return
	}

	project.ID = c.Param("id")
	affected, err := project.ModifyProject()
	if nil != err {
		c.String(500, "更新项目操作失败:"+err.Error())
		return
	}

	c.String(200, "更新项目成功:"+strconv.Itoa(int(affected)))
	return
}

// DeleteProject destroy a project
func DeleteProject(c *gin.Context) {
	var project models.Project
	project.ID = c.Param("id")
	affected, err := project.Destroy()
	if nil != err {
		c.String(500, "删除项目操作失败:"+err.Error())
		return
	}

	c.String(200, "删除项目成功:"+strconv.Itoa(int(affected)))
	return
}

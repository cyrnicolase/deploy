package project

import (
	"net/http"
	"strconv"

	"deploy/models"

	"github.com/gin-gonic/gin"
)

// PostServer 新增服务器
func PostServer(c *gin.Context) {
	var server models.Server
	//rawData, _ := c.GetRawData()			// 这个地方不能先获取原始数据，获取后，下面的Bind会出错
	//logger.Log.Info("新增服务器参数："+string(rawData), nil)
	if err := c.ShouldBindJSON(&server); nil != err {
		c.String(http.StatusBadRequest, "新增服务器参数错误:"+err.Error())
		return
	}

	affected, err := server.Create()
	if nil != err {
		c.String(http.StatusInternalServerError, "新增服务器操作失败:"+err.Error())
		return
	}

	c.String(http.StatusOK, "新增服务器成功:"+strconv.Itoa(int(affected)))
	return
}

// PutServer 修改服务器
func PutServer(c *gin.Context) {
	var server = new(models.Server)
	if err := c.ShouldBind(server); nil != err {
		c.String(http.StatusBadRequest, "修改服务器参数错误:"+err.Error())
		return
	}

	server.ID = c.Param("id")
	affected, err := server.ModifyServer()
	if nil != err {
		c.String(http.StatusInternalServerError, "修改服务器操作失败:"+err.Error())
		return
	}

	c.String(http.StatusOK, "修改服务器成功:"+strconv.Itoa(int(affected)))
	return
}

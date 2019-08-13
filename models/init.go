package models

import (
	"deploy/config"

	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq" // 引入pq的连接
	"xorm.io/core"
)

var x *xorm.Engine

// Boot 是启动
func Boot() {
	engine, err := xorm.NewEngine(config.Database.Dialect, config.Database.String())
	if nil != err {
		panic("连接数据库失败: " + err.Error())
	}

	engine.ShowSQL(true)
	engine.SetMapper(core.GonicMapper{})

	x = engine
}

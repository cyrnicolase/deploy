package config

import (
	"fmt"
	"strings"
)

// DatabaseConfig 是app的配置结构体
type DatabaseConfig struct {
	Dialect  string
	Host     string
	Port     int
	User     string
	Dbname   string `mapstructure:"database"`
	Password string
	Sslmode  string
}

// DatabaseConfigs 是整个数据库配置文件结构体
type DatabaseConfigs struct {
	Development DatabaseConfig
	Testing     DatabaseConfig
	Production  DatabaseConfig
}

// Databases 是所有环境的配置情况
var Databases DatabaseConfigs

// Database 是DatabaseConfig的具体对象
var Database DatabaseConfig

// 解析数据库配置
func unmarshalDatabases() {
	v, _ := readYaml("database")
	v.Unmarshal(&Databases)

	// 注册当前运行模式的数据库配置信息
	registerDatabase()
}

func (d DatabaseConfig) String() string {
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%v",
		d.Host, d.Port, d.User, d.Dbname, d.Password, d.Sslmode)
}

func registerDatabase() {
	runMode := App.RunMode
	switch strings.ToLower(runMode) {
	case "development":
		Database = Databases.Development
	case "testing":
		Database = Databases.Testing
	case "production":
		Database = Databases.Production
	default:
		Database = Databases.Development
	}
}

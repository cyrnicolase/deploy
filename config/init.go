package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var configPath string

func init() {
	currentDir, _ := os.Getwd()
	configPath = currentDir + "/yaml"
}

// Boot is a bootstrap func
func Boot() {
	unmarshalApp()
	unmarshalDatabases()
	unmarshalPrivileges()
}

// SetConfigPath is reset config files path
func SetConfigPath(path string) {
	configPath = path
}

func readYaml(name string) (v *viper.Viper, err error) {
	v = viper.New()
	v.SetConfigName(name)
	v.AddConfigPath(configPath)
	err = v.ReadInConfig()
	if nil != err {
		err = fmt.Errorf("读取配置文件%s失败:%v", name, err)
		return
	}
	return
}

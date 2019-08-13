package config

// AppConfig 是app的配置结构体
type AppConfig struct {
	AppName    string `mapstructure:"app_name"`
	Version    string `mapstructure:"version"`
	DateFormat string `mapstructure:"date_format"`
	RunMode    string `mapstructure:"run_mode"`
}

// App 是AppConfig的具体对象
var App AppConfig

func unmarshalApp() {
	v, _ := readYaml("app")
	v.Unmarshal(&App)
}

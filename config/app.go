package config

// AppConfig 是app的配置结构体
type AppConfig struct {
	AppName    string `mapstructure:"app_name"`
	AppKey     string `mapstructure:"app_key"`
	Version    string `mapstructure:"version"`
	DateFormat string `mapstructure:"date_format"`
	RunMode    string `mapstructure:"run_mode"`
	Jwt        JwtConfig
}

// JwtConfig 是JwtToken相关配置
type JwtConfig struct {
	Timeout int `mapstructure:"timeout"`
	Refresh int `mapstructure:"refresh"`
}

// App 是AppConfig的具体对象
var App AppConfig

func unmarshalApp() {
	v, _ := readYaml("app")
	v.Unmarshal(&App)
}

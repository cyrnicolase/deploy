package config

import (
	"testing"
)

func TestUnmarshalApp(t *testing.T) {
	unmarshalApp()

	tApp := AppConfig{
		AppName:    "一个应用名",
		AppKey:     "base64:mIIRo13uoF/m0Oze81WmlLLQxzGcMZVL+Gr0qSKhivc=",
		Version:    "0.0.1",
		DateFormat: "2006-01-02 15:04:05",
		RunMode:    "development",
		Jwt: JwtConfig{
			Timeout: 120,
			Refresh: 120,
		},
	}

	if tApp != App {
		t.Errorf("期望：%v, 实际:%v\n", tApp, App)
	}
}

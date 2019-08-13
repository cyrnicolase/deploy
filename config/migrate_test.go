package config

import (
	"testing"
)

func TestUnmarshalMigrate(t *testing.T) {
	unmarshalMigrates()
	App.RunMode = "development"
	tMigrates := MigrateConfigs{
		Development: MigrateConfig{
			Dialect:    "postgres",
			Datasource: "host=127.0.0.1 port=5432 user=opmp password=123456 dbname=opmp sslmode=disable",
			Dir:        "migrations",
			Table:      "migrations",
		},
	}

	if tMigrates != Migrates {
		t.Errorf("期望：%v, 实际:%v", tMigrates, Migrates)
	}
}

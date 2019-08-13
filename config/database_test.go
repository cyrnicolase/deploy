package config

import (
	"testing"
)

func TestUnmarshalDatabases(t *testing.T) {
	App.RunMode = "development"
	unmarshalDatabases()
	tDatabases := DatabaseConfigs{
		Development: DatabaseConfig{
			Dialect:  "postgres",
			Host:     "127.0.0.1",
			Port:     5432,
			User:     "opmp",
			Dbname:   "opmp",
			Password: "123456",
			Sslmode:  "disable",
		},
		Testing: DatabaseConfig{
			Dialect:  "postgres",
			Host:     "192.168.0.11",
			Port:     5432,
			User:     "opmp",
			Dbname:   "opmp",
			Password: "abc123",
			Sslmode:  "disable",
		},
		Production: DatabaseConfig{
			Dialect:  "postgres",
			Host:     "10.0.0.1",
			Port:     5432,
			User:     "opmp",
			Dbname:   "opmp",
			Password: "hook@761&%ass,11",
			Sslmode:  "disable",
		},
	}

	if tDatabases != Databases {
		t.Errorf("期望： %v, 实际： %v", tDatabases, Databases)
	}
}

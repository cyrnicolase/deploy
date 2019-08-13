package config

import "strings"

// MigrateConfig is struct
type MigrateConfig struct {
	Dialect    string
	Datasource string
	Dir        string
	Table      string
}

// MigrateConfigs is set of MigrateConfig
type MigrateConfigs struct {
	Development MigrateConfig
}

// Migrates is var of MigrateConfigs
var Migrates MigrateConfigs

// Migrate is var of MigrateConfig
var Migrate MigrateConfig

func unmarshalMigrates() {
	v, _ := readYaml("migrate")
	v.Unmarshal(&Migrates)

	registerMigrate()
}

func registerMigrate() {
	runMode := App.RunMode
	switch strings.ToLower(runMode) {
	case "development":
		Migrate = Migrates.Development
	default:
		Migrate = Migrates.Development
	}
}

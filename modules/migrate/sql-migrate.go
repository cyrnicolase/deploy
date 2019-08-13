package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"

	"deploy/config"
	"deploy/models"

	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
)

var opt string

func init() {
	flag.StringVar(&opt, "opt", "", "数据库迁移操作：up/down")
}

func main() {
	flag.Parse()
	Init()

	migrate.SetTable("migrations")
	migrations := &migrate.FileMigrationSource{
		Dir: "../../migrations",
	}

	db, err := sql.Open(config.Database.Dialect, config.Database.String())
	if nil != err {
		panic("连接数据库操作失败:" + err.Error())
	}

	var (
		n int
	)
	if "up" == opt {
		n, err = migrate.Exec(db, config.Database.Dialect, migrations, migrate.Up)
	}
	if "down" == opt {
		n, err = migrate.Exec(db, config.Database.Dialect, migrations, migrate.Down)
	}

	if nil != err {
		fmt.Errorf("执行迁移操作失败: %v\n", err)
	}

	fmt.Printf("Applied %d migrations!\n", n)
}

// Init is a initial function
func Init() {
	currentDir, _ := os.Getwd()
	path := currentDir + "/../../config/yaml"
	config.SetConfigPath(path)
	config.Boot()
	models.Boot()
}

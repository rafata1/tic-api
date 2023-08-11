package main

import (
	"fmt"
	"git.teko.vn/dung.cda/tic-26-be/config"
	"git.teko.vn/dung.cda/tic-26-be/package/migration"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	conf := config.Load()
	cmd := migration.MigrateCommand(conf.MySQL.DSN())
	err := cmd.Execute()
	if err != nil {
		fmt.Println("ERROR:", err)
	}
}

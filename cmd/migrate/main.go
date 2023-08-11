package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/rafata1/tic-api/config"
	"github.com/rafata1/tic-api/package/migration"
)

func main() {
	conf := config.Load()
	cmd := migration.MigrateCommand(conf.MySQL.DSN())
	err := cmd.Execute()
	if err != nil {
		fmt.Println("ERROR:", err)
	}
}

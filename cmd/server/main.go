package main

import (
	"fmt"
	"git.teko.vn/dung.cda/tic-26-be/config"
	"git.teko.vn/dung.cda/tic-26-be/service/project"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"
	"os"
)

func main() {
	rootCmd := cobra.Command{
		Use: "server",
	}
	rootCmd.AddCommand(
		runServerCommand(),
	)
	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}

func runServerCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "run",
		Short: "run http server",
		Run: func(cmd *cobra.Command, args []string) {
			runServer()
		},
	}
}

func runServer() {
	conf := config.Load()
	db := conf.MySQL.MustConnect()
	fmt.Println("Connected to MySQL")

	gin.SetMode(os.Getenv("GIN_MODE"))
	router := gin.Default()
	router.Use(cors.Default())

	projectServer := project.NewServer(db)
	router.POST("/api/v1/projects", projectServer.CreateProject)

	if err := router.Run(); err != nil {
		panic(err)
	}
}

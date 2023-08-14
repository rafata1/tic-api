package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rafata1/tic-api/config"
	"github.com/rafata1/tic-api/service/auth"
	"github.com/rafata1/tic-api/service/chat"
	"github.com/rafata1/tic-api/service/project"
	"github.com/rafata1/tic-api/service/ticket"
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
	router.Use(getCorsMiddleware())
	authService := auth.NewService(conf.IAM.Endpoint)
	router.Use(authService.AuthenticationInterceptor())

	projectServer := project.NewServer(db)
	router.GET("/api/v1/projects/:project_id/faqs", projectServer.ListFAQs)
	router.POST("/api/v1/projects/:project_id/faqs", projectServer.CreateFAQ)
	router.POST("/api/v1/projects", projectServer.CreateProject)
	router.GET("/api/v1/projects/:project_id", projectServer.GetProject)
	router.GET("/api/v1/projects", projectServer.ListProjects)

	chatServer := chat.NewServer()
	router.POST("/api/v1/chat", chatServer.Answer)

	ticketServer := ticket.NewServer(conf)
	router.POST("/api/v1/tickets", ticketServer.CreateTicket)

	if err := router.Run(); err != nil {
		panic(err)
	}
}

func getCorsMiddleware() gin.HandlerFunc {
	conf := cors.DefaultConfig()
	conf.AllowAllOrigins = true
	conf.AllowHeaders = []string{
		"Origin",
		"Accept",
		"Content-Type",
		"Content-Length",
		"Accept-Encoding",
		"Authorization",
		"X-CSRF-Token",
	}
	return cors.New(conf)
}

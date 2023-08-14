package ticket

import (
	"github.com/andygrunwald/go-jira"
	"github.com/gin-gonic/gin"
	"github.com/rafata1/tic-api/config"
	"github.com/rafata1/tic-api/service/common"
)

type IServer interface {
	CreateTicket(c *gin.Context)
}

type server struct {
	service IService
}

func (s server) CreateTicket(c *gin.Context) {
	var input inputTicket
	err := c.Bind(&input)
	if err != nil {
		common.WriteError(c, common.ErrBadRequest)
		return
	}
	output, err := s.service.CreateTicket(c, input)
	if err != nil {
		common.WriteError(c, err)
		return
	}
	common.WriteSuccess(c, output)
}

func NewServer(conf config.Config) IServer {
	tp := jira.BasicAuthTransport{
		Username: conf.Jira.Username,
		Password: conf.Jira.Password,
	}

	jiraClient, err := jira.NewClient(tp.Client(), conf.Jira.BaseURL)
	if err != nil {
		panic(err)
	}

	return &server{
		service: NewService(jiraClient),
	}
}

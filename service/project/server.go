package project

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/rafata1/tic-api/service/common"
)

type IServer interface {
	CreateProject(c *gin.Context)
	ListProjects(c *gin.Context)
}

type server struct {
	service IService
}

func (s server) CreateProject(c *gin.Context) {
	var dtoProject DTOProject
	err := c.Bind(&dtoProject)
	if err != nil {
		common.WriteError(c, err)
		return
	}

	output, err := s.service.CreateProject(c, dtoProject.Name)
	if err != nil {
		common.WriteError(c, err)
		return
	}

	common.WriteSuccess(c, output)
	return
}

func (s server) ListProjects(c *gin.Context) {
	output, err := s.service.ListProject(c)
	if err != nil {
		common.WriteError(c, err)
		return
	}

	common.WriteSuccess(c, output)
	return
}

func NewServer(db *sqlx.DB) IServer {
	return &server{
		service: NewService(db),
	}
}

package project

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/rafata1/tic-api/service/common"
)

type IServer interface {
	CreateProject(c *gin.Context)
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

	ctx := context.Background()
	output, err := s.service.CreateProject(ctx, dtoProject.Name)
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

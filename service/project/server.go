package project

import (
	"context"
	"git.teko.vn/dung.cda/tic-26-be/service/common"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
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

package project

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/rafata1/tic-api/service/common"
	"github.com/rafata1/tic-api/util"
)

type IServer interface {
	CreateProject(c *gin.Context)
	ListProjects(c *gin.Context)
	GetProject(c *gin.Context)
	CreateFAQ(c *gin.Context)
	ListFAQs(c *gin.Context)
}

type server struct {
	service IService
}

func (s server) CreateProject(c *gin.Context) {
	var input inputProject
	err := c.Bind(&input)
	if err != nil {
		common.WriteError(c, common.ErrBadRequest)
		return
	}

	output, err := s.service.CreateProject(c, input.Name)
	if err != nil {
		common.WriteError(c, err)
		return
	}

	common.WriteSuccess(c, output)
}

func (s server) ListProjects(c *gin.Context) {
	output, err := s.service.ListProject(c)
	if err != nil {
		common.WriteError(c, err)
		return
	}

	common.WriteSuccess(c, output)
}

func (s server) GetProject(c *gin.Context) {
	projectID, err := util.ParseInt64(
		c.Param("project_id"),
	)
	if err != nil {
		common.WriteError(c, common.ErrBadRequest)
		return
	}

	output, err := s.service.GetProject(c, projectID)
	if err != nil {
		common.WriteError(c, err)
		return
	}
	common.WriteSuccess(c, output)
}

func (s server) CreateFAQ(c *gin.Context) {
	projectID, err := util.ParseInt64(
		c.Param("project_id"),
	)
	if err != nil {
		common.WriteError(c, common.ErrBadRequest)
		return
	}

	var input inputFAQ
	err = c.Bind(&input)
	if err != nil {
		common.WriteError(c, common.ErrBadRequest)
		return
	}

	input.ProjectID = projectID
	output, err := s.service.CreateFAQ(c, input)
	if err != nil {
		common.WriteError(c, err)
		return
	}
	common.WriteSuccess(c, output)
}

func (s server) ListFAQs(c *gin.Context) {
	projectID, err := util.ParseInt64(
		c.Param("project_id"),
	)
	if err != nil {
		common.WriteError(c, common.ErrBadRequest)
		return
	}
	output, err := s.service.ListFAQs(c, projectID)
	if err != nil {
		common.WriteError(c, err)
		return
	}
	common.WriteSuccess(c, output)
}

func NewServer(db *sqlx.DB) IServer {
	return &server{
		service: NewService(db),
	}
}

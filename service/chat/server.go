package chat

import (
	"github.com/gin-gonic/gin"
	"github.com/rafata1/tic-api/service/common"
)

type IServer interface {
	Answer(c *gin.Context)
}

type server struct {
	service IService
}

func (s server) Answer(c *gin.Context) {
	var input inputChat
	err := c.Bind(&input)
	if err != nil {
		common.WriteError(c, common.ErrBadRequest)
		return
	}

	output, err := s.service.Answer(c, input)
	if err != nil {
		common.WriteError(c, err)
		return
	}
	common.WriteSuccess(c, output)
}

func NewServer() IServer {
	return &server{
		service: NewService(),
	}
}

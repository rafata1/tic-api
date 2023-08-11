package common

import (
	"github.com/gin-gonic/gin"
	"github.com/rafata1/tic-api/package/errors"
	"net/http"
)

func WriteError(c *gin.Context, err error) {
	code := err.(errors.DomainError).Code
	c.JSON(code, BaseRes{
		Message: err.Error(),
	})
}

func WriteSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, BaseRes{
		Message: "success",
		Data:    data,
	})
}

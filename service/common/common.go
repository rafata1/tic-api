package common

import (
	"git.teko.vn/dung.cda/tic-26-be/package/errors"
	"github.com/gin-gonic/gin"
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

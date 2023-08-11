package common

import (
	"git.teko.vn/dung.cda/tic-26-be/package/errors"
	"net/http"
)

var ErrExecuteIntoDB = errors.New(http.StatusInternalServerError, "errors executing into database")

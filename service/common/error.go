package common

import (
	"github.com/rafata1/tic-api/package/errors"
	"net/http"
)

var ErrExecuteIntoDB = errors.New(http.StatusInternalServerError, "errors executing into database")

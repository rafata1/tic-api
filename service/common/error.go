package common

import (
	"github.com/rafata1/tic-api/package/errors"
	"net/http"
)

var ErrBadRequest = errors.New(http.StatusBadRequest, "bad request")
var ErrUnauthorized = errors.New(http.StatusUnauthorized, "unauthorized")
var ErrExecuteIntoDB = errors.New(http.StatusInternalServerError, "errors executing into database")
var ErrQueryIntoDB = errors.New(http.StatusInternalServerError, "errors querying in to database")

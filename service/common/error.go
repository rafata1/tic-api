package common

import (
	"github.com/rafata1/tic-api/package/errors"
	"net/http"
)

var ErrBadRequest = errors.New(http.StatusBadRequest, "bad request")
var ErrUnauthorized = errors.New(http.StatusUnauthorized, "unauthorized")
var ErrExecuteIntoDB = errors.New(http.StatusInternalServerError, "error executing into database")
var ErrQueryIntoDB = errors.New(http.StatusInternalServerError, "error querying in to database")
var ErrCall3rdParty = errors.New(http.StatusInternalServerError, "error calling 3rd")

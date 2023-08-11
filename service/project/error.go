package project

import (
	"github.com/rafata1/tic-api/package/errors"
	"net/http"
)

var ErrProjectNameRequired = errors.New(http.StatusBadRequest, "project name is required")

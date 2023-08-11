package project

import (
	"git.teko.vn/dung.cda/tic-26-be/package/errors"
	"net/http"
)

var ErrProjectNameRequired = errors.New(http.StatusBadRequest, "project name is required")

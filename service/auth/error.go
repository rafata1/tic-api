package auth

import "github.com/rafata1/tic-api/package/errors"

// ErrInvalidIAMToken ...
var ErrInvalidIAMToken = errors.New(401, "Invalid iam token")

// ErrCallIAM ...
var ErrCallIAM = errors.New(500, "Error calling iam service")

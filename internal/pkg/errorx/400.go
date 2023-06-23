package errorx

import (
	"net/http"

	"github.com/blackhorseya/project-goapi/pkg/er"
)

var (
	// ErrInvalidUsername means the username is invalid
	ErrInvalidUsername = er.New(http.StatusBadRequest, 400100, "invalid username", "invalid username")

	// ErrInvalidPassword means the password is invalid
	ErrInvalidPassword = er.New(http.StatusBadRequest, 400101, "invalid password", "invalid password")
)

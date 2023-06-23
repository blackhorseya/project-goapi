package errorx

import (
	"net/http"

	"github.com/hashgreen/hg-goapi/pkg/er"
)

var (
	// ErrGetProfile means get profile error
	ErrGetProfile = er.New(http.StatusInternalServerError, 500100, "get profile error", "get profile error")
)

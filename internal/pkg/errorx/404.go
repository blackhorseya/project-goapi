package errorx

import (
	"net/http"

	"github.com/hashgreen/hg-goapi/pkg/er"
)

var (
	// ErrNotFoundProfile means can't find profile
	ErrNotFoundProfile = er.New(http.StatusNotFound, 404100, "can't find profile", "can't find profile")
)

//go:generate mockgen -destination=./mock_${GOFILE} -package=biz -source=${GOFILE}

package biz

import (
	userM "github.com/blackhorseya/project-goapi/entity/domain/user/model"
	"github.com/blackhorseya/project-goapi/pkg/contextx"
)

// IBiz is a biz interface
type IBiz interface {
	// Login serve caller to give username and password to login system
	Login(ctx contextx.Contextx, username, password string) (info *userM.Profile, err error)

	// Register serve caller to give username and password to register system
	Register(ctx contextx.Contextx, username, password string) (info *userM.Profile, err error)
}

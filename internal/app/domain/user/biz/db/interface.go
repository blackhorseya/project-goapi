//go:generate mockgen -destination=./mock_${GOFILE} -package=db -source=${GOFILE}

package db

import (
	userM "github.com/blackhorseya/project-goapi/entity/domain/user/model"
	"github.com/blackhorseya/project-goapi/pkg/contextx"
)

// Reader declare a reader interface
type Reader interface {
	// GetUserByUsername serve caller to get user profile by username
	GetUserByUsername(ctx contextx.Contextx, username string) (info *userM.Profile, err error)
}

// Writer declare a writer interface
type Writer interface {
	// CreateUser serve caller to create a user profile
	CreateUser(ctx contextx.Contextx, username, password string) (info *userM.Profile, err error)

	// UpdateUser serve caller to update a user profile
	UpdateUser(ctx contextx.Contextx, user *userM.Profile) error
}

// ReaderWriter declare a reader and writer interface
type ReaderWriter interface {
	Reader
	Writer
}

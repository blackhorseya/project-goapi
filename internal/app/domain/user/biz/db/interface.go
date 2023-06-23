//go:generate mockgen -destination=./mock_${GOFILE} -package=db -source=${GOFILE}

package db

import (
	userM "github.com/blackhorseya/project-goapi/entity/domain/user/model"
	"github.com/blackhorseya/project-goapi/pkg/contextx"
)

// Reader declare a reader interface
type Reader interface {
	// GetUserByName serve caller to get user profile by name
	GetUserByName(ctx contextx.Contextx, name string) (info *userM.Profile, err error)
}

// Writer declare a writer interface
type Writer interface {
	// CreateUser serve caller to create a user profile
	CreateUser(ctx contextx.Contextx, name, password string) (info *userM.Profile, err error)

	// UpdateUser serve caller to update a user profile
	UpdateUser(ctx contextx.Contextx, user *userM.Profile) error
}

// ReaderWriter declare a reader and writer interface
type ReaderWriter interface {
	Reader
	Writer
}

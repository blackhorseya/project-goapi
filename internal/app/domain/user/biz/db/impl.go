package db

import (
	"github.com/blackhorseya/project-goapi/entity/domain/user/model"
	"github.com/blackhorseya/project-goapi/pkg/contextx"
)

type impl struct {
}

// NewImpl new a db impl
func NewImpl() ReaderWriter {
	return &impl{}
}

func (i *impl) GetUserByUsername(ctx contextx.Contextx, username string) (info *model.Profile, err error) {
	// todo: 2023/6/24|sean|impl me
	panic("implement me")
}

func (i *impl) CreateUser(ctx contextx.Contextx, username, password string) (info *model.Profile, err error) {
	// todo: 2023/6/24|sean|impl me
	panic("implement me")
}

func (i *impl) UpdateUser(ctx contextx.Contextx, user *model.Profile) error {
	// todo: 2023/6/24|sean|impl me
	panic("implement me")
}

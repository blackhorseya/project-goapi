package biz

import (
	userB "github.com/blackhorseya/project-goapi/entity/domain/user/biz"
	"github.com/blackhorseya/project-goapi/entity/domain/user/model"
	"github.com/blackhorseya/project-goapi/internal/app/domain/user/biz/db"
	"github.com/blackhorseya/project-goapi/pkg/contextx"
)

type impl struct {
	rw db.ReaderWriter
}

// NewImpl serve caller to create a user biz
func NewImpl(rw db.ReaderWriter) userB.Bizer {
	return &impl{
		rw: rw,
	}
}

func (i *impl) Login(ctx contextx.Contextx, username, password string) (info *model.Profile, err error) {
	// todo: 2023/6/24|sean|impl me
	panic("implement me")
}

func (i *impl) Register(ctx contextx.Contextx, username, password string) (info *model.Profile, err error) {
	// todo: 2023/6/24|sean|impl me
	panic("implement me")
}

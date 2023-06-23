package biz

import (
	"strings"

	userB "github.com/blackhorseya/project-goapi/entity/domain/user/biz"
	"github.com/blackhorseya/project-goapi/entity/domain/user/model"
	"github.com/blackhorseya/project-goapi/internal/app/domain/user/biz/db"
	"github.com/blackhorseya/project-goapi/internal/pkg/errorx"
	"github.com/blackhorseya/project-goapi/pkg/contextx"
	"go.uber.org/zap"
)

type impl struct {
	db db.ReaderWriter
}

// NewImpl serve caller to create a user biz
func NewImpl(db db.ReaderWriter) userB.Bizer {
	return &impl{
		db: db,
	}
}

func (i *impl) Login(ctx contextx.Contextx, username, password string) (info *model.Profile, err error) {
	username = strings.ReplaceAll(username, " ", "")
	if len(username) == 0 {
		ctx.Error("error empty username", zap.String("username", username))
		return nil, errorx.ErrInvalidUsername
	}

	profile, err := i.db.GetUserByUsername(ctx, username)
	if err != nil {
		ctx.Error("get user by username error from db", zap.Error(err), zap.String("username", username))
		return nil, errorx.ErrGetProfile
	}
	if profile == nil {
		ctx.Error("can't find profile from db", zap.String("username", username))
		return nil, errorx.ErrNotFoundProfile
	}

	// todo: 2023/6/24|sean|impl me
	panic("implement me")
}

func (i *impl) Register(ctx contextx.Contextx, username, password string) (info *model.Profile, err error) {
	// todo: 2023/6/24|sean|impl me
	panic("implement me")
}

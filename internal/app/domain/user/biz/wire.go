//go:generate wire
//go:build wireinject

package biz

import (
	userB "github.com/blackhorseya/project-goapi/entity/domain/user/biz"
	"github.com/blackhorseya/project-goapi/internal/app/domain/user/biz/db"
	"github.com/blackhorseya/project-goapi/internal/pkg/storage/mongodb"
	"github.com/google/wire"
)

var testProviderSet = wire.NewSet(NewImpl)

func CreateBiz(db db.ReaderWriter) userB.IBiz {
	panic(wire.Build(testProviderSet))
}

var integrationProviderSet = wire.NewSet(NewImpl, db.ProviderSet, mongodb.ProviderSet)

func CreateBizIntegration(opts *mongodb.Options) (userB.IBiz, error) {
	panic(wire.Build(integrationProviderSet))
}

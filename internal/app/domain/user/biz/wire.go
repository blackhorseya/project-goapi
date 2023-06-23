//go:generate wire
//go:build wireinject

package biz

import (
	userB "github.com/blackhorseya/project-goapi/entity/domain/user/biz"
	"github.com/blackhorseya/project-goapi/internal/app/domain/user/biz/db"
	"github.com/google/wire"
)

var testProviderSet = wire.NewSet(NewImpl)

func CreateBizer(db db.ReaderWriter) userB.Bizer {
	panic(wire.Build(testProviderSet))
}

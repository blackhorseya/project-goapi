// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package biz

import (
	"github.com/blackhorseya/project-goapi/entity/domain/user/biz"
	"github.com/blackhorseya/project-goapi/internal/app/domain/user/biz/db"
	"github.com/blackhorseya/project-goapi/internal/pkg/storage/mongodb"
	"github.com/google/wire"
)

// Injectors from wire.go:

func CreateBiz(db2 db.ReaderWriter) biz.IBiz {
	iBiz := NewImpl(db2)
	return iBiz
}

func CreateBizIntegration(opts *mongodb.Options) (biz.IBiz, error) {
	client, err := mongodb.NewMongoClient(opts)
	if err != nil {
		return nil, err
	}
	readerWriter := db.NewImpl(client)
	iBiz := NewImpl(readerWriter)
	return iBiz, nil
}

// wire.go:

var testProviderSet = wire.NewSet(NewImpl)

var integrationProviderSet = wire.NewSet(NewImpl, db.ProviderSet, mongodb.ProviderSet)

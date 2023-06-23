// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package db

import (
	"github.com/blackhorseya/project-goapi/internal/pkg/storage/mongodb"
	"github.com/google/wire"
)

// Injectors from wire.go:

func CreateReaderWriter(opts *mongodb.Options) (ReaderWriter, error) {
	client, err := mongodb.NewMongoClient(opts)
	if err != nil {
		return nil, err
	}
	readerWriter := NewImpl(client)
	return readerWriter, nil
}

// wire.go:

var testProviderSet = wire.NewSet(NewImpl, mongodb.ProviderSet)

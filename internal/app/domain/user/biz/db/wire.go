//go:generate wire
//go:build wireinject

package db

import (
	"github.com/blackhorseya/project-goapi/internal/pkg/storage/mongodb"
	"github.com/google/wire"
)

var testProviderSet = wire.NewSet(NewImpl, mongodb.ProviderSet)

func CreateReaderWriter(opts *mongodb.Options) (ReaderWriter, error) {
	panic(wire.Build(testProviderSet))
}

package mongodb

import (
	"github.com/google/wire"
	"github.com/hashgreen/hg-goapi/pkg/contextx"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Options is the options for mongodb.
type Options struct {
	URL string `json:"url" yaml:"url"`
}

// NewMongoClient creates a new mongo client.
func NewMongoClient(opts *Options) (*mongo.Client, error) {
	client, err := mongo.Connect(contextx.Background(), options.Client().ApplyURI(opts.URL))
	if err != nil {
		return nil, errors.Wrap(err, "connect mongo client failed")
	}

	return client, nil
}

// ProviderSet is mongodb providers
var ProviderSet = wire.NewSet(NewMongoClient)

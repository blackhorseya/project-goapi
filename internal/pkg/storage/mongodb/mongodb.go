package mongodb

import (
	"github.com/blackhorseya/project-goapi/pkg/contextx"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
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

// Container represents the mongodb container type used in the module
type Container struct {
	testcontainers.Container
}

// NewContainer creates an instance of the mongodb container type
func NewContainer(ctx contextx.Contextx) (*Container, error) {
	req := testcontainers.ContainerRequest{
		Image:        "mongo:6",
		ExposedPorts: []string{"27017/tcp"},
		WaitingFor: wait.ForAll(
			wait.ForLog("Waiting for connections"),
			wait.ForListeningPort("27017/tcp"),
		),
	}
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, err
	}

	return &Container{Container: container}, nil
}

// ProviderSet is mongodb providers
var ProviderSet = wire.NewSet(NewMongoClient)

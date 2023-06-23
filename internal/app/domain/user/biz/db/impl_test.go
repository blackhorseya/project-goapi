package db

import (
	"context"
	"reflect"
	"testing"

	userM "github.com/blackhorseya/project-goapi/entity/domain/user/model"
	"github.com/blackhorseya/project-goapi/internal/pkg/storage/mongodb"
	"github.com/blackhorseya/project-goapi/pkg/contextx"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

// mongodbContainer represents the mongodb container type used in the module
type mongodbContainer struct {
	testcontainers.Container
}

// startContainer creates an instance of the mongodb container type
func startContainer(ctx context.Context) (*mongodbContainer, error) {
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

	return &mongodbContainer{Container: container}, nil
}

type suiteTester struct {
	suite.Suite

	logger  *zap.Logger
	mongodb *mongodbContainer
	rw      *mongo.Client
	db      ReaderWriter
}

func (s *suiteTester) SetupTest() {
	s.logger, _ = zap.NewDevelopment()

	ctx := contextx.BackgroundWithLogger(s.logger)
	s.mongodb, _ = startContainer(ctx)
	endpoint, _ := s.mongodb.Endpoint(ctx, "mongodb")
	s.rw, _ = mongo.Connect(ctx, options.Client().ApplyURI(endpoint))

	opts := &mongodb.Options{URL: endpoint}
	db, err := CreateReaderWriter(opts)
	if err != nil {
		panic(err)
	}
	s.db = db
}

func (s *suiteTester) TearDownTest() {
	_ = s.rw.Disconnect(context.Background())
	_ = s.mongodb.Terminate(context.Background())
}

func TestAll(t *testing.T) {
	t.Parallel()
	t.Helper()

	suite.Run(t, new(suiteTester))
}

func (s *suiteTester) Test_impl_GetUserByUsername() {
	user1 := &userM.Profile{
		Id:          0,
		Username:    "user1",
		Password:    "",
		Name:        "",
		Email:       "",
		Picture:     "",
		AccessToken: "",
		CreatedAt:   nil,
		UpdatedAt:   nil,
	}

	type args struct {
		username string
		mock     func()
	}
	tests := []struct {
		name     string
		args     args
		wantInfo *userM.Profile
		wantErr  bool
	}{
		{
			name:     "get user by username then not found return nil",
			args:     args{username: "user1"},
			wantInfo: nil,
			wantErr:  false,
		},
		{
			name: "get user by username then ok",
			args: args{username: "user1", mock: func() {
				_, _ = s.rw.Database(dbName).Collection(collName).InsertOne(contextx.Background(), user1)
			}},
			wantInfo: user1,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotInfo, err := s.db.GetUserByUsername(contextx.BackgroundWithLogger(s.logger), tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserByUsername() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInfo, tt.wantInfo) {
				t.Errorf("GetUserByUsername() gotInfo = %v, want %v", gotInfo, tt.wantInfo)
			}
		})
	}
}

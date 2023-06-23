package db

import (
	"time"

	userM "github.com/blackhorseya/project-goapi/entity/domain/user/model"
	"github.com/blackhorseya/project-goapi/pkg/contextx"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	dbName = "user"

	collName = "users"
)

type impl struct {
	rw *mongo.Client
}

// NewImpl new a db impl
func NewImpl(rw *mongo.Client) ReaderWriter {
	return &impl{
		rw: rw,
	}
}

func (i *impl) GetUserByUsername(ctx contextx.Contextx, username string) (info *userM.Profile, err error) {
	timeout, cancelFunc := contextx.WithTimeout(ctx, 5*time.Second)
	defer cancelFunc()

	coll := i.rw.Database(dbName).Collection(collName)
	filter := bson.M{"username": username}
	var got *userM.Profile
	err = coll.FindOne(timeout, filter).Decode(&got)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}

		return nil, errors.Wrap(err, "find one user by username failed")
	}

	return got, nil
}

func (i *impl) CreateUser(ctx contextx.Contextx, username, password string) (info *userM.Profile, err error) {
	// todo: 2023/6/24|sean|impl me
	panic("implement me")
}

func (i *impl) UpdateUser(ctx contextx.Contextx, user *userM.Profile) error {
	// todo: 2023/6/24|sean|impl me
	panic("implement me")
}

package biz

import (
	"testing"

	userB "github.com/blackhorseya/project-goapi/entity/domain/user/biz"
	"github.com/blackhorseya/project-goapi/internal/pkg/storage/mongodb"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type suiteIntegration struct {
	suite.Suite

	logger  *zap.Logger
	mongodb *mongodb.Container
	rw      *mongo.Client
	biz     userB.IBiz
}

func (s *suiteIntegration) SetupTest() {
}

func (s *suiteIntegration) TearDownTest() {
}

func TestIntegration(t *testing.T) {
	t.Parallel()
	t.Helper()

	suite.Run(t, new(suiteIntegration))
}

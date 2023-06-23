package biz

import (
	"errors"
	"reflect"
	"testing"

	userB "github.com/blackhorseya/project-goapi/entity/domain/user/biz"
	userM "github.com/blackhorseya/project-goapi/entity/domain/user/model"
	"github.com/blackhorseya/project-goapi/internal/app/domain/user/biz/db"
	"github.com/blackhorseya/project-goapi/pkg/contextx"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type suiteTester struct {
	suite.Suite

	logger *zap.Logger
	ctrl   *gomock.Controller
	db     *db.MockReaderWriter
	biz    userB.IBiz
}

func (s *suiteTester) SetupTest() {
	s.logger, _ = zap.NewDevelopment()
	s.ctrl = gomock.NewController(s.T())
	s.db = db.NewMockReaderWriter(s.ctrl)
	s.biz = CreateBiz(s.db)
}

func (s *suiteTester) TearDownTest() {
	s.ctrl.Finish()
}

func TestAll(t *testing.T) {
	t.Parallel()
	t.Helper()

	suite.Run(t, new(suiteTester))
}

func (s *suiteTester) Test_impl_Login() {
	type args struct {
		username string
		password string
		mock     func()
	}
	tests := []struct {
		name     string
		args     args
		wantInfo *userM.Profile
		wantErr  bool
	}{
		{
			name:     "invalid username then error",
			args:     args{username: "   ", password: "password"},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "get user by username then error",
			args: args{username: "username", password: "password", mock: func() {
				s.db.EXPECT().GetUserByUsername(gomock.Any(), "username").Return(nil, errors.New("error")).Times(1)
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "get user by username not found then error",
			args: args{username: "username", password: "password", mock: func() {
				s.db.EXPECT().GetUserByUsername(gomock.Any(), "username").Return(nil, nil).Times(1)
			}},
			wantInfo: nil,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotInfo, err := s.biz.Login(contextx.BackgroundWithLogger(s.logger), tt.args.username, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInfo, tt.wantInfo) {
				t.Errorf("Login() gotInfo = %v, want %v", gotInfo, tt.wantInfo)
			}
		})
	}
}

package user

import (
	"errors"
	"reflect"
	"testing"

	"github.com/blackhorseya/lobster/internal/biz/user/repo/mocks"
	"github.com/blackhorseya/lobster/internal/pkg/contextx"
	"github.com/blackhorseya/lobster/internal/pkg/pb"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var (
	uuid1 = "d76f4f51-f141-41ba-ba57-c4749319586b"

	time1 = int64(1610548520788105000)

	token1 = "b54c851b9d9e030f2afd6f6119b9c84e59f02590"

	email1 = "test@gmail.com"

	user1 = pb.Profile{
		ID:          uuid1,
		AccessToken: token1,
		Email:       email1,
		SignupAt:    time1,
	}
)

type bizSuite struct {
	suite.Suite
	mock *mocks.IRepo
	biz  IBiz
}

func (s *bizSuite) SetupTest() {
	s.mock = new(mocks.IRepo)
	biz, err := CreateIBiz(s.mock)
	if err != nil {
		panic(err)
		return
	}

	s.biz = biz
}

func (s *bizSuite) TearDownTest() {
	s.mock.AssertExpectations(s.T())
}

func TestBizSuite(t *testing.T) {
	suite.Run(t, new(bizSuite))
}

func (s *bizSuite) Test_impl_GetInfoByEmail() {
	type args struct {
		email string
		mock  func()
	}
	tests := []struct {
		name     string
		args     args
		wantInfo *pb.Profile
		wantErr  bool
	}{
		{
			name: "mail then nil error",
			args: args{email: email1, mock: func() {
				s.mock.On("QueryInfoByEmail", mock.Anything, email1).Return(nil, errors.New("error")).Once()
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "mail then profile nil",
			args: args{email: email1, mock: func() {
				s.mock.On("QueryInfoByEmail", mock.Anything, email1).Return(&user1, nil).Once()
			}},
			wantInfo: &user1,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotInfo, err := s.biz.GetInfoByEmail(contextx.Background(), tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetInfoByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInfo, tt.wantInfo) {
				t.Errorf("GetInfoByEmail() gotInfo = %v, want %v", gotInfo, tt.wantInfo)
			}

			s.TearDownTest()
		})
	}
}

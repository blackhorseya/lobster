package user

import (
	"reflect"
	"testing"

	"github.com/blackhorseya/lobster/internal/app/lobster/biz/user/repo/mocks"
	"github.com/blackhorseya/lobster/internal/pkg/base/contextx"
	"github.com/blackhorseya/lobster/internal/pkg/base/encrypt"
	"github.com/blackhorseya/lobster/internal/pkg/entity/user"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var (
	email1 = "test@google.com"

	pass1 = "password"

	encPWD, _ = encrypt.HashAndSalt(pass1)

	user1 = &user.Profile{Email: email1, Password: encPWD}
)

type bizSuite struct {
	suite.Suite
	mock *mocks.IRepo
	biz  IBiz
}

func (s *bizSuite) SetupTest() {
	s.mock = new(mocks.IRepo)
	biz, err := CreateIBiz("../../../../../configs/app.yaml", s.mock)
	if err != nil {
		panic(err)
	}

	s.biz = biz
}

func (s *bizSuite) TearDownTest() {
	s.mock.AssertExpectations(s.T())
}

func TestBizSuite(t *testing.T) {
	suite.Run(t, new(bizSuite))
}

func (s *bizSuite) Test_impl_Signup() {
	type args struct {
		email    string
		password string
		mock     func()
	}
	tests := []struct {
		name     string
		args     args
		wantInfo *user.Profile
		wantErr  bool
	}{
		{
			name:     "missing email then error",
			args:     args{email: "", password: pass1},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name:     "missing password then error",
			args:     args{email: email1, password: ""},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "get by email then error",
			args: args{email: email1, password: pass1, mock: func() {
				s.mock.On("GetByEmail", mock.Anything, email1).Return(nil, errors.New("error")).Once()
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "get by email is exists then error",
			args: args{email: email1, password: pass1, mock: func() {
				s.mock.On("GetByEmail", mock.Anything, email1).Return(user1, nil).Once()
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "signup then error",
			args: args{email: email1, password: pass1, mock: func() {
				s.mock.On("GetByEmail", mock.Anything, email1).Return(nil, nil).Once()
				s.mock.On("Register", mock.Anything, email1, mock.Anything).Return(nil, errors.New("error")).Once()
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "signup then success",
			args: args{email: email1, password: pass1, mock: func() {
				s.mock.On("GetByEmail", mock.Anything, email1).Return(nil, nil).Once()
				s.mock.On("Register", mock.Anything, email1, mock.Anything).Return(user1, nil).Once()
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

			gotInfo, err := s.biz.Signup(contextx.Background(), tt.args.email, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("Signup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInfo, tt.wantInfo) {
				t.Errorf("Signup() gotInfo = %v, want %v", gotInfo, tt.wantInfo)
			}

			s.TearDownTest()
		})
	}
}

func (s *bizSuite) Test_impl_Login() {
	type args struct {
		email    string
		password string
		mock     func()
	}
	tests := []struct {
		name     string
		args     args
		wantInfo *user.Profile
		wantErr  bool
	}{
		{
			name:     "missing email then error",
			args:     args{email: "", password: pass1},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name:     "missing password then error",
			args:     args{email: email1, password: ""},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "get by email then error",
			args: args{email: email1, password: pass1, mock: func() {
				s.mock.On("GetByEmail", mock.Anything, email1).Return(nil, errors.New("error")).Once()
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "get by email then not found error",
			args: args{email: email1, password: pass1, mock: func() {
				s.mock.On("GetByEmail", mock.Anything, email1).Return(nil, nil).Once()
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "login then password not equal error",
			args: args{email: email1, password: "123", mock: func() {
				s.mock.On("GetByEmail", mock.Anything, email1).Return(user1, nil).Once()
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "update token then error",
			args: args{email: email1, password: pass1, mock: func() {
				s.mock.On("GetByEmail", mock.Anything, email1).Return(user1, nil).Once()
				s.mock.On("UpdateToken", mock.Anything, user1).Return(nil, errors.New("error")).Once()
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "login then success",
			args: args{email: email1, password: pass1, mock: func() {
				s.mock.On("GetByEmail", mock.Anything, email1).Return(user1, nil).Once()
				s.mock.On("UpdateToken", mock.Anything, user1).Return(user1, nil).Once()
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

			gotInfo, err := s.biz.Login(contextx.Background(), tt.args.email, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInfo, tt.wantInfo) {
				t.Errorf("Login() gotInfo = %v, want %v", gotInfo, tt.wantInfo)
			}

			s.TearDownTest()
		})
	}
}

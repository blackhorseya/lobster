package user

import (
	"reflect"
	"testing"

	"github.com/blackhorseya/lobster/internal/app/lobster/biz/user/repo/mocks"
	"github.com/blackhorseya/lobster/internal/pkg/base/contextx"
	"github.com/blackhorseya/lobster/internal/pkg/base/encrypt"
	"github.com/blackhorseya/lobster/internal/pkg/entity/user"
	"github.com/blackhorseya/lobster/internal/pkg/infra/token"
	"github.com/bwmarrin/snowflake"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

var (
	id1 = int64(0)

	token1 = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJsb2JzdGVyIiwiaWQiOjAsImVtYWlsIjoiZW1haWwiLCJleHAiOjE5MDk5NTU2NDJ9.I2tByuRnyMtTEOWihGX3_RcKFS-3AwjdRxsW_YzZ-0c"

	email1 = "email"

	pass1 = "password"

	salt1, _ = encrypt.HashAndSalt(pass1)

	info1 = &user.Profile{ID: id1, Token: token1, Password: salt1}
)

type bizSuite struct {
	suite.Suite
	mock *mocks.IRepo
	biz  IBiz
}

func (s *bizSuite) SetupTest() {
	logger, _ := zap.NewDevelopment()
	node, _ := snowflake.NewNode(1)
	factory, _ := token.New(&token.Options{Signature: "7d01eb0200bc730a2c58"}, logger)

	s.mock = new(mocks.IRepo)
	if biz, err := CreateIBiz(logger, s.mock, node, factory); err != nil {
		panic(err)
	} else {
		s.biz = biz
	}
}

func (s *bizSuite) TearDownTest() {
	s.mock.AssertExpectations(s.T())
}

func TestBizSuite(t *testing.T) {
	suite.Run(t, new(bizSuite))
}

func (s *bizSuite) Test_impl_GetByID() {
	type args struct {
		ctx  contextx.Contextx
		id   int64
		mock func()
	}
	tests := []struct {
		name     string
		args     args
		wantInfo *user.Profile
		wantErr  bool
	}{
		{
			name: "get by id then error",
			args: args{id: id1, mock: func() {
				s.mock.On("GetByID", mock.Anything, id1).Return(nil, errors.New("error")).Once()
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "get by id then not exists",
			args: args{id: id1, mock: func() {
				s.mock.On("GetByID", mock.Anything, id1).Return(nil, nil).Once()
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "get by id then user",
			args: args{id: id1, mock: func() {
				s.mock.On("GetByID", mock.Anything, id1).Return(info1, nil).Once()
			}},
			wantInfo: info1,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotInfo, err := s.biz.GetByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInfo, tt.wantInfo) {
				t.Errorf("GetByID() gotInfo = %v, want %v", gotInfo, tt.wantInfo)
			}

			s.TearDownTest()
		})
	}
}

func (s *bizSuite) Test_impl_GetByToken() {
	type args struct {
		ctx   contextx.Contextx
		token string
		mock  func()
	}
	tests := []struct {
		name     string
		args     args
		wantInfo *user.Profile
		wantErr  bool
	}{
		{
			name:     "missing token then error",
			args:     args{token: ""},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "get by id then error",
			args: args{token: token1, mock: func() {
				s.mock.On("GetByID", mock.Anything, id1).Return(nil, errors.New("error")).Once()
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "get by token then not exists",
			args: args{token: token1, mock: func() {
				s.mock.On("GetByID", mock.Anything, id1).Return(nil, nil).Once()

			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "get by token then user",
			args: args{token: token1, mock: func() {
				s.mock.On("GetByID", mock.Anything, id1).Return(info1, nil).Once()
			}},
			wantInfo: info1,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotInfo, err := s.biz.GetByToken(tt.args.ctx, tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInfo, tt.wantInfo) {
				t.Errorf("GetByToken() gotInfo = %v, want %v", gotInfo, tt.wantInfo)
			}

			s.TearDownTest()
		})
	}
}

func (s *bizSuite) Test_impl_Signup() {
	type args struct {
		ctx      contextx.Contextx
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
			name: "get by email then exists",
			args: args{email: email1, password: pass1, mock: func() {
				s.mock.On("GetByEmail", mock.Anything, email1).Return(info1, nil).Once()
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "register then error",
			args: args{email: email1, password: pass1, mock: func() {
				s.mock.On("GetByEmail", mock.Anything, email1).Return(nil, nil).Once()
				s.mock.On("Register", mock.Anything, mock.Anything).Return(nil, errors.New("error")).Once()
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "register then user",
			args: args{email: email1, password: pass1, mock: func() {
				s.mock.On("GetByEmail", mock.Anything, email1).Return(nil, nil).Once()
				s.mock.On("Register", mock.Anything, mock.Anything).Return(info1, nil).Once()
			}},
			wantInfo: info1,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotInfo, err := s.biz.Signup(tt.args.ctx, tt.args.email, tt.args.password)
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
		ctx      contextx.Contextx
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
			name: "get by email then not exists",
			args: args{email: email1, password: pass1, mock: func() {
				s.mock.On("GetByEmail", mock.Anything, email1).Return(nil, nil).Once()
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "login then incorrect password",
			args: args{email: email1, password: "error", mock: func() {
				s.mock.On("GetByEmail", mock.Anything, email1).Return(info1, nil).Once()
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "update token then error",
			args: args{email: email1, password: pass1, mock: func() {
				s.mock.On("GetByEmail", mock.Anything, email1).Return(info1, nil).Once()
				s.mock.On("UpdateToken", mock.Anything, info1).Return(nil, errors.New("error")).Once()
			}},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name: "login then user",
			args: args{email: email1, password: pass1, mock: func() {
				s.mock.On("GetByEmail", mock.Anything, email1).Return(info1, nil).Once()
				s.mock.On("UpdateToken", mock.Anything, info1).Return(info1, nil).Once()
			}},
			wantInfo: info1,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotInfo, err := s.biz.Login(tt.args.ctx, tt.args.email, tt.args.password)
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

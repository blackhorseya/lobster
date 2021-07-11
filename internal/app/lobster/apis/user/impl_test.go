package user

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/blackhorseya/lobster/internal/app/lobster/biz/user/mocks"
	"github.com/blackhorseya/lobster/internal/pkg/base/encrypt"
	"github.com/blackhorseya/lobster/internal/pkg/entity/er"
	"github.com/blackhorseya/lobster/internal/pkg/entity/user"
	"github.com/blackhorseya/lobster/internal/pkg/infra/transports/http/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

var (
	id1 = int64(200)

	token1 = "token"

	email1 = "email"

	pass1 = "password"

	salt1, _ = encrypt.HashAndSalt(pass1)

	info1 = &user.Profile{ID: id1, Token: token1, Password: salt1}
)

type handlerSuite struct {
	suite.Suite
	r       *gin.Engine
	mock    *mocks.IBiz
	handler IHandler
}

func (s *handlerSuite) SetupTest() {
	logger := zap.NewNop()

	gin.SetMode(gin.TestMode)
	s.r = gin.New()
	s.r.Use(middlewares.ContextMiddleware())
	s.r.Use(middlewares.ResponseMiddleware())

	s.mock = new(mocks.IBiz)
	if handler, err := CreateIHandler(logger, s.mock); err != nil {
		panic(err)
	} else {
		s.handler = handler
	}
}

func (s *handlerSuite) TearDownTest() {
	s.mock.AssertExpectations(s.T())
}

func TestHandlerSuite(t *testing.T) {
	suite.Run(t, new(handlerSuite))
}

func (s *handlerSuite) Test_impl_Signup() {
	s.r.POST("/api/v1/auth/signup", s.handler.Signup)

	type args struct {
		email    string
		password string
		mock     func()
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
	}{
		{
			name: "signup then error",
			args: args{email: email1, password: pass1, mock: func() {
				s.mock.On("Signup", mock.Anything, email1, pass1).Return(nil, er.ErrSignup).Once()
			}},
			wantCode: 500,
		},
		{
			name: "signup then user",
			args: args{email: email1, password: pass1, mock: func() {
				s.mock.On("Signup", mock.Anything, email1, pass1).Return(info1, nil).Once()
			}},
			wantCode: 201,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			uri := fmt.Sprintf("/api/v1/auth/signup")
			val := url.Values{}
			val.Add("email", tt.args.email)
			val.Add("password", tt.args.password)
			req := httptest.NewRequest(http.MethodPost, uri, strings.NewReader(val.Encode()))
			req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
			w := httptest.NewRecorder()
			s.r.ServeHTTP(w, req)

			got := w.Result()
			defer got.Body.Close()

			s.EqualValuesf(tt.wantCode, got.StatusCode, "Signup() code = %v, wantCode = %v", got.StatusCode, tt.wantCode)

			s.TearDownTest()
		})
	}
}

func (s *handlerSuite) Test_impl_Login() {
	s.r.POST("/api/v1/auth/login", s.handler.Login)

	type args struct {
		email    string
		password string
		mock     func()
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
	}{
		{
			name: "login then error",
			args: args{email: email1, password: pass1, mock: func() {
				s.mock.On("Login", mock.Anything, email1, pass1).Return(nil, er.ErrLogin).Once()
			}},
			wantCode: 500,
		},
		{
			name: "login then user",
			args: args{email: email1, password: pass1, mock: func() {
				s.mock.On("Login", mock.Anything, email1, pass1).Return(info1, nil).Once()
			}},
			wantCode: 201,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			uri := fmt.Sprintf("/api/v1/auth/login")
			val := url.Values{}
			val.Add("email", tt.args.email)
			val.Add("password", tt.args.password)
			req := httptest.NewRequest(http.MethodPost, uri, strings.NewReader(val.Encode()))
			req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
			w := httptest.NewRecorder()
			s.r.ServeHTTP(w, req)

			got := w.Result()
			defer got.Body.Close()

			s.EqualValuesf(tt.wantCode, got.StatusCode, "Login() code = %v, wantCode = %v", got.StatusCode, tt.wantCode)

			s.TearDownTest()
		})
	}
}

func (s *handlerSuite) Test_impl_GetByID() {
	s.r.GET("/api/v1/users/:id", s.handler.GetByID)

	type args struct {
		id   int64
		mock func()
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
	}{
		{
			name: "get by id then error",
			args: args{id: id1, mock: func() {
				s.mock.On("GetByID", mock.Anything, id1).Return(nil, er.ErrGetUserByID).Once()
			}},
			wantCode: 500,
		},
		{
			name: "get by id then user",
			args: args{id: id1, mock: func() {
				s.mock.On("GetByID", mock.Anything, id1).Return(info1, nil).Once()
			}},
			wantCode: 200,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			uri := fmt.Sprintf("/api/v1/users/%v", tt.args.id)
			req := httptest.NewRequest(http.MethodGet, uri, nil)
			w := httptest.NewRecorder()
			s.r.ServeHTTP(w, req)

			got := w.Result()
			defer got.Body.Close()

			s.EqualValuesf(tt.wantCode, got.StatusCode, "GetByID() code = %v, wantCode = %v", got.StatusCode, tt.wantCode)

			s.TearDownTest()
		})
	}
}

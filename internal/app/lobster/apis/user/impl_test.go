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
	id1 = int64(0)

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

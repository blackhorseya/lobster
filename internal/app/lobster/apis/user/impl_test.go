package user

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/blackhorseya/lobster/internal/app/lobster/biz/user/mocks"
	"github.com/blackhorseya/lobster/internal/pkg/entities/user"
	"github.com/blackhorseya/lobster/internal/pkg/transports/http/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var (
	uuid1 = "d76f4f51-f141-41ba-ba57-c4749319586b"

	time1 = int64(1610548520788105000)

	token1 = "b54c851b9d9e030f2afd6f6119b9c84e59f02590"

	email1 = "test@gmail.com"

	user1 = user.Profile{
		ID:          uuid1,
		AccessToken: token1,
		Email:       email1,
		SignupAt:    time1,
	}
)

type handlerSuite struct {
	suite.Suite
	r       *gin.Engine
	mock    *mocks.IBiz
	handler IHandler
}

func (s *handlerSuite) SetupTest() {
	gin.SetMode(gin.TestMode)

	s.r = gin.New()
	s.r.Use(middlewares.ContextMiddleware())

	s.mock = new(mocks.IBiz)
	handler, err := CreateIHandler(s.mock)
	if err != nil {
		panic(err)
		return
	}

	s.handler = handler
}

func (s *handlerSuite) TearDownTest() {
	s.mock.AssertExpectations(s.T())
}

func TestHandlerSuite(t *testing.T) {
	suite.Run(t, new(handlerSuite))
}

func (s *handlerSuite) Test_impl_Signup() {
	s.r.POST("/api/v1/users/signup", s.handler.Signup)

	type args struct {
		email string
		token string
		mock  func()
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
		wantBody *user.Profile
	}{
		{
			name: "profile then 500 error",
			args: args{email: email1, token: token1, mock: func() {
				s.mock.On("Signup", mock.Anything, email1, token1).Return(nil, errors.New("error")).Once()
			}},
			wantCode: 500,
			wantBody: nil,
		},
		{
			name: "profile then 201",
			args: args{email: email1, token: token1, mock: func() {
				s.mock.On("Signup", mock.Anything, email1, token1).Return(&user1, nil).Once()
			}},
			wantCode: 201,
			wantBody: nil,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			uri := fmt.Sprintf("/api/v1/users/signup")
			data, _ := json.Marshal(&user.Profile{Email: tt.args.email, AccessToken: tt.args.token})
			req := httptest.NewRequest(http.MethodPost, uri, bytes.NewBuffer(data))
			w := httptest.NewRecorder()
			s.r.ServeHTTP(w, req)

			got := w.Result()
			defer got.Body.Close()

			body, _ := ioutil.ReadAll(got.Body)
			var gotBody *user.Profile
			_ = json.Unmarshal(body, &gotBody)

			s.EqualValuesf(tt.wantCode, got.StatusCode, "Signup() code = %v, wantCode = %v", got.StatusCode, tt.wantCode)
			if tt.wantBody != nil && !reflect.DeepEqual(gotBody, tt.wantBody) {
				s.T().Errorf("Signup() got = %v, wantBody = %v", gotBody, tt.wantBody)
			}

			s.TearDownTest()
		})
	}
}

func (s *handlerSuite) Test_impl_Login() {
	s.r.POST("/api/v1/users/login", s.handler.Login)

	type args struct {
		email string
		token string
		mock  func()
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
		wantBody *user.Profile
	}{
		{
			name: "profile then 500 error",
			args: args{email: email1, token: token1, mock: func() {
				s.mock.On("Login", mock.Anything, email1, token1).Return(nil, errors.New("error")).Once()
			}},
			wantCode: 500,
			wantBody: nil,
		},
		{
			name: "profile then 201",
			args: args{email: email1, token: token1, mock: func() {
				s.mock.On("Login", mock.Anything, email1, token1).Return(&user1, nil).Once()
			}},
			wantCode: 201,
			wantBody: nil,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			uri := fmt.Sprintf("/api/v1/users/login")
			data, _ := json.Marshal(&user.Profile{Email: tt.args.email, AccessToken: tt.args.token})
			req := httptest.NewRequest(http.MethodPost, uri, bytes.NewBuffer(data))
			w := httptest.NewRecorder()
			s.r.ServeHTTP(w, req)

			got := w.Result()
			defer got.Body.Close()

			body, _ := ioutil.ReadAll(got.Body)
			var gotBody *user.Profile
			_ = json.Unmarshal(body, &gotBody)

			s.EqualValuesf(tt.wantCode, got.StatusCode, "Signup() code = %v, wantCode = %v", got.StatusCode, tt.wantCode)
			if tt.wantBody != nil && !reflect.DeepEqual(gotBody, tt.wantBody) {
				s.T().Errorf("Signup() got = %v, wantBody = %v", gotBody, tt.wantBody)
			}

			s.TearDownTest()
		})
	}
}

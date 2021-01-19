package health

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/blackhorseya/lobster/internal/app/biz/health/mocks"
	"github.com/blackhorseya/lobster/internal/pkg/transports/http/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
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
	s.r.Use(middlewares.LoggerMiddleware())

	s.mock = new(mocks.IBiz)
	if handler, err := CreateIHandler(s.mock); err != nil {
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

func (s *handlerSuite) Test_impl_Readiness() {
	s.r.GET("/api/readiness", s.handler.Readiness)

	type args struct {
		mock func()
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
	}{
		{
			name: "readiness then 500 error",
			args: args{mock: func() {
				s.mock.On("Readiness", mock.Anything).Return(errors.New("error")).Once()
			}},
			wantCode: 500,
		},
		{
			name: "readiness then 200 nil",
			args: args{mock: func() {
				s.mock.On("Readiness", mock.Anything).Return(nil).Once()
			}},
			wantCode: 200,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			uri := fmt.Sprintf("/api/readiness")
			req := httptest.NewRequest(http.MethodGet, uri, nil)
			w := httptest.NewRecorder()
			s.r.ServeHTTP(w, req)

			got := w.Result()
			defer func() {
				_ = got.Body.Close()
			}()

			s.EqualValuesf(tt.wantCode, got.StatusCode, "Readiness() code = %v, wantCode = %v", got.StatusCode, tt.wantCode)

			s.TearDownTest()
		})
	}
}

func (s *handlerSuite) Test_impl_Liveness() {
	s.r.GET("/api/liveness", s.handler.Liveness)

	type args struct {
		mock func()
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
	}{
		{
			name: "readiness then 500 error",
			args: args{mock: func() {
				s.mock.On("Liveness", mock.Anything).Return(errors.New("error")).Once()
			}},
			wantCode: 500,
		},
		{
			name: "readiness then 200 nil",
			args: args{mock: func() {
				s.mock.On("Liveness", mock.Anything).Return(nil).Once()
			}},
			wantCode: 200,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			uri := fmt.Sprintf("/api/liveness")
			req := httptest.NewRequest(http.MethodGet, uri, nil)
			w := httptest.NewRecorder()
			s.r.ServeHTTP(w, req)

			got := w.Result()
			defer func() {
				_ = got.Body.Close()
			}()

			s.EqualValuesf(tt.wantCode, got.StatusCode, "Liveness() code = %v, wantCode = %v", got.StatusCode, tt.wantCode)

			s.TearDownTest()
		})
	}
}

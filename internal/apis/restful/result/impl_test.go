package result

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

	"github.com/blackhorseya/lobster/internal/biz/result/mocks"
	"github.com/blackhorseya/lobster/internal/pkg/pb"
	"github.com/blackhorseya/lobster/internal/pkg/transports/http/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var (
	krID   = "d76f4f51-f141-41ba-ba57-c4749319586b"
	goalID = "0829ee06-1f04-43d9-8565-812e1826f805"

	time1 = int64(1611059529208050000)

	kr1 = &pb.Result{
		ID:       krID,
		GoalID:   goalID,
		Title:    "kr1",
		CreateAt: time1,
	}

	created1 = &pb.Result{
		Title:  "created kr1",
		GoalID: goalID,
	}

	updated1 = &pb.Result{
		ID:       krID,
		GoalID:   goalID,
		Title:    "updated kr1",
		CreateAt: time1,
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
	s.r.Use(middlewares.LoggerMiddleware())

	s.mock = new(mocks.IBiz)
	handler, err := CreateIHandler(s.mock)
	if err != nil {
		panic(err)
	}

	s.handler = handler
}

func (s *handlerSuite) TearDownTest() {
	s.mock.AssertExpectations(s.T())
}

func TestHandlerSuite(t *testing.T) {
	suite.Run(t, new(handlerSuite))
}

func (s *handlerSuite) Test_impl_GetByID() {
	s.r.GET("/api/v1/krs/:id", s.handler.GetByID)

	type args struct {
		id   string
		mock func()
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
		wantBody *pb.Result
	}{
		{
			name:     "id then 400 error",
			args:     args{id: "id"},
			wantCode: 400,
			wantBody: nil,
		},
		{
			name: "uuid then 500 error",
			args: args{id: krID, mock: func() {
				s.mock.On("GetByID", mock.Anything, krID).Return(nil, errors.New("error")).Once()
			}},
			wantCode: 500,
			wantBody: nil,
		},
		{
			name: "uuid then not found error",
			args: args{id: krID, mock: func() {
				s.mock.On("GetByID", mock.Anything, krID).Return(nil, nil).Once()
			}},
			wantCode: 404,
			wantBody: nil,
		},
		{
			name: "uuid then kr nil",
			args: args{id: krID, mock: func() {
				s.mock.On("GetByID", mock.Anything, krID).Return(kr1, nil).Once()
			}},
			wantCode: 200,
			wantBody: kr1,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			uri := fmt.Sprintf("/api/v1/krs/%v", tt.args.id)
			req := httptest.NewRequest(http.MethodGet, uri, nil)
			w := httptest.NewRecorder()
			s.r.ServeHTTP(w, req)

			got := w.Result()
			defer got.Body.Close()

			var gotBody *pb.Result
			body, _ := ioutil.ReadAll(got.Body)
			err := json.Unmarshal(body, &gotBody)
			if err != nil {
				s.Errorf(err, "unmarshal response body is failure")
			}

			s.EqualValuesf(tt.wantCode, got.StatusCode, "GetByID() code = %v, wantCode = %v", got.StatusCode, tt.wantCode)
			if tt.wantBody != nil && !reflect.DeepEqual(gotBody, tt.wantBody) {
				s.T().Errorf("GetByID() got = %v, wantBody = %v", gotBody, tt.wantBody)
			}

			s.TearDownTest()
		})
	}
}

func (s *handlerSuite) Test_impl_List() {
	s.r.GET("/api/v1/krs", s.handler.List)

	type args struct {
		page string
		size string
		mock func()
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
		wantBody []*pb.Result
	}{
		{
			name:     "a 10 then 400 error",
			args:     args{page: "a", size: "10"},
			wantCode: 400,
			wantBody: nil,
		},
		{
			name:     "10 b then 400 error",
			args:     args{page: "10", size: "b"},
			wantCode: 400,
			wantBody: nil,
		},
		{
			name: "1 1 then 500 error",
			args: args{page: "1", size: "1", mock: func() {
				s.mock.On("List", mock.Anything, 1, 1).Return(nil, errors.New("error")).Once()
			}},
			wantCode: 500,
			wantBody: nil,
		},
		{
			name: "1 1 then 404 error",
			args: args{page: "1", size: "1", mock: func() {
				s.mock.On("List", mock.Anything, 1, 1).Return(nil, nil).Once()
			}},
			wantCode: 404,
			wantBody: nil,
		},
		{
			name: "1 1 then 200 nil",
			args: args{page: "1", size: "1", mock: func() {
				s.mock.On("List", mock.Anything, 1, 1).Return([]*pb.Result{kr1}, nil).Once()
			}},
			wantCode: 200,
			wantBody: []*pb.Result{kr1},
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			uri := fmt.Sprintf("/api/v1/krs?page=%v&size=%v", tt.args.page, tt.args.size)
			req := httptest.NewRequest(http.MethodGet, uri, nil)
			w := httptest.NewRecorder()
			s.r.ServeHTTP(w, req)

			got := w.Result()
			defer got.Body.Close()

			var gotBody []*pb.Result
			body, _ := ioutil.ReadAll(got.Body)
			err := json.Unmarshal(body, &gotBody)
			if err != nil {
				s.Errorf(err, "unmarshal response body is failure")
			}

			s.EqualValuesf(tt.wantCode, got.StatusCode, "List() code = %v, wantCode = %v", got.StatusCode, tt.wantCode)
			if tt.wantBody != nil && !reflect.DeepEqual(gotBody, tt.wantBody) {
				s.T().Errorf("List() got = %v, wantBody = %v", gotBody, tt.wantBody)
			}

			s.TearDownTest()
		})
	}
}

func (s *handlerSuite) Test_impl_Create() {
	s.r.POST("/api/v1/krs", s.handler.Create)

	type args struct {
		created *pb.Result
		mock    func()
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
		wantBody *pb.Result
	}{
		{
			name:     "missing title then 400 error",
			args:     args{created: &pb.Result{Title: "", GoalID: goalID}},
			wantCode: 400,
			wantBody: nil,
		},
		{
			name:     "missing parent goal then 400 error",
			args:     args{created: &pb.Result{Title: "title", GoalID: ""}},
			wantCode: 400,
			wantBody: nil,
		},
		{
			name: "kr then 500 error",
			args: args{created: created1, mock: func() {
				s.mock.On("LinkToGoal", mock.Anything, created1).Return(nil, errors.New("error")).Once()
			}},
			wantCode: 500,
			wantBody: nil,
		},
		{
			name: "kr then 201 nil",
			args: args{created: created1, mock: func() {
				s.mock.On("LinkToGoal", mock.Anything, created1).Return(kr1, nil).Once()
			}},
			wantCode: 201,
			wantBody: kr1,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			uri := fmt.Sprintf("/api/v1/krs")
			data, _ := json.Marshal(tt.args.created)
			req := httptest.NewRequest(http.MethodPost, uri, bytes.NewBuffer(data))
			w := httptest.NewRecorder()
			s.r.ServeHTTP(w, req)

			got := w.Result()
			defer got.Body.Close()

			var gotBody *pb.Result
			body, _ := ioutil.ReadAll(got.Body)
			err := json.Unmarshal(body, &gotBody)
			if err != nil {
				s.Errorf(err, "unmarshal response body is failure")
			}

			s.EqualValuesf(tt.wantCode, got.StatusCode, "Create() code = %v, wantCode = %v", got.StatusCode, tt.wantCode)
			if tt.wantBody != nil && !reflect.DeepEqual(gotBody, tt.wantBody) {
				s.T().Errorf("Create() got = %v, wantBody = %v", gotBody, tt.wantBody)
			}

			s.TearDownTest()
		})
	}
}

func (s *handlerSuite) Test_impl_Update() {
	s.r.PUT("/api/v1/krs/:id", s.handler.Update)

	type args struct {
		id      string
		updated *pb.Result
		mock    func()
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
		wantBody *pb.Result
	}{
		{
			name:     "id then 400 error",
			args:     args{id: "id"},
			wantCode: 400,
			wantBody: nil,
		},
		{
			name:     "missing title then 400 error",
			args:     args{id: krID, updated: &pb.Result{Title: "", GoalID: goalID}},
			wantCode: 400,
			wantBody: nil,
		},
		{
			name:     "missing parent goal then 400 error",
			args:     args{id: krID, updated: &pb.Result{Title: "title", GoalID: "id"}},
			wantCode: 400,
			wantBody: nil,
		},
		{
			name: "update then 500 error",
			args: args{id: krID, updated: kr1, mock: func() {
				s.mock.On("Update", mock.Anything, kr1).Return(nil, errors.New("error")).Once()
			}},
			wantCode: 500,
			wantBody: nil,
		},
		{
			name: "update then 200 nil",
			args: args{id: krID, updated: kr1, mock: func() {
				s.mock.On("Update", mock.Anything, kr1).Return(updated1, nil).Once()
			}},
			wantCode: 200,
			wantBody: updated1,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			uri := fmt.Sprintf("/api/v1/krs/%v", tt.args.id)
			body, _ := json.Marshal(tt.args.updated)
			req := httptest.NewRequest(http.MethodPut, uri, bytes.NewBuffer(body))
			w := httptest.NewRecorder()
			s.r.ServeHTTP(w, req)

			got := w.Result()
			defer got.Body.Close()

			var gotBody *pb.Result
			body, _ = ioutil.ReadAll(got.Body)
			err := json.Unmarshal(body, &gotBody)
			if err != nil {
				s.Errorf(err, "unmarshal response body is failure")
			}

			s.EqualValuesf(tt.wantCode, got.StatusCode, "Update() code = %v, wantCode = %v", got.StatusCode, tt.wantCode)
			if tt.wantBody != nil && !reflect.DeepEqual(gotBody, tt.wantBody) {
				s.T().Errorf("Update() got = %v, wantBody = %v", gotBody, tt.wantBody)
			}

			s.TearDownTest()
		})
	}
}

func (s *handlerSuite) Test_impl_Delete() {
	s.r.DELETE("/api/v1/krs/:id", s.handler.Delete)

	type args struct {
		id   string
		mock func()
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
	}{
		{
			name:     "id then 400 error",
			args:     args{id: "id"},
			wantCode: 400,
		},
		{
			name: "uuid then 500 error",
			args: args{id: krID, mock: func() {
				s.mock.On("Delete", mock.Anything, krID).Return(errors.New("error")).Once()
			}},
			wantCode: 500,
		},
		{
			name: "uuid then 204 nil",
			args: args{id: krID, mock: func() {
				s.mock.On("Delete", mock.Anything, krID).Return(nil).Once()
			}},
			wantCode: 204,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			uri := fmt.Sprintf("/api/v1/krs/%v", tt.args.id)
			req := httptest.NewRequest(http.MethodDelete, uri, nil)
			w := httptest.NewRecorder()
			s.r.ServeHTTP(w, req)

			got := w.Result()
			defer got.Body.Close()

			s.EqualValuesf(tt.wantCode, got.StatusCode, "Delete() code = %v, wantCode = %v", got.StatusCode, tt.wantCode)

			s.TearDownTest()
		})
	}
}

func (s *handlerSuite) Test_impl_GetByGoalID() {
	s.r.GET("/api/v1/goals/:id/results", s.handler.GetByGoalID)

	type args struct {
		id   string
		mock func()
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
		wantBody []*pb.Result
	}{
		{
			name:     "id then 400 error",
			args:     args{id: "id"},
			wantCode: 400,
			wantBody: nil,
		},
		{
			name: "uuid then 500 error",
			args: args{id: goalID, mock: func() {
				s.mock.On("GetByGoalID", mock.Anything, goalID).Return(nil, errors.New("error")).Once()
			}},
			wantCode: 500,
			wantBody: nil,
		},
		{
			name: "uuid then 404 error",
			args: args{id: goalID, mock: func() {
				s.mock.On("GetByGoalID", mock.Anything, goalID).Return(nil, nil).Once()
			}},
			wantCode: 404,
			wantBody: nil,
		},
		{
			name: "uuid then 200 results",
			args: args{id: goalID, mock: func() {
				s.mock.On("GetByGoalID", mock.Anything, goalID).Return([]*pb.Result{kr1}, nil).Once()
			}},
			wantCode: 200,
			wantBody: []*pb.Result{kr1},
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			uri := fmt.Sprintf("/api/v1/goals/%v/results", tt.args.id)
			req := httptest.NewRequest(http.MethodGet, uri, nil)
			w := httptest.NewRecorder()
			s.r.ServeHTTP(w, req)

			got := w.Result()
			defer got.Body.Close()

			var gotBody []*pb.Result
			body, _ := ioutil.ReadAll(got.Body)
			err := json.Unmarshal(body, &gotBody)
			if err != nil {
				s.Errorf(err, "unmarshal response body is failure")
			}

			s.EqualValuesf(tt.wantCode, got.StatusCode, "QueryByGoalID() code = %v, wantCode = %v", got.StatusCode, tt.wantCode)
			if tt.wantBody != nil && !reflect.DeepEqual(gotBody, tt.wantBody) {
				s.T().Errorf("QueryByGoalID() got = %v, wantBody = %v", gotBody, tt.wantBody)
			}

			s.TearDownTest()
		})
	}
}

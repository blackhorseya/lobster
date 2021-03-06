package task

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

	"github.com/blackhorseya/lobster/internal/app/lobster/biz/task/mocks"
	"github.com/blackhorseya/lobster/internal/pkg/entity/er"
	"github.com/blackhorseya/lobster/internal/pkg/entity/response"
	"github.com/blackhorseya/lobster/internal/pkg/entity/todo"
	"github.com/blackhorseya/lobster/internal/pkg/infra/transports/http/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

var (
	uuid1 = int64(1)

	time1 = int64(1610548520788105000)

	task1 = &todo.Task{
		ID:        uuid1,
		Title:     "task1",
		CreatedAt: time1,
	}

	created1 = &todo.Task{
		Title: "create task1",
	}

	updated1 = &todo.Task{
		ID:        uuid1,
		Title:     "updated task1",
		CreatedAt: time1,
	}

	updated2 = &todo.Task{
		ID:        uuid1,
		Status:    todo.Status_INPROGRESS,
		Title:     "task1",
		CreatedAt: time1,
	}
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

func (s *handlerSuite) Test_impl_GetByID() {
	s.r.GET("/api/v1/tasks/:id", s.handler.GetByID)

	type args struct {
		id   int64
		mock func()
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
		wantBody *response.Response
	}{
		{
			name: "uuid then 500 error",
			args: args{id: uuid1, mock: func() {
				s.mock.On("GetByID", mock.Anything, uuid1).Return(nil, er.ErrGetTaskByID).Once()
			}},
			wantCode: 500,
			wantBody: nil,
		},
		{
			name: "uuid then 200 task",
			args: args{id: uuid1, mock: func() {
				s.mock.On("GetByID", mock.Anything, uuid1).Return(task1, nil).Once()
			}},
			wantCode: 200,
			wantBody: &response.Response{Code: 200, Msg: "ok", Data: task1},
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			uri := fmt.Sprintf("/api/v1/tasks/%v", tt.args.id)
			req := httptest.NewRequest(http.MethodGet, uri, nil)
			w := httptest.NewRecorder()
			s.r.ServeHTTP(w, req)

			got := w.Result()
			defer func() {
				_ = got.Body.Close()
			}()

			body, _ := ioutil.ReadAll(got.Body)
			var gotBody *response.Response
			if err := json.Unmarshal(body, &gotBody); err != nil {
				s.Errorf(err, "unmarshal response body is failure")
			}

			s.EqualValuesf(tt.wantCode, got.StatusCode, "GetByID() code = %v, wantCode = %v", got.StatusCode, tt.wantCode)
			if tt.wantBody != nil && !reflect.DeepEqual(gotBody, tt.wantBody) {
				s.Errorf(fmt.Errorf("GetByID() got = %v, wantBody = %v", gotBody, tt.wantBody), "GetByID")
			}

			s.TearDownTest()
		})
	}
}

func (s *handlerSuite) Test_impl_List() {
	s.r.GET("/api/v1/tasks", s.handler.List)

	type args struct {
		page string
		size string
		mock func()
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
		wantBody []*todo.Task
	}{
		{
			name:     "a 10 then 400 error",
			args:     args{page: "a", size: "10"},
			wantCode: 400,
			wantBody: nil,
		},
		{
			name:     "1 b then 400 error",
			args:     args{page: "1", size: "b"},
			wantCode: 400,
			wantBody: nil,
		},
		{
			name: "1 1 then 500 error",
			args: args{page: "1", size: "1", mock: func() {
				s.mock.On("List", mock.Anything, 1, 1).Return(nil, er.ErrListTasks).Once()
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
			name: "1 1 then 200",
			args: args{page: "1", size: "1", mock: func() {
				s.mock.On("List", mock.Anything, 1, 1).Return([]*todo.Task{task1}, nil).Once()
			}},
			wantCode: 200,
			wantBody: []*todo.Task{task1},
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			uri := fmt.Sprintf("/api/v1/tasks?page=%v&size=%v", tt.args.page, tt.args.size)
			req := httptest.NewRequest(http.MethodGet, uri, nil)
			w := httptest.NewRecorder()
			s.r.ServeHTTP(w, req)

			got := w.Result()
			defer func() {
				_ = got.Body.Close()
			}()

			body, _ := ioutil.ReadAll(got.Body)
			var gotBody []*todo.Task
			if err := json.Unmarshal(body, &gotBody); err != nil {
				s.Errorf(err, "unmarshal response body is failure")
			}

			s.EqualValuesf(tt.wantCode, got.StatusCode, "List() code = %v, wantCode = %v", got.StatusCode, tt.wantCode)
			s.ElementsMatchf(tt.wantBody, gotBody, "List() body = %v, wantBody = %v", gotBody, tt.wantBody)

			s.TearDownTest()
		})
	}
}

func (s *handlerSuite) Test_impl_Create() {
	s.r.POST("/api/v1/tasks", s.handler.Create)

	type args struct {
		title string
		mock  func()
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
		wantBody *todo.Task
	}{
		{
			name:     "empty title then 400 error",
			args:     args{title: ""},
			wantCode: 400,
			wantBody: nil,
		},
		{
			name: "task then 500 error",
			args: args{title: task1.Title, mock: func() {
				s.mock.On("Create", mock.Anything, task1.Title).Return(nil, errors.New("error")).Once()
			}},
			wantCode: 500,
			wantBody: nil,
		},
		{
			name: "task then 201 nil",
			args: args{title: task1.Title, mock: func() {
				s.mock.On("Create", mock.Anything, task1.Title).Return(task1, nil).Once()
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

			uri := fmt.Sprintf("/api/v1/tasks")
			data, _ := json.Marshal(&todo.Task{Title: tt.args.title})
			req := httptest.NewRequest(http.MethodPost, uri, bytes.NewBuffer(data))
			w := httptest.NewRecorder()
			s.r.ServeHTTP(w, req)

			got := w.Result()
			defer func() {
				_ = got.Body.Close()
			}()

			body, _ := ioutil.ReadAll(got.Body)
			var gotBody *todo.Task
			if err := json.Unmarshal(body, &gotBody); err != nil {
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

func (s *handlerSuite) Test_impl_Delete() {
	s.r.DELETE("/api/v1/tasks/:id", s.handler.Delete)

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
			name: "uuid then 500 error",
			args: args{id: uuid1, mock: func() {
				s.mock.On("Delete", mock.Anything, uuid1).Return(errors.New("error")).Once()
			}},
			wantCode: 500,
		},
		{
			name: "uuid then 204 nil",
			args: args{id: uuid1, mock: func() {
				s.mock.On("Delete", mock.Anything, uuid1).Return(nil).Once()
			}},
			wantCode: 204,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			uri := fmt.Sprintf("/api/v1/tasks/%v", tt.args.id)
			req := httptest.NewRequest(http.MethodDelete, uri, nil)
			w := httptest.NewRecorder()
			s.r.ServeHTTP(w, req)

			got := w.Result()
			defer func() {
				_ = got.Body.Close()
			}()

			s.EqualValuesf(tt.wantCode, got.StatusCode, "Delete() code = %v, wantCode = %v", got.StatusCode, tt.wantCode)

			s.TearDownTest()
		})
	}
}

func (s *handlerSuite) Test_impl_UpdateStatus() {
	s.r.PATCH("/api/v1/tasks/:id/status", s.handler.UpdateStatus)

	type args struct {
		id     int64
		status todo.Status
		mock   func()
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
		wantBody *response.Response
	}{
		{
			name: "uuid then 500 error",
			args: args{id: uuid1, status: todo.Status_INPROGRESS, mock: func() {
				s.mock.On("UpdateStatus", mock.Anything, uuid1, todo.Status_INPROGRESS).Return(nil, errors.New("error")).Once()
			}},
			wantCode: 500,
		},
		{
			name: "uuid status then 200 task",
			args: args{id: uuid1, status: todo.Status_INPROGRESS, mock: func() {
				s.mock.On("UpdateStatus", mock.Anything, uuid1, todo.Status_INPROGRESS).Return(updated2, nil).Once()
			}},
			wantCode: 200,
			wantBody: response.OK.WithData(updated2),
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			uri := fmt.Sprintf("/api/v1/tasks/%v/status", tt.args.id)
			data, _ := json.Marshal(&todo.Task{Status: tt.args.status})
			req := httptest.NewRequest(http.MethodPatch, uri, bytes.NewBuffer(data))
			w := httptest.NewRecorder()
			s.r.ServeHTTP(w, req)

			got := w.Result()
			defer got.Body.Close()

			body, _ := ioutil.ReadAll(got.Body)
			var gotBody *todo.Task
			_ = json.Unmarshal(body, &gotBody)

			s.EqualValuesf(tt.wantCode, got.StatusCode, "Delete() code = %v, wantCode = %v", got.StatusCode, tt.wantCode)
			if tt.wantBody != nil && !reflect.DeepEqual(gotBody, tt.wantBody) {
				s.Errorf(fmt.Errorf("Update() got = %v, wantBody = %v", gotBody, tt.wantBody), "Update")
			}

			s.TearDownTest()
		})
	}
}

func (s *handlerSuite) Test_impl_ModifyTitle() {
	s.r.PATCH("/api/v1/tasks/:id/title", s.handler.ModifyTitle)

	type args struct {
		id    int64
		title string
		mock  func()
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
		wantBody *response.Response
	}{
		{
			name:     "uuid missing title then error",
			args:     args{id: uuid1, title: ""},
			wantCode: 400,
			wantBody: nil,
		},
		{
			name: "uuid title then 500 error",
			args: args{id: uuid1, title: "title", mock: func() {
				s.mock.On("ModifyTitle", mock.Anything, uuid1, "title").Return(nil, errors.New("error")).Once()
			}},
			wantCode: 500,
			wantBody: nil,
		},
		{
			name: "uuid title then 200 nil",
			args: args{id: uuid1, title: "title", mock: func() {
				s.mock.On("ModifyTitle", mock.Anything, uuid1, "title").Return(task1, nil).Once()
			}},
			wantCode: 200,
			wantBody: response.OK.WithData(task1),
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			uri := fmt.Sprintf("/api/v1/tasks/%v/title", tt.args.id)
			data, _ := json.Marshal(&todo.Task{ID: tt.args.id, Title: tt.args.title})
			req := httptest.NewRequest(http.MethodPatch, uri, bytes.NewBuffer(data))
			w := httptest.NewRecorder()
			s.r.ServeHTTP(w, req)

			got := w.Result()
			defer got.Body.Close()

			body, _ := ioutil.ReadAll(got.Body)
			var gotBody *todo.Task
			_ = json.Unmarshal(body, &gotBody)

			s.EqualValuesf(tt.wantCode, got.StatusCode, "ModifyTitle() code = %v, wantCode = %v", got.StatusCode, tt.wantCode)
			if tt.wantBody != nil && !reflect.DeepEqual(gotBody, tt.wantBody) {
				s.Errorf(fmt.Errorf("ModifyTitle() got = %v, wantBody = %v", gotBody, tt.wantBody), "ModifyTitle")
			}

			s.TearDownTest()
		})
	}
}

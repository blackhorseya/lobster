package objective

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

	"github.com/blackhorseya/lobster/internal/biz/goal/mocks"
	"github.com/blackhorseya/lobster/internal/pkg/entities/biz/okr"
	"github.com/blackhorseya/lobster/internal/pkg/transports/http/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var (
	uuid1 = "d76f4f51-f141-41ba-ba57-c4749319586b"

	time1 = int64(1610548520788105000)

	obj1 = &okr.Objective{
		ID:       uuid1,
		Title:    "obj1",
		CreateAt: time1,
	}

	created1 = &okr.Objective{Title: "created obj1"}

	updated1 = &okr.Objective{Title: "updated obj1"}
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

func (s *handlerSuite) Test_impl_GetByID() {
	s.r.GET("/api/v1/objectives/:id", s.handler.GetByID)

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
			name: "uuid then 200 error",
			args: args{id: uuid1, mock: func() {
				s.mock.On("GetByID", mock.Anything, uuid1).Return(nil, errors.New("error")).Once()
			}},
			wantCode: 200,
		},
		{
			name: "uuid then 200 obj",
			args: args{id: uuid1, mock: func() {
				s.mock.On("GetByID", mock.Anything, uuid1).Return(obj1, nil).Once()
			}},
			wantCode: 200,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			uri := fmt.Sprintf("/api/v1/objectives/%v", tt.args.id)
			req := httptest.NewRequest(http.MethodGet, uri, nil)
			w := httptest.NewRecorder()
			s.r.ServeHTTP(w, req)

			got := w.Result()
			defer func() {
				_ = got.Body.Close()
			}()

			s.EqualValuesf(tt.wantCode, got.StatusCode, "GetByID() code = %v, wantCode = %v", got.StatusCode, tt.wantCode)

			s.TearDownTest()
		})
	}
}

func (s *handlerSuite) Test_impl_List() {
	s.r.GET("/api/v1/objectives", s.handler.List)

	type args struct {
		page string
		size string
		mock func()
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
		wantBody []*okr.Objective
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
			name: "1 1 then 200 error",
			args: args{page: "1", size: "1", mock: func() {
				s.mock.On("List", mock.Anything, 1, 1).Return(nil, errors.New("error")).Once()
			}},
			wantCode: 200,
			wantBody: nil,
		},
		{
			name: "10 10 then 404 error",
			args: args{page: "10", size: "10", mock: func() {
				s.mock.On("List", mock.Anything, 10, 10).Return(nil, nil).Once()
			}},
			wantCode: 404,
			wantBody: nil,
		},
		{
			name: "1 1 then 200 error",
			args: args{page: "1", size: "1", mock: func() {
				s.mock.On("List", mock.Anything, 1, 1).Return(
					[]*okr.Objective{obj1}, nil).Once()
			}},
			wantCode: 200,
			wantBody: []*okr.Objective{obj1},
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			uri := fmt.Sprintf("/api/v1/objectives?page=%v&size=%v", tt.args.page, tt.args.size)
			req := httptest.NewRequest(http.MethodGet, uri, nil)
			w := httptest.NewRecorder()
			s.r.ServeHTTP(w, req)

			got := w.Result()
			defer func() {
				_ = got.Body.Close()
			}()

			body, _ := ioutil.ReadAll(got.Body)
			var gotBody []*okr.Objective
			if err := json.Unmarshal(body, &gotBody); err != nil {
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
	s.r.POST("/api/v1/objectives", s.handler.Create)

	type args struct {
		created *okr.Objective
		mock    func()
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
		wantBody *okr.Objective
	}{
		{
			name:     "empty title then 400 error",
			args:     args{created: &okr.Objective{Title: ""}},
			wantCode: 400,
			wantBody: nil,
		},
		{
			name: "created then 200 error",
			args: args{created: created1, mock: func() {
				s.mock.On("Create", mock.Anything, created1).Return(nil, errors.New("error")).Once()
			}},
			wantCode: 200,
			wantBody: nil,
		},
		{
			name: "created then 201 obj1",
			args: args{created: created1, mock: func() {
				s.mock.On("Create", mock.Anything, created1).Return(obj1, nil).Once()
			}},
			wantCode: 201,
			wantBody: obj1,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			uri := fmt.Sprintf("/api/v1/objectives")
			data, _ := json.Marshal(tt.args.created)
			req := httptest.NewRequest(http.MethodPost, uri, bytes.NewBuffer(data))
			w := httptest.NewRecorder()
			s.r.ServeHTTP(w, req)

			got := w.Result()
			defer func() {
				_ = got.Body.Close()
			}()

			body, _ := ioutil.ReadAll(got.Body)
			var gotBody *okr.Objective
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

func (s *handlerSuite) Test_impl_Update() {
	s.r.PUT("/api/v1/objectives/:id", s.handler.Update)

	type args struct {
		id      string
		updated *okr.Objective
		mock    func()
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
		wantBody *okr.Objective
	}{
		{
			name:     "id then 400 error",
			args:     args{id: "id", updated: &okr.Objective{Title: "updated obj1"}},
			wantCode: 400,
			wantBody: nil,
		},
		{
			name:     "uuid empty title then 400 error",
			args:     args{id: uuid1, updated: &okr.Objective{Title: ""}},
			wantCode: 400,
			wantBody: nil,
		},
		{
			name: "uuid updated then 200 error",
			args: args{id: uuid1, updated: updated1, mock: func() {
				s.mock.On("Update", mock.Anything, updated1).Return(nil, errors.New("error")).Once()
			}},
			wantCode: 200,
			wantBody: nil,
		},
		{
			name: "uuid updated then 200 nil",
			args: args{id: uuid1, updated: updated1, mock: func() {
				s.mock.On("Update", mock.Anything, updated1).Return(obj1, nil).Once()
			}},
			wantCode: 200,
			wantBody: obj1,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			uri := fmt.Sprintf("/api/v1/objectives/%v", tt.args.id)
			data, _ := json.Marshal(tt.args.updated)
			req := httptest.NewRequest(http.MethodPut, uri, bytes.NewBuffer(data))
			w := httptest.NewRecorder()
			s.r.ServeHTTP(w, req)

			got := w.Result()
			defer func() {
				_ = got.Body.Close()
			}()

			body, _ := ioutil.ReadAll(got.Body)
			var gotBody *okr.Objective
			if err := json.Unmarshal(body, &gotBody); err != nil {
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
	s.r.DELETE("/api/v1/objectives/:id", s.handler.Delete)

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
			name: "uuid then 200 error",
			args: args{id: uuid1, mock: func() {
				s.mock.On("Delete", mock.Anything, uuid1).Return(errors.New("error")).Once()
			}},
			wantCode: 200,
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

			uri := fmt.Sprintf("/api/v1/objectives/%v", tt.args.id)
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

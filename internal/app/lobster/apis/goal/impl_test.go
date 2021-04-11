package goal

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

	"github.com/blackhorseya/lobster/internal/app/lobster/biz/goal/mocks"
	"github.com/blackhorseya/lobster/internal/pkg/entities/okr"
	"github.com/blackhorseya/lobster/internal/pkg/entities/response"
	"github.com/blackhorseya/lobster/internal/pkg/transports/http/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

var (
	uuid1 = "d76f4f51-f141-41ba-ba57-c4749319586b"

	time1 = int64(1610548520788105000)

	obj1 = &okr.Goal{
		ID:       uuid1,
		Title:    "obj1",
		CreateAt: time1,
	}

	created1 = &okr.Goal{Title: "created obj1"}

	updated1 = &okr.Goal{Title: "updated obj1"}
)

type handlerSuite struct {
	suite.Suite
	r       *gin.Engine
	mock    *mocks.IBiz
	handler IHandler
}

func (s *handlerSuite) SetupTest() {
	logger, _ := zap.NewDevelopment()

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
			name: "uuid then 500 error",
			args: args{id: uuid1, mock: func() {
				s.mock.On("GetByID", mock.Anything, uuid1).Return(nil, errors.New("error")).Once()
			}},
			wantCode: 500,
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
		wantBody *response.Response
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
				s.mock.On("List", mock.Anything, 1, 1).Return(nil, errors.New("error")).Once()
			}},
			wantCode: 500,
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
					[]*okr.Goal{obj1}, nil).Once()
			}},
			wantCode: 200,
			wantBody: response.OK.WithData([]*okr.Goal{obj1}),
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
			var gotBody *response.Response
			if err := json.Unmarshal(body, &gotBody); err != nil {
				s.Errorf(err, "unmarshal response body is failure")
			}

			s.EqualValuesf(tt.wantCode, got.StatusCode, "List() code = %v, wantCode = %v", got.StatusCode, tt.wantCode)
			if tt.wantBody != nil && !reflect.DeepEqual(gotBody, tt.wantBody) {
				s.Errorf(fmt.Errorf("List() got = %v, wantBody = %v", gotBody, tt.wantBody), "List")
			}

			s.TearDownTest()
		})
	}
}

func (s *handlerSuite) Test_impl_Create() {
	s.r.POST("/api/v1/objectives", s.handler.Create)

	type args struct {
		created *okr.Goal
		mock    func()
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
		wantBody *response.Response
	}{
		{
			name:     "empty title then 400 error",
			args:     args{created: &okr.Goal{Title: ""}},
			wantCode: 400,
			wantBody: nil,
		},
		{
			name: "created then 500 error",
			args: args{created: created1, mock: func() {
				s.mock.On("Create", mock.Anything, created1).Return(nil, errors.New("error")).Once()
			}},
			wantCode: 500,
			wantBody: nil,
		},
		{
			name: "created then 201 obj1",
			args: args{created: created1, mock: func() {
				s.mock.On("Create", mock.Anything, created1).Return(obj1, nil).Once()
			}},
			wantCode: 201,
			wantBody: response.OK.WithData(obj1),
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
			var gotBody *okr.Goal
			if err := json.Unmarshal(body, &gotBody); err != nil {
				s.Errorf(err, "unmarshal response body is failure")
			}

			s.EqualValuesf(tt.wantCode, got.StatusCode, "Create() code = %v, wantCode = %v", got.StatusCode, tt.wantCode)
			if tt.wantBody != nil && !reflect.DeepEqual(gotBody, tt.wantBody) {
				s.Errorf(fmt.Errorf("Create() got = %v, wantBody = %v", gotBody, tt.wantBody), "Create")
			}

			s.TearDownTest()
		})
	}
}

func (s *handlerSuite) Test_impl_ModifyTitle() {
	s.r.PATCH("/api/v1/goals/:id/title", s.handler.ModifyTitle)

	type args struct {
		id    string
		title string
		mock  func()
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
		wantBody *okr.Goal
	}{
		{
			name:     "id then parse id error 400",
			args:     args{id: "id", title: "title"},
			wantCode: 400,
			wantBody: nil,
		},
		{
			name:     "uuid missing title then error 400",
			args:     args{id: uuid1, title: ""},
			wantCode: 400,
			wantBody: nil,
		},
		{
			name: "uuid title then error 500",
			args: args{id: uuid1, title: "title", mock: func() {
				s.mock.On("ModifyTitle", mock.Anything, uuid1, "title").Return(nil, errors.New("error")).Once()
			}},
			wantCode: 500,
			wantBody: nil,
		},
		{
			name: "uuid title then 200",
			args: args{id: uuid1, title: "obj1", mock: func() {
				s.mock.On("ModifyTitle", mock.Anything, uuid1, "obj1").Return(obj1, nil).Once()
			}},
			wantCode: 200,
			wantBody: nil,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			uri := fmt.Sprintf("/api/v1/goals/%v/title", tt.args.id)
			data, _ := json.Marshal(&okr.Goal{Title: tt.args.title})
			req := httptest.NewRequest(http.MethodPatch, uri, bytes.NewBuffer(data))
			w := httptest.NewRecorder()
			s.r.ServeHTTP(w, req)

			got := w.Result()
			defer func() {
				_ = got.Body.Close()
			}()

			body, _ := ioutil.ReadAll(got.Body)
			var gotBody *okr.Goal
			if err := json.Unmarshal(body, &gotBody); err != nil {
				s.Errorf(err, "unmarshal response body is failure")
			}

			s.EqualValuesf(tt.wantCode, got.StatusCode, "ModifyTitle() code = %v, wantCode = %v", got.StatusCode, tt.wantCode)
			if tt.wantBody != nil && !reflect.DeepEqual(gotBody, tt.wantBody) {
				s.T().Errorf("ModifyTitle() got = %v, wantBody = %v", gotBody, tt.wantBody)
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

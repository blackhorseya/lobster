package task

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/blackhorseya/lobster/internal/app/lobster/biz/task/repo/mocks"
	"github.com/blackhorseya/lobster/internal/pkg/base/contextx"
	"github.com/blackhorseya/lobster/internal/pkg/entity/todo"
	"github.com/blackhorseya/lobster/internal/pkg/entity/user"
	"github.com/bwmarrin/snowflake"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

var (
	id1 = int64(1)

	userID1 = int64(1414348729661562880)

	info1 = &user.Profile{ID: userID1}

	ctx1 = contextx.WithValue(contextx.Background(), "user", info1)

	time1 = time.Now().UnixNano()

	task1 = &todo.Task{
		ID:        id1,
		UserID:    userID1,
		Title:     "task1",
		Status:    todo.Status_BACKLOG,
		CreatedAt: time1,
	}

	updated1 = &todo.Task{
		ID:        id1,
		UserID:    userID1,
		Title:     "updated task1",
		CreatedAt: time1,
	}

	updateStatus = &todo.Task{
		ID:        id1,
		UserID:    userID1,
		Title:     "task1",
		Status:    todo.Status_INPROGRESS,
		CreatedAt: time1,
	}
)

type bizSuite struct {
	suite.Suite
	mock *mocks.IRepo
	biz  IBiz
}

func (s *bizSuite) SetupTest() {
	logger, _ := zap.NewDevelopment()
	node, _ := snowflake.NewNode(1)

	s.mock = new(mocks.IRepo)
	if biz, err := CreateIBiz(logger, s.mock, node); err != nil {
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
		name    string
		args    args
		want    *todo.Task
		wantErr bool
	}{
		{
			name:    "missing user info then error",
			args:    args{id: id1, ctx: contextx.Background()},
			want:    nil,
			wantErr: true,
		},
		{
			name: "get by id then error",
			args: args{id: id1, ctx: ctx1, mock: func() {
				s.mock.On("QueryByID", mock.Anything, info1.ID, id1).Return(nil, errors.New("err")).Once()
			}},
			want:    nil,
			wantErr: true,
		},
		{
			name: "get by id then not exists",
			args: args{id: id1, ctx: ctx1, mock: func() {
				s.mock.On("QueryByID", mock.Anything, info1.ID, id1).Return(nil, nil).Once()
			}},
			want:    nil,
			wantErr: true,
		},
		{
			name: "get by id then task",
			args: args{id: id1, ctx: ctx1, mock: func() {
				s.mock.On("QueryByID", mock.Anything, info1.ID, id1).Return(
					task1, nil).Once()
			}},
			want:    task1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			got, err := s.biz.GetByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByID() got = %v, want %v", got, tt.want)
			}

			s.TearDownTest()
		})
	}
}

func (s *bizSuite) Test_impl_List() {
	type args struct {
		ctx  contextx.Contextx
		page int
		size int
		mock func()
	}
	tests := []struct {
		name    string
		args    args
		want    []*todo.Task
		wantErr bool
	}{
		{
			name:    "-1 10 then nil error",
			args:    args{page: -1, size: 10, ctx: ctx1},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "10 -1 then nil error",
			args:    args{page: 10, size: -1, ctx: ctx1},
			want:    nil,
			wantErr: true,
		},
		{
			name: "1 1 then nil error",
			args: args{page: 1, size: 1, ctx: ctx1, mock: func() {
				s.mock.On("List", mock.Anything, info1.ID, 0, 1).Return(
					nil, errors.New("err")).Once()
			}},
			want:    nil,
			wantErr: true,
		},
		{
			name: "1 1 then tasks nil",
			args: args{page: 1, size: 1, ctx: ctx1, mock: func() {
				s.mock.On("List", mock.Anything, info1.ID, 0, 1).Return(
					[]*todo.Task{task1}, nil).Once()
			}},
			want: []*todo.Task{
				task1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			got, err := s.biz.List(tt.args.ctx, tt.args.page, tt.args.size)
			if (err != nil) != tt.wantErr {
				t.Errorf("List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("List() got = %v, want %v", got, tt.want)
			}

			s.TearDownTest()
		})
	}
}

func (s *bizSuite) Test_impl_Create() {
	type args struct {
		ctx   contextx.Contextx
		title string
		mock  func()
	}
	tests := []struct {
		name    string
		args    args
		want    *todo.Task
		wantErr bool
	}{
		{
			name:    "missing user info in ctx then error",
			args:    args{ctx: contextx.Background(), title: "title"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "missing title then nil error",
			args:    args{title: "", ctx: ctx1},
			want:    nil,
			wantErr: true,
		},
		{
			name: "task then nil error",
			args: args{title: "task1", ctx: ctx1, mock: func() {
				s.mock.On("Create", mock.Anything, mock.Anything).Return(nil, errors.New("err")).Once()
			}},
			want:    nil,
			wantErr: true,
		},
		{
			name: "task then task nil",
			args: args{title: "task1", ctx: ctx1, mock: func() {
				s.mock.On("Create", mock.Anything, mock.Anything).Return(
					task1, nil).Once()
			}},
			want:    task1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			got, err := s.biz.Create(tt.args.ctx, tt.args.title)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
			}

			s.TearDownTest()
		})
	}
}

func (s *bizSuite) Test_impl_Delete() {
	type args struct {
		id   int64
		mock func()
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "uuid then nil error",
			args: args{id: id1, mock: func() {
				s.mock.On("Delete", mock.Anything, id1).Return(0, errors.New("err")).Once()
			}},
			wantErr: true,
		},
		{
			name: "uuid then not found",
			args: args{id: id1, mock: func() {
				s.mock.On("Delete", mock.Anything, id1).Return(0, nil).Once()
			}},
			wantErr: true,
		},
		{
			name: "uuid then nil",
			args: args{id: id1, mock: func() {
				s.mock.On("Delete", mock.Anything, id1).Return(1, nil).Once()
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			if err := s.biz.Delete(contextx.Background(), tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}

			s.TearDownTest()
		})
	}
}

func (s *bizSuite) Test_impl_UpdateStatus() {
	type args struct {
		ctx    contextx.Contextx
		id     int64
		status todo.Status
		mock   func()
	}
	tests := []struct {
		name    string
		args    args
		wantT   *todo.Task
		wantErr bool
	}{
		{
			name:    "missing user info in ctx then error",
			args:    args{id: id1, ctx: contextx.Background()},
			wantT:   nil,
			wantErr: true,
		},
		{
			name: "get by id then error",
			args: args{id: id1, ctx: ctx1, mock: func() {
				s.mock.On("QueryByID", mock.Anything, info1.ID, id1).Return(nil, errors.New("error")).Once()
			}},
			wantT:   nil,
			wantErr: true,
		},
		{
			name: "get by id then not exists",
			args: args{id: id1, ctx: ctx1, mock: func() {
				s.mock.On("QueryByID", mock.Anything, info1.ID, id1).Return(nil, nil).Once()
			}},
			wantT:   nil,
			wantErr: true,
		},
		{
			name: "update status then error",
			args: args{id: id1, ctx: ctx1, status: todo.Status_INPROGRESS, mock: func() {
				s.mock.On("QueryByID", mock.Anything, info1.ID, id1).Return(task1, nil).Once()
				s.mock.On("Update", mock.Anything, mock.Anything).Return(nil, errors.New("error")).Once()
			}},
			wantT:   nil,
			wantErr: true,
		},
		{
			name: "uuid then updated nil",
			args: args{id: id1, ctx: ctx1, status: todo.Status_INPROGRESS, mock: func() {
				s.mock.On("QueryByID", mock.Anything, info1.ID, id1).Return(task1, nil).Once()
				s.mock.On("Update", mock.Anything, mock.Anything).Return(updateStatus, nil).Once()
			}},
			wantT:   updateStatus,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotT, err := s.biz.UpdateStatus(tt.args.ctx, tt.args.id, tt.args.status)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotT, tt.wantT) {
				t.Errorf("UpdateStatus() gotT = %v, want %v", gotT, tt.wantT)
			}

			s.TearDownTest()
		})
	}
}

func (s *bizSuite) Test_impl_ModifyTitle() {
	type args struct {
		ctx   contextx.Contextx
		id    int64
		title string
		mock  func()
	}
	tests := []struct {
		name    string
		args    args
		wantT   *todo.Task
		wantErr bool
	}{
		{
			name:    "missing user info in ctx then error",
			args:    args{id: id1, ctx: contextx.Background(), title: "task"},
			wantT:   nil,
			wantErr: true,
		},
		{
			name:    "missing title then error",
			args:    args{id: id1, ctx: ctx1, title: ""},
			wantT:   nil,
			wantErr: true,
		},
		{
			name: "get by id then error",
			args: args{id: id1, ctx: ctx1, title: "title", mock: func() {
				s.mock.On("QueryByID", mock.Anything, info1.ID, id1).Return(nil, errors.New("error")).Once()
			}},
			wantT:   nil,
			wantErr: true,
		},
		{
			name: "get by id then not exists",
			args: args{id: id1, ctx: ctx1, title: "title", mock: func() {
				s.mock.On("QueryByID", mock.Anything, info1.ID, id1).Return(nil, nil).Once()
			}},
			wantT:   nil,
			wantErr: true,
		},
		{
			name: "modify title then error",
			args: args{id: id1, ctx: ctx1, title: "updated task1", mock: func() {
				s.mock.On("QueryByID", mock.Anything, info1.ID, id1).Return(task1, nil).Once()
				s.mock.On("Update", mock.Anything, updated1).Return(nil, errors.New("error")).Once()
			}},
			wantT:   nil,
			wantErr: true,
		},
		{
			name: "modify title then task",
			args: args{id: id1, ctx: ctx1, title: "updated task1", mock: func() {
				s.mock.On("QueryByID", mock.Anything, info1.ID, id1).Return(task1, nil).Once()
				s.mock.On("Update", mock.Anything, updated1).Return(updated1, nil).Once()
			}},
			wantT:   updated1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotT, err := s.biz.ModifyTitle(tt.args.ctx, tt.args.id, tt.args.title)
			if (err != nil) != tt.wantErr {
				t.Errorf("ModifyTitle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotT, tt.wantT) {
				t.Errorf("ModifyTitle() gotT = %v, want %v", gotT, tt.wantT)
			}

			s.TearDownTest()
		})
	}
}

package task

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/blackhorseya/lobster/internal/app/lobster/biz/task/repo/mocks"
	"github.com/blackhorseya/lobster/internal/pkg/contextx"
	"github.com/blackhorseya/lobster/internal/pkg/entities/task"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

var (
	uuid1 = "d76f4f51-f141-41ba-ba57-c4749319586b"

	time1 = time.Now().UnixNano()

	task1 = &task.Task{
		ID:        uuid1,
		Title:     "task1",
		Status:    task.Status_BACKLOG,
		Completed: false,
		CreateAt:  time1,
	}

	updated1 = &task.Task{
		ID:        uuid1,
		Title:     "updated task1",
		Completed: false,
		CreateAt:  time1,
	}

	updateStatus = &task.Task{
		ID:        uuid1,
		Title:     "updated task1",
		Status:    task.Status_INPROGRESS,
		Completed: false,
		CreateAt:  time1,
	}
)

type bizSuite struct {
	suite.Suite
	mock *mocks.IRepo
	biz  IBiz
}

func (s *bizSuite) SetupTest() {
	logger, _ := zap.NewDevelopment()

	s.mock = new(mocks.IRepo)
	if biz, err := CreateIBiz(logger, s.mock); err != nil {
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
		id   string
		mock func()
	}
	tests := []struct {
		name    string
		args    args
		want    *task.Task
		wantErr bool
	}{
		{
			name:    "id then nil error",
			args:    args{id: "id"},
			want:    nil,
			wantErr: true,
		},
		{
			name: "uuid then nil error",
			args: args{id: uuid1, mock: func() {
				s.mock.On("QueryByID", mock.Anything, uuid1).Return(
					nil, errors.New("err")).Once()
			}},
			want:    nil,
			wantErr: true,
		},
		{
			name: "uuid then nil not exists",
			args: args{id: uuid1, mock: func() {
				s.mock.On("QueryByID", mock.Anything, uuid1).Return(nil, nil).Once()
			}},
			want:    nil,
			wantErr: true,
		},
		{
			name: "uuid then task nil",
			args: args{id: uuid1, mock: func() {
				s.mock.On("QueryByID", mock.Anything, uuid1).Return(
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

			got, err := s.biz.GetByID(contextx.Background(), tt.args.id)
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
		page int
		size int
		mock func()
	}
	tests := []struct {
		name    string
		args    args
		want    []*task.Task
		wantErr bool
	}{
		{
			name:    "-1 10 then nil error",
			args:    args{page: -1, size: 10},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "10 -1 then nil error",
			args:    args{page: 10, size: -1},
			want:    nil,
			wantErr: true,
		},
		{
			name: "1 1 then nil error",
			args: args{page: 1, size: 1, mock: func() {
				s.mock.On("List", mock.Anything, 0, 1).Return(
					nil, errors.New("err")).Once()
			}},
			want:    nil,
			wantErr: true,
		},
		{
			name: "1 1 then tasks nil",
			args: args{page: 1, size: 1, mock: func() {
				s.mock.On("List", mock.Anything, 0, 1).Return(
					[]*task.Task{task1}, nil).Once()
			}},
			want: []*task.Task{
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

			got, err := s.biz.List(contextx.Background(), tt.args.page, tt.args.size)
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

func (s *bizSuite) Test_impl_Count() {
	type args struct {
		mock func()
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "count then 0 error",
			args: args{mock: func() {
				s.mock.On("Count", mock.Anything).Return(0, errors.New("err")).Once()
			}},
			want:    0,
			wantErr: true,
		},
		{
			name: "count then 0 not found",
			args: args{mock: func() {
				s.mock.On("Count", mock.Anything).Return(0, nil).Once()
			}},
			want:    0,
			wantErr: true,
		},
		{
			name: "count then 10 nil",
			args: args{mock: func() {
				s.mock.On("Count", mock.Anything).Return(10, nil).Once()
			}},
			want:    10,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			got, err := s.biz.Count(contextx.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("Count() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Count() got = %v, want %v", got, tt.want)
			}

			s.TearDownTest()
		})
	}
}

func (s *bizSuite) Test_impl_Create() {
	type args struct {
		task *task.Task
		mock func()
	}
	tests := []struct {
		name    string
		args    args
		want    *task.Task
		wantErr bool
	}{
		{
			name:    "missing title then nil error",
			args:    args{task: &task.Task{Title: ""}},
			want:    nil,
			wantErr: true,
		},
		{
			name: "task then nil error",
			args: args{task: &task.Task{Title: "task1"}, mock: func() {
				s.mock.On("Create", mock.Anything, mock.Anything).Return(
					nil, errors.New("err")).Once()
			}},
			want:    nil,
			wantErr: true,
		},
		{
			name: "task then task nil",
			args: args{task: &task.Task{Title: "task1"}, mock: func() {
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

			got, err := s.biz.Create(contextx.Background(), tt.args.task)
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
		id   string
		mock func()
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "id then nil error",
			args:    args{id: "id"},
			wantErr: true,
		},
		{
			name: "uuid then nil error",
			args: args{id: uuid1, mock: func() {
				s.mock.On("Delete", mock.Anything, uuid1).Return(0, errors.New("err")).Once()
			}},
			wantErr: true,
		},
		{
			name: "uuid then not found",
			args: args{id: uuid1, mock: func() {
				s.mock.On("Delete", mock.Anything, uuid1).Return(0, nil).Once()
			}},
			wantErr: true,
		},
		{
			name: "uuid then nil",
			args: args{id: uuid1, mock: func() {
				s.mock.On("Delete", mock.Anything, uuid1).Return(1, nil).Once()
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
		id     string
		status task.Status
		mock   func()
	}
	tests := []struct {
		name    string
		args    args
		wantT   *task.Task
		wantErr bool
	}{
		{
			name:    "id then nil error",
			args:    args{id: "id"},
			wantT:   nil,
			wantErr: true,
		},
		{
			name: "uuid then query id error",
			args: args{id: uuid1, mock: func() {
				s.mock.On("QueryByID", mock.Anything, uuid1).Return(nil, errors.New("error")).Once()
			}},
			wantT:   nil,
			wantErr: true,
		},
		{
			name: "uuid then query id not found",
			args: args{id: uuid1, mock: func() {
				s.mock.On("QueryByID", mock.Anything, uuid1).Return(nil, nil).Once()
			}},
			wantT:   nil,
			wantErr: true,
		},
		{
			name: "uuid then update error",
			args: args{id: uuid1, status: task.Status_INPROGRESS, mock: func() {
				s.mock.On("QueryByID", mock.Anything, uuid1).Return(task1, nil).Once()
				s.mock.On("Update", mock.Anything, updateStatus).Return(nil, errors.New("error")).Once()
			}},
			wantT:   nil,
			wantErr: true,
		},
		{
			name: "uuid then updated nil",
			args: args{id: uuid1, status: task.Status_INPROGRESS, mock: func() {
				s.mock.On("QueryByID", mock.Anything, uuid1).Return(task1, nil).Once()
				s.mock.On("Update", mock.Anything, updateStatus).Return(updateStatus, nil).Once()
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

			gotT, err := s.biz.UpdateStatus(contextx.Background(), tt.args.id, tt.args.status)
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
		id    string
		title string
		mock  func()
	}
	tests := []struct {
		name    string
		args    args
		wantT   *task.Task
		wantErr bool
	}{
		{
			name:    "id then parse id error",
			args:    args{id: "id", title: "title"},
			wantT:   nil,
			wantErr: true,
		},
		{
			name:    "uuid missing title then error",
			args:    args{id: uuid1, title: ""},
			wantT:   nil,
			wantErr: true,
		},
		{
			name: "uuid title then query id error",
			args: args{id: uuid1, title: "title", mock: func() {
				s.mock.On("QueryByID", mock.Anything, uuid1).Return(nil, errors.New("error")).Once()
			}},
			wantT:   nil,
			wantErr: true,
		},
		{
			name: "uuid title then query id not found",
			args: args{id: uuid1, title: "title", mock: func() {
				s.mock.On("QueryByID", mock.Anything, uuid1).Return(nil, nil).Once()
			}},
			wantT:   nil,
			wantErr: true,
		},
		{
			name: "uuid title then modify title error",
			args: args{id: uuid1, title: "updated task1", mock: func() {
				s.mock.On("QueryByID", mock.Anything, uuid1).Return(task1, nil).Once()
				s.mock.On("Update", mock.Anything, updated1).Return(nil, errors.New("error")).Once()
			}},
			wantT:   nil,
			wantErr: true,
		},
		{
			name: "uuid title then modify title",
			args: args{id: uuid1, title: "updated task1", mock: func() {
				s.mock.On("QueryByID", mock.Anything, uuid1).Return(task1, nil).Once()
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

			gotT, err := s.biz.ModifyTitle(contextx.Background(), tt.args.id, tt.args.title)
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

package todo

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/blackhorseya/lobster/internal/app/biz/todo/repo/mocks"
	"github.com/blackhorseya/lobster/internal/pkg/contextx"
	"github.com/blackhorseya/lobster/internal/pkg/entities/biz/todo"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var (
	uuid1 = "d76f4f51-f141-41ba-ba57-c4749319586b"

	time1 = time.Now().UnixNano()

	task1 = &todo.Task{
		ID:        uuid1,
		Title:     "task1",
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
	s.mock = new(mocks.IRepo)
	if biz, err := CreateIBiz(s.mock); err != nil {
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
		want    *todo.Task
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

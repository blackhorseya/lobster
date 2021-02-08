package objective

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/blackhorseya/lobster/internal/biz/objective/repo/mocks"
	"github.com/blackhorseya/lobster/internal/pkg/contextx"
	"github.com/blackhorseya/lobster/internal/pkg/entities/biz/okr"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var (
	uuid1 = "d76f4f51-f141-41ba-ba57-c4749319586b"

	time1 = time.Now().UnixNano()

	emptyTitle = &okr.Objective{Title: ""}

	created1 = &okr.Objective{
		Title: "obj1",
	}

	obj1 = &okr.Objective{
		ID:       uuid1,
		Title:    "obj1",
		CreateAt: time1,
	}

	updated1 = &okr.Objective{
		ID:       uuid1,
		Title:    "updated obj1",
		CreateAt: time1,
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

func (s *bizSuite) Test_impl_Create() {
	type args struct {
		obj  *okr.Objective
		mock func()
	}
	tests := []struct {
		name    string
		args    args
		want    *okr.Objective
		wantErr bool
	}{
		{
			name:    "empty title then nil error",
			args:    args{obj: emptyTitle},
			want:    nil,
			wantErr: true,
		},
		{
			name: "created then nil error",
			args: args{obj: created1, mock: func() {
				s.mock.On("Create", mock.Anything, mock.Anything).Return(
					nil, errors.New("error")).Once()
			}},
			want:    nil,
			wantErr: true,
		},
		{
			name: "created then obj1 nil",
			args: args{obj: created1, mock: func() {
				s.mock.On("Create", mock.Anything, mock.Anything).Return(obj1, nil).Once()
			}},
			want:    obj1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			got, err := s.biz.Create(contextx.Background(), tt.args.obj)
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

func (s *bizSuite) Test_impl_List() {
	type args struct {
		page int
		size int
		mock func()
	}
	tests := []struct {
		name    string
		args    args
		want    []*okr.Objective
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
				s.mock.On("List", mock.Anything, 0, 1).Return(nil, errors.New("error")).Once()
			}},
			want:    nil,
			wantErr: true,
		},
		{
			name: "10 10 then nil error",
			args: args{page: 10, size: 10, mock: func() {
				s.mock.On("List", mock.Anything, 90, 10).Return(nil, nil).Once()
			}},
			want:    nil,
			wantErr: true,
		},
		{
			name: "1 1 then objs nil",
			args: args{page: 1, size: 1, mock: func() {
				s.mock.On("List", mock.Anything, 0, 1).Return([]*okr.Objective{obj1}, nil).Once()
			}},
			want:    []*okr.Objective{obj1},
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
				s.mock.On("Count", mock.Anything).Return(0, errors.New("error")).Once()
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

func (s *bizSuite) Test_impl_Update() {
	type args struct {
		updated *okr.Objective
		mock    func()
	}
	tests := []struct {
		name    string
		args    args
		want    *okr.Objective
		wantErr bool
	}{
		{
			name:    "empty title then nil error",
			args:    args{updated: &okr.Objective{ID: uuid1, Title: ""}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "id title then nil error",
			args:    args{updated: &okr.Objective{ID: "id", Title: "obj1"}},
			want:    nil,
			wantErr: true,
		},
		{
			name: "uuid title then query error",
			args: args{updated: &okr.Objective{ID: uuid1, Title: "obj1"}, mock: func() {
				s.mock.On("QueryByID", mock.Anything, uuid1).Return(nil, errors.New("error")).Once()
			}},
			want:    nil,
			wantErr: true,
		},
		{
			name: "uuid title then query not found error",
			args: args{updated: &okr.Objective{ID: uuid1, Title: "obj1"}, mock: func() {
				s.mock.On("QueryByID", mock.Anything, uuid1).Return(nil, nil).Once()
			}},
			want:    nil,
			wantErr: true,
		},
		{
			name: "uuid title then nil error",
			args: args{updated: &okr.Objective{ID: uuid1, Title: "obj1"}, mock: func() {
				s.mock.On("QueryByID", mock.Anything, uuid1).Return(obj1, nil).Once()
				s.mock.On("Update", mock.Anything, mock.Anything).Return(nil, errors.New("error")).Once()
			}},
			want:    nil,
			wantErr: true,
		},
		{
			name: "uuid title then obj1 nil",
			args: args{updated: &okr.Objective{ID: uuid1, Title: "updated obj1"}, mock: func() {
				s.mock.On("QueryByID", mock.Anything, uuid1).Return(obj1, nil).Once()
				s.mock.On("Update", mock.Anything, mock.Anything).Return(updated1, nil).Once()
			}},
			want:    updated1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			got, err := s.biz.Update(contextx.Background(), tt.args.updated)
			if (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Update() got = %v, want %v", got, tt.want)
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
			name: "uuid then error",
			args: args{id: uuid1, mock: func() {
				s.mock.On("Delete", mock.Anything, uuid1).Return(0, errors.New("error")).Once()
			}},
			wantErr: true,
		},
		{
			name: "uuid then not found error",
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

func (s *bizSuite) Test_impl_GetByID() {
	type args struct {
		id   string
		mock func()
	}
	tests := []struct {
		name    string
		args    args
		want    *okr.Objective
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
				s.mock.On("QueryByID", mock.Anything, uuid1).Return(nil, errors.New("error")).Once()
			}},
			want:    nil,
			wantErr: true,
		},
		{
			name: "uuid then obj nil",
			args: args{id: uuid1, mock: func() {
				s.mock.On("QueryByID", mock.Anything, uuid1).Return(obj1, nil).Once()
			}},
			want:    obj1,
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

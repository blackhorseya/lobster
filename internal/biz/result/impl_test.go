package result

import (
	"errors"
	"reflect"
	"testing"

	"github.com/blackhorseya/lobster/internal/biz/result/repo/mocks"
	"github.com/blackhorseya/lobster/internal/pkg/contextx"
	"github.com/blackhorseya/lobster/internal/pkg/pb"
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

	updated1 = &pb.Result{
		ID:       krID,
		GoalID:   goalID,
		Title:    "updated kr1",
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
	biz, err := CreateIBiz(s.mock)
	if err != nil {
		panic(err)
	}

	s.biz = biz
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
		wantKr  *pb.Result
		wantErr bool
	}{
		{
			name:    "id then nil error",
			args:    args{id: "id"},
			wantKr:  nil,
			wantErr: true,
		},
		{
			name: "uuid then nil error",
			args: args{id: krID, mock: func() {
				s.mock.On("QueryByID", mock.Anything, krID).Return(nil, errors.New("error")).Once()
			}},
			wantKr:  nil,
			wantErr: true,
		},
		{
			name: "uuid then not found error",
			args: args{id: krID, mock: func() {
				s.mock.On("QueryByID", mock.Anything, krID).Return(nil, nil).Once()
			}},
			wantKr:  nil,
			wantErr: true,
		},
		{
			name: "uuid then kr nil",
			args: args{id: krID, mock: func() {
				s.mock.On("QueryByID", mock.Anything, krID).Return(kr1, nil).Once()
			}},
			wantKr:  kr1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotKr, err := s.biz.GetByID(contextx.Background(), tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotKr, tt.wantKr) {
				t.Errorf("GetByID() gotKr = %v, want %v", gotKr, tt.wantKr)
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
		wantKrs []*pb.Result
		wantErr bool
	}{
		{
			name:    "-1 10 then nil error",
			args:    args{page: -1, size: 10},
			wantKrs: nil,
			wantErr: true,
		},
		{
			name:    "10 -1 then nil error",
			args:    args{page: 10, size: -1},
			wantKrs: nil,
			wantErr: true,
		},
		{
			name: "1 1 then nil error",
			args: args{page: 1, size: 1, mock: func() {
				s.mock.On("QueryList", mock.Anything, 0, 1).Return(nil, errors.New("error")).Once()
			}},
			wantKrs: nil,
			wantErr: true,
		},
		{
			name: "1 1 then krs nil",
			args: args{page: 1, size: 1, mock: func() {
				s.mock.On("QueryList", mock.Anything, 0, 1).Return([]*pb.Result{kr1}, nil).Once()
			}},
			wantKrs: []*pb.Result{kr1},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotKrs, err := s.biz.List(contextx.Background(), tt.args.page, tt.args.size)
			if (err != nil) != tt.wantErr {
				t.Errorf("List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotKrs, tt.wantKrs) {
				t.Errorf("List() gotKrs = %v, want %v", gotKrs, tt.wantKrs)
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
			name:    "id then error",
			args:    args{id: "id"},
			wantErr: true,
		},
		{
			name: "uuid then error",
			args: args{id: krID, mock: func() {
				s.mock.On("Delete", mock.Anything, krID).Return(errors.New("error")).Once()
			}},
			wantErr: true,
		},
		{
			name: "uuid then nil",
			args: args{id: krID, mock: func() {
				s.mock.On("Delete", mock.Anything, krID).Return(nil).Once()
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

func (s *bizSuite) Test_impl_Update() {
	type args struct {
		updated *pb.Result
		mock    func()
	}
	tests := []struct {
		name    string
		args    args
		wantKr  *pb.Result
		wantErr bool
	}{
		{
			name:    "id title then nil error",
			args:    args{updated: &pb.Result{ID: "id", GoalID: goalID, Title: "title"}},
			wantKr:  nil,
			wantErr: true,
		},
		{
			name:    "uuid missing title then nil error",
			args:    args{updated: &pb.Result{ID: krID, GoalID: goalID, Title: ""}},
			wantKr:  nil,
			wantErr: true,
		},
		{
			name:    "goal id then nil error",
			args:    args{updated: &pb.Result{ID: krID, GoalID: "id", Title: "title"}},
			wantKr:  nil,
			wantErr: true,
		},
		{
			name: "uuid then query error",
			args: args{updated: updated1, mock: func() {
				s.mock.On("QueryByID", mock.Anything, krID).Return(nil, errors.New("error")).Once()
			}},
			wantKr:  nil,
			wantErr: true,
		},
		{
			name: "uuid then query not found error",
			args: args{updated: updated1, mock: func() {
				s.mock.On("QueryByID", mock.Anything, krID).Return(nil, nil).Once()
			}},
			wantKr:  nil,
			wantErr: true,
		},
		{
			name: "uuid then update error",
			args: args{updated: updated1, mock: func() {
				s.mock.On("QueryByID", mock.Anything, krID).Return(kr1, nil).Once()
				s.mock.On("Update", mock.Anything, updated1).Return(nil, errors.New("error")).Once()
			}},
			wantKr:  nil,
			wantErr: true,
		},
		{
			name: "uuid then updated1 nil",
			args: args{updated: updated1, mock: func() {
				s.mock.On("QueryByID", mock.Anything, updated1.ID).Return(kr1, nil).Once()
				s.mock.On("Update", mock.Anything, updated1).Return(updated1, nil).Once()
			}},
			wantKr:  updated1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotKr, err := s.biz.Update(contextx.Background(), tt.args.updated)
			if (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotKr, tt.wantKr) {
				t.Errorf("Update() gotKr = %v, want %v", gotKr, tt.wantKr)
			}

			s.TearDownTest()
		})
	}
}

func (s *bizSuite) Test_impl_LinkToGoal() {
	type args struct {
		created *pb.Result
		mock    func()
	}
	tests := []struct {
		name    string
		args    args
		wantKr  *pb.Result
		wantErr bool
	}{
		{
			name:    "missing title then nil error",
			args:    args{created: &pb.Result{ID: krID, GoalID: goalID, Title: ""}},
			wantKr:  nil,
			wantErr: true,
		},
		{
			name:    "goal id not uuid then nil error",
			args:    args{created: &pb.Result{ID: krID, GoalID: "id", Title: "title"}},
			wantKr:  nil,
			wantErr: true,
		},
		{
			name: "create then nil error",
			args: args{created: kr1, mock: func() {
				s.mock.On("Create", mock.Anything, kr1).Return(nil, errors.New("error")).Once()
			}},
			wantKr:  nil,
			wantErr: true,
		},
		{
			name: "create then kr nil",
			args: args{created: kr1, mock: func() {
				s.mock.On("Create", mock.Anything, kr1).Return(kr1, nil).Once()
			}},
			wantKr:  kr1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotKr, err := s.biz.LinkToGoal(contextx.Background(), tt.args.created)
			if (err != nil) != tt.wantErr {
				t.Errorf("LinkToGoal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotKr, tt.wantKr) {
				t.Errorf("LinkToGoal() gotKr = %v, want %v", gotKr, tt.wantKr)
			}

			s.TearDownTest()
		})
	}
}

func (s *bizSuite) Test_impl_GetByGoalID() {
	type args struct {
		id   string
		mock func()
	}
	tests := []struct {
		name    string
		args    args
		wantKrs []*pb.Result
		wantErr bool
	}{
		{
			name:    "id then parse id error",
			args:    args{id: "id"},
			wantKrs: nil,
			wantErr: true,
		},
		{
			name: "uuid then query error",
			args: args{id: goalID, mock: func() {
				s.mock.On("QueryByGoalID", mock.Anything, goalID).Return(nil, errors.New("error")).Once()
			}},
			wantKrs: nil,
			wantErr: true,
		},
		{
			name: "uuid then query krs",
			args: args{id: goalID, mock: func() {
				s.mock.On("QueryByGoalID", mock.Anything, goalID).Return([]*pb.Result{kr1}, nil).Once()
			}},
			wantKrs: []*pb.Result{kr1},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			gotKrs, err := s.biz.GetByGoalID(contextx.Background(), tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByGoalID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotKrs, tt.wantKrs) {
				t.Errorf("GetByGoalID() gotKrs = %v, want %v", gotKrs, tt.wantKrs)
			}

			s.TearDownTest()
		})
	}
}

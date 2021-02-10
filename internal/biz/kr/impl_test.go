package kr

import (
	"errors"
	"reflect"
	"testing"

	"github.com/blackhorseya/lobster/internal/biz/kr/repo/mocks"
	"github.com/blackhorseya/lobster/internal/pkg/contextx"
	"github.com/blackhorseya/lobster/internal/pkg/entities/biz/okr"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

var (
	krID   = "d76f4f51-f141-41ba-ba57-c4749319586b"
	goalID = "0829ee06-1f04-43d9-8565-812e1826f805"

	time1 = int64(1611059529208050000)

	kr1 = &okr.KeyResult{
		ID:       krID,
		GoalID:   goalID,
		Title:    "kr1",
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
		wantKr  *okr.KeyResult
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

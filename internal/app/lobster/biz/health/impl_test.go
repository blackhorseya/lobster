package health

import (
	"errors"
	"testing"

	"github.com/blackhorseya/lobster/internal/app/lobster/biz/health/repo/mocks"
	"github.com/blackhorseya/lobster/internal/pkg/contextx"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
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

func (s *bizSuite) Test_impl_Readiness() {
	type args struct {
		mock func()
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "readiness then error",
			args: args{mock: func() {
				s.mock.On("Ping", mock.Anything).Return(false, errors.New("error")).Once()
			}},
			wantErr: true,
		},
		{
			name: "readiness then false error",
			args: args{mock: func() {
				s.mock.On("Ping", mock.Anything).Return(false, nil).Once()
			}},
			wantErr: true,
		},
		{
			name: "readiness then nil",
			args: args{mock: func() {
				s.mock.On("Ping", mock.Anything).Return(true, nil).Once()
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			if err := s.biz.Readiness(contextx.Background()); (err != nil) != tt.wantErr {
				t.Errorf("Readiness() error = %v, wantErr %v", err, tt.wantErr)
			}

			s.TearDownTest()
		})
	}
}

func (s *bizSuite) Test_impl_Liveness() {
	type args struct {
		mock func()
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "liveness then error",
			args: args{mock: func() {
				s.mock.On("Ping", mock.Anything).Return(false, errors.New("error")).Once()
			}},
			wantErr: true,
		},
		{
			name: "liveness then false error",
			args: args{mock: func() {
				s.mock.On("Ping", mock.Anything).Return(false, nil).Once()
			}},
			wantErr: true,
		},
		{
			name: "liveness then nil",
			args: args{mock: func() {
				s.mock.On("Ping", mock.Anything).Return(true, nil).Once()
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			if err := s.biz.Liveness(contextx.Background()); (err != nil) != tt.wantErr {
				t.Errorf("Liveness() error = %v, wantErr %v", err, tt.wantErr)
			}

			s.TearDownTest()
		})
	}
}

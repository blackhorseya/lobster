// +build integration

package repo

import (
	"reflect"
	"testing"

	"github.com/blackhorseya/lobster/internal/pkg/contextx"
	"github.com/blackhorseya/lobster/internal/pkg/pb"
	"github.com/stretchr/testify/suite"
)

var (
	krID1 = "d0f6de2d-d8cd-4259-85e9-220bdd81e1cf"
	krID2 = "471322e2-4682-40dd-b344-dbf07800ad5c"

	goalID = "d0f6de2d-d8cd-4259-85e9-220bdd81e1cf"

	time1 = int64(1611059529208050000)

	kr1 = &pb.KeyResult{
		ID:       krID1,
		GoalID:   goalID,
		Title:    "kr1",
		Target: 99,
		Actual: 10,
		CreateAt: int64(1613114039486249000),
	}

	kr2 = &pb.KeyResult{
		ID:       krID2,
		GoalID:   goalID,
		Title:    "kr",
		Target: 100,
		Actual: 20,
		CreateAt: int64(1614032353580795000),
	}

	updated1 = &pb.KeyResult{
		ID:       krID1,
		GoalID:   goalID,
		Title:    "updated kr1",
		CreateAt: time1,
	}
)

type repoSuite struct {
	suite.Suite
	repo IRepo
}

func (s *repoSuite) SetupTest() {
	if repo, err := CreateRepo("../../../../configs/app.yaml"); err != nil {
		panic(err)
	} else {
		s.repo = repo
	}
}

func TestRepoSuite(t *testing.T) {
	suite.Run(t, new(repoSuite))
}

func (s *repoSuite) Test_impl_Create() {
	type args struct {
		created *pb.KeyResult
	}
	tests := []struct {
		name    string
		args    args
		wantKr  *pb.KeyResult
		wantErr bool
	}{
		{
			name:    "created then kr nil",
			args:    args{created: kr1},
			wantKr:  kr1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			gotKr, err := s.repo.Create(contextx.Background(), tt.args.created)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotKr, tt.wantKr) {
				t.Errorf("Create() gotKr = %v, want %v", gotKr, tt.wantKr)
			}
		})
	}
}

func (s *repoSuite) Test_impl_QueryKRByID() {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    *pb.KeyResult
		wantErr bool
	}{
		{
			name:    "goalID id then kr nil",
			args:    args{id: krID1},
			want:    kr1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			got, err := s.repo.QueryByID(contextx.Background(), tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QueryByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func (s *repoSuite) Test_impl_Update() {
	type args struct {
		updated *pb.KeyResult
	}
	tests := []struct {
		name    string
		args    args
		wantKr  *pb.KeyResult
		wantErr bool
	}{
		{
			name:    "update then kr nil",
			args:    args{updated: updated1},
			wantKr:  updated1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			gotKr, err := s.repo.Update(contextx.Background(), tt.args.updated)
			if (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotKr, tt.wantKr) {
				t.Errorf("Update() gotKr = %v, want %v", gotKr, tt.wantKr)
			}
		})
	}
}

func (s *repoSuite) Test_impl_QueryList() {
	type args struct {
		offset int
		limit  int
	}
	tests := []struct {
		name    string
		args    args
		wantKrs []*pb.KeyResult
		wantErr bool
	}{
		{
			name:    "list then krs nil",
			args:    args{offset: 0, limit: 1},
			wantKrs: []*pb.KeyResult{kr1},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			gotKrs, err := s.repo.QueryList(contextx.Background(), tt.args.offset, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotKrs, tt.wantKrs) {
				t.Errorf("QueryList() gotKrs = %v, want %v", gotKrs, tt.wantKrs)
			}
		})
	}
}

func (s *repoSuite) Test_impl_Delete() {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "uuid then nil",
			args:    args{id: krID1},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if err := s.repo.Delete(contextx.Background(), tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func (s *repoSuite) Test_impl_QueryByGoalID() {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		wantKrs []*pb.KeyResult
		wantErr bool
	}{
		{
			name:    "uuid then krs nil",
			args:    args{id: goalID},
			wantKrs: []*pb.KeyResult{kr2, kr1},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			gotKrs, err := s.repo.QueryByGoalID(contextx.Background(), tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryByGoalID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotKrs, tt.wantKrs) {
				t.Errorf("QueryByGoalID() gotKrs = %v, want %v", gotKrs, tt.wantKrs)
			}
		})
	}
}

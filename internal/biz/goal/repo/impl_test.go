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
	uuid1 = "d76f4f51-f141-41ba-ba57-c4749319586b"

	time1 = int64(1611059529208050000)

	obj1 = &pb.Objective{
		ID:       uuid1,
		Title:    "obj1",
		CreateAt: time1,
	}

	updated1 = &pb.Objective{
		ID:       uuid1,
		Title:    "updated obj1",
		CreateAt: time1,
	}
)

type repoSuite struct {
	suite.Suite
	repo IRepo
}

func (s *repoSuite) SetupTest() {
	if repo, err := CreateRepo("../../../../../configs/app.yaml"); err != nil {
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
		created *pb.Objective
	}
	tests := []struct {
		name    string
		args    args
		want    *pb.Objective
		wantErr bool
	}{
		{
			name:    "created then obj nil",
			args:    args{created: obj1},
			want:    obj1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			got, err := s.repo.Create(contextx.Background(), tt.args.created)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func (s *repoSuite) Test_impl_QueryByID() {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    *pb.Objective
		wantErr bool
	}{
		{
			name:    "uuid then obj1 nil",
			args:    args{id: uuid1},
			want:    obj1,
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

func (s *repoSuite) Test_impl_List() {
	type args struct {
		offset int
		limit  int
	}
	tests := []struct {
		name    string
		args    args
		want    []*pb.Objective
		wantErr bool
	}{
		{
			name:    "0 1 then objs nil",
			args:    args{offset: 0, limit: 1},
			want:    []*pb.Objective{obj1},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			got, err := s.repo.List(contextx.Background(), tt.args.offset, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("List() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func (s *repoSuite) Test_impl_Count() {
	type args struct {
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    "count then 1 nil",
			args:    args{},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			got, err := s.repo.Count(contextx.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("Count() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Count() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func (s *repoSuite) Test_impl_Update() {
	type args struct {
		updated *pb.Objective
	}
	tests := []struct {
		name    string
		args    args
		want    *pb.Objective
		wantErr bool
	}{
		{
			name:    "update then obj1 nil",
			args:    args{updated: updated1},
			want:    updated1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			got, err := s.repo.Update(contextx.Background(), tt.args.updated)
			if (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Update() got = %v, want %v", got, tt.want)
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
		want    int
		wantErr bool
	}{
		{
			name:    "uuid then 1 nil",
			args:    args{id: uuid1},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			got, err := s.repo.Delete(contextx.Background(), tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Delete() got = %v, want %v", got, tt.want)
			}
		})
	}
}

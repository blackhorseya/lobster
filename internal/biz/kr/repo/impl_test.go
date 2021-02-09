// +build integration

package repo

import (
	"reflect"
	"testing"

	"github.com/blackhorseya/lobster/internal/pkg/contextx"
	"github.com/blackhorseya/lobster/internal/pkg/entities/biz/okr"
	"github.com/stretchr/testify/suite"
)

var (
	krID = "d76f4f51-f141-41ba-ba57-c4749319586b"

	goalID = "0829ee06-1f04-43d9-8565-812e1826f805"

	time1 = int64(1611059529208050000)

	kr1 = &okr.KeyResult{
		ID:       krID,
		Title:    "kr1",
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

func (s *repoSuite) Test_impl_QueryKRByID() {
	type args struct {
		goalID string
		krID   string
	}
	tests := []struct {
		name    string
		args    args
		want    *okr.KeyResult
		wantErr bool
	}{
		{
			name:    "goalID krID then kr nil",
			args:    args{goalID: goalID, krID: krID},
			want:    kr1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			got, err := s.repo.QueryKRByID(contextx.Background(), tt.args.goalID, tt.args.krID)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryKRByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QueryKRByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

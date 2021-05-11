// +build integration

package repo

import (
	"reflect"
	"testing"

	"github.com/blackhorseya/lobster/internal/pkg/base/contextx"
	"github.com/blackhorseya/lobster/internal/pkg/entity/user"
	"github.com/stretchr/testify/suite"
)

var (
	uuid1 = "d76f4f51-f141-41ba-ba57-c4749319586b"

	time1 = int64(1610548520788105000)

	token1 = "b54c851b9d9e030f2afd6f6119b9c84e59f02590"

	email1 = "test@gmail.com"

	user1 = user.Profile{
		ID:          uuid1,
		AccessToken: token1,
		Email:       email1,
		SignupAt:    time1,
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

func (s *repoSuite) Test_impl_UserRegister() {
	type args struct {
		newUser user.Profile
	}
	tests := []struct {
		name     string
		args     args
		wantInfo *user.Profile
		wantErr  bool
	}{
		{
			name:     "profile then profile nil",
			args:     args{newUser: user1},
			wantInfo: &user1,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			gotInfo, err := s.repo.UserRegister(contextx.Background(), tt.args.newUser)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRegister() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInfo, tt.wantInfo) {
				t.Errorf("UserRegister() gotInfo = %v, want %v", gotInfo, tt.wantInfo)
			}
		})
	}
}

func (s *repoSuite) Test_impl_QueryInfoByEmail() {
	type args struct {
		email string
	}
	tests := []struct {
		name     string
		args     args
		wantInfo *user.Profile
		wantErr  bool
	}{
		{
			name:     "email then profile nil",
			args:     args{email: email1},
			wantInfo: &user1,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			gotInfo, err := s.repo.QueryInfoByEmail(contextx.Background(), tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryInfoByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInfo, tt.wantInfo) {
				t.Errorf("QueryInfoByEmail() gotInfo = %v, want %v", gotInfo, tt.wantInfo)
			}
		})
	}
}

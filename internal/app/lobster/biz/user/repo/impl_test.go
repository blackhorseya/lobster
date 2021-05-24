// +build integration

package repo

import (
	"reflect"
	"testing"

	"github.com/blackhorseya/lobster/internal/pkg/base/contextx"
	"github.com/blackhorseya/lobster/internal/pkg/base/encrypt"
	"github.com/blackhorseya/lobster/internal/pkg/entity/user"
	"github.com/stretchr/testify/suite"
)

var (
	email1 = "test@gmail.com"

	password1 = "password"

	encPWD, _ = encrypt.HashAndSalt(password1)

	user1 = &user.Profile{ID: 1, Email: email1, Password: encPWD, AccessToken: ""}
)

type repoSuite struct {
	suite.Suite
	repo IRepo
}

func (s *repoSuite) SetupTest() {
	repo, err := CreateIRepo("../../../../../../configs/app.yaml")
	if err != nil {
		panic(err)
	}

	s.repo = repo
}

func TestRepoSuite(t *testing.T) {
	suite.Run(t, new(repoSuite))
}

func (s *repoSuite) Test_impl_Register() {
	type args struct {
		email    string
		password string
	}
	tests := []struct {
		name     string
		args     args
		wantInfo *user.Profile
		wantErr  bool
	}{
		{
			name:     "register then success",
			args:     args{email: email1, password: encPWD},
			wantInfo: user1,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			gotInfo, err := s.repo.Register(contextx.Background(), tt.args.email, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInfo, tt.wantInfo) {
				t.Errorf("Register() gotInfo = %v, want %v", gotInfo, tt.wantInfo)
			}
		})
	}
}

func (s *repoSuite) Test_impl_GetByEmail() {
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
			name:     "get by email then success",
			args:     args{email: email1},
			wantInfo: user1,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			gotInfo, err := s.repo.GetByEmail(contextx.Background(), tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInfo, tt.wantInfo) {
				t.Errorf("GetByEmail() gotInfo = %v, want %v", gotInfo, tt.wantInfo)
			}
		})
	}
}

func (s *repoSuite) Test_impl_UpdateToken() {
	type args struct {
		updated *user.Profile
	}
	tests := []struct {
		name     string
		args     args
		wantInfo *user.Profile
		wantErr  bool
	}{
		{
			name:     "update token then success",
			args:     args{updated: user1},
			wantInfo: user1,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			gotInfo, err := s.repo.UpdateToken(contextx.Background(), tt.args.updated)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInfo, tt.wantInfo) {
				t.Errorf("UpdateToken() gotInfo = %v, want %v", gotInfo, tt.wantInfo)
			}
		})
	}
}

package repo

import (
	"reflect"
	"testing"

	"github.com/blackhorseya/lobster/internal/pkg/base/contextx"
	"github.com/blackhorseya/lobster/internal/pkg/entity/user"
	"github.com/stretchr/testify/suite"
)

var (
	email1 = "email"

	pass1 = "$2a$04$W7XkHbwTrBUistouvflijuB2JOnYW4iEZEHVGgTX1bSERjRPZgZR."

	token1 = "token"

	info1 = &user.Profile{
		ID:       0,
		Email:    email1,
		Password: pass1,
		Token:    token1,
	}
)

type repoSuite struct {
	suite.Suite
	repo IRepo
}

func (s *repoSuite) SetupTest() {
	if repo, err := CreateRepo("../../../../../../configs/app.yaml"); err != nil {
		panic(err)
	} else {
		s.repo = repo
	}
}

func TestRepoSuite(t *testing.T) {
	suite.Run(t, new(repoSuite))
}

func (s *repoSuite) Test_impl_Register() {
	type args struct {
		newUser *user.Profile
	}
	tests := []struct {
		name     string
		args     args
		wantInfo *user.Profile
		wantErr  bool
	}{
		{
			name:     "register then user",
			args:     args{newUser: info1},
			wantInfo: info1,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			gotInfo, err := s.repo.Register(contextx.Background(), tt.args.newUser)
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

func (s *repoSuite) Test_impl_GetByID() {
	type args struct {
		id int64
	}
	tests := []struct {
		name     string
		args     args
		wantInfo *user.Profile
		wantErr  bool
	}{
		{
			name:     "get by id then user",
			args:     args{id: 0},
			wantInfo: info1,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			gotInfo, err := s.repo.GetByID(contextx.Background(), tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInfo, tt.wantInfo) {
				t.Errorf("GetByID() gotInfo = %v, want %v", gotInfo, tt.wantInfo)
			}
		})
	}
}

func (s *repoSuite) Test_impl_GetByToken() {
	type args struct {
		token string
	}
	tests := []struct {
		name     string
		args     args
		wantInfo *user.Profile
		wantErr  bool
	}{
		{
			name:     "get by token then user",
			args:     args{token: token1},
			wantInfo: info1,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			gotInfo, err := s.repo.GetByToken(contextx.Background(), tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInfo, tt.wantInfo) {
				t.Errorf("GetByToken() gotInfo = %v, want %v", gotInfo, tt.wantInfo)
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
			name:     "get by email then user",
			args:     args{email: email1},
			wantInfo: info1,
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
			name:     "update token then user",
			args:     args{info1},
			wantInfo: info1,
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

package kr

import "github.com/blackhorseya/lobster/internal/biz/kr/repo"

type impl struct {
	repo repo.IRepo
}

// NewImpl serve caller to create an IBiz
func NewImpl(repo repo.IRepo) IBiz {
	return &impl{repo: repo}
}

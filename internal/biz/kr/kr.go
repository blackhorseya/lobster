package kr

import (
	"github.com/blackhorseya/lobster/internal/biz/kr/repo"
	"github.com/google/wire"
)

// IBiz declare key result biz service function
type IBiz interface {
	// todo: 2021-01-25|10:07|doggy|implement me
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewImpl, repo.ProviderSet)

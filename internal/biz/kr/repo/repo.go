package repo

import "github.com/google/wire"

// IRepo declare key result repo service function
type IRepo interface {
	// todo: 2021-01-25|10:07|doggy|implement me
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewImpl)

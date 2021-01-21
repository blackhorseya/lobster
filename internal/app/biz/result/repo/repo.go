package repo

import "github.com/google/wire"

// IRepo declare key result repo service function
type IRepo interface {
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewImpl)

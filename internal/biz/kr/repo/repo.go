package repo

import (
	"github.com/blackhorseya/lobster/internal/pkg/contextx"
	"github.com/blackhorseya/lobster/internal/pkg/entities/biz/okr"
	"github.com/google/wire"
)

// IRepo declare key result repo service function
type IRepo interface {
	// QueryByID means query key result by key result's id
	QueryByID(ctx contextx.Contextx, id string) (kr *okr.KeyResult, err error)

	// QueryList means query key result list
	QueryList(ctx contextx.Contextx, offset, limit int) (krs []*okr.KeyResult, err error)

	// Delete means delete a key result
	Delete(ctx contextx.Contextx, id string) (err error)

	// Create means create a key result for objective
	Create(ctx contextx.Contextx, created *okr.KeyResult) (kr *okr.KeyResult, err error)

	// todo: 2021-01-25|10:07|doggy|implement me
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewImpl)

package repo

import (
	"github.com/blackhorseya/lobster/internal/pkg/contextx"
	"github.com/blackhorseya/lobster/internal/pkg/entities/biz/okr"
	"github.com/google/wire"
)

// IRepo declare key result repo service function
type IRepo interface {
	// QueryKRByID means query key result by goal's id and key result's id
	QueryKRByID(ctx contextx.Contextx, goalID, krID string) (*okr.KeyResult, error)

	// todo: 2021-01-25|10:07|doggy|implement me
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewImpl)

package repo

import (
	"github.com/blackhorseya/lobster/internal/pkg/contextx"
	"github.com/blackhorseya/lobster/internal/pkg/pb"
	"github.com/google/wire"
)

// IRepo declare key result repo service function
type IRepo interface {
	// QueryByID means query key result by key result's id
	QueryByID(ctx contextx.Contextx, id string) (kr *pb.Result, err error)

	// QueryByGoalID means query key results by goal's id
	QueryByGoalID(ctx contextx.Contextx, id string) (krs []*pb.Result, err error)

	// QueryList means query key result list
	QueryList(ctx contextx.Contextx, offset, limit int) (krs []*pb.Result, err error)

	// Delete means delete a key result
	Delete(ctx contextx.Contextx, id string) (err error)

	// Create means create a key result for objective
	Create(ctx contextx.Contextx, created *pb.Result) (kr *pb.Result, err error)

	// Update means update a key result
	Update(ctx contextx.Contextx, updated *pb.Result) (kr *pb.Result, err error)
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewImpl)
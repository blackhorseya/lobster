package repo

import (
	"github.com/blackhorseya/lobster/internal/pkg/contextx"
	"github.com/blackhorseya/lobster/internal/pkg/entities/biz/okr"
	"go.mongodb.org/mongo-driver/mongo"
)

type impl struct {
	mongo *mongo.Client
}

// NewImpl serve caller to create an IRepo
func NewImpl(mongo *mongo.Client) IRepo {
	return &impl{mongo: mongo}
}

func (i *impl) QueryByID(ctx contextx.Contextx, id string) (*okr.Objective, error) {
	panic("implement me")
}

func (i *impl) Create(ctx contextx.Contextx, created *okr.Objective) (*okr.Objective, error) {
	panic("implement me")
}

func (i *impl) List(ctx contextx.Contextx, offset, limit int) ([]*okr.Objective, error) {
	panic("implement me")
}

func (i *impl) Count(ctx contextx.Contextx) (int, error) {
	panic("implement me")
}

func (i *impl) Update(ctx contextx.Contextx, updated *okr.Objective) (*okr.Objective, error) {
	panic("implement me")
}

func (i *impl) Delete(ctx contextx.Contextx, id string) (int, error) {
	panic("implement me")
}

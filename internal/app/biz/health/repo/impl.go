package repo

import (
	"time"

	"github.com/blackhorseya/lobster/internal/pkg/contextx"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type impl struct {
	mongo *mongo.Client
}

// NewImpl serve caller to create an IRepo
func NewImpl(mongo *mongo.Client) IRepo {
	return &impl{mongo: mongo}
}

func (i *impl) Ping(ctx contextx.Contextx) (bool, error) {
	timeout, cancel := contextx.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	if err := i.mongo.Ping(timeout, readpref.Primary()); err != nil {
		return false, err
	}

	return true, nil
}

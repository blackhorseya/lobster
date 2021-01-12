package databases

import (
	"time"

	"github.com/blackhorseya/lobster/internal/pkg/config"
	"github.com/blackhorseya/lobster/internal/pkg/contextx"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// NewMongoDB init mongodb client
func NewMongoDB(cfg *config.Config) (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(cfg.DB.URL))
	if err != nil {
		return nil, err
	}

	ctx, cancel := contextx.WithTimeout(contextx.Background(), 5*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return client, nil
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewMongoDB)

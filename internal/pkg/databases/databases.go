package databases

import (
	"time"

	"github.com/blackhorseya/lobster/internal/pkg/config"
	"github.com/blackhorseya/lobster/internal/pkg/contextx"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	// import mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Options is configuration of database
type Options struct {
	URL   string `yaml:"url"`
	Debug bool
}

// NewOptions serve caller to create an Options
func NewOptions(v *viper.Viper, logger *zap.Logger) (*Options, error) {
	var err error
	o := new(Options)
	if err = v.UnmarshalKey("db", o); err != nil {
		return nil, errors.Wrap(err, "unmarshal db option error")
	}

	logger.Info("load database options success", zap.String("url", o.URL))

	return o, err
}

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

// NewMariaDB init mariadb client
func NewMariaDB(o *Options) (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", o.URL)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db, nil
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewOptions, NewMariaDB)

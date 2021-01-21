package repo

import "go.mongodb.org/mongo-driver/mongo"

type impl struct {
	mongo *mongo.Client
}

// NewImpl serve caller to create an IRepo
func NewImpl(mongo *mongo.Client) IRepo {
	return &impl{mongo: mongo}

}

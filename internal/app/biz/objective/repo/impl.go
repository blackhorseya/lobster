package repo

import (
	"time"

	"github.com/blackhorseya/lobster/internal/pkg/contextx"
	"github.com/blackhorseya/lobster/internal/pkg/entities/biz/okr"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type impl struct {
	mongo *mongo.Client
}

// NewImpl serve caller to create an IRepo
func NewImpl(mongo *mongo.Client) IRepo {
	return &impl{mongo: mongo}
}

func (i *impl) Create(ctx contextx.Contextx, created *okr.Objective) (*okr.Objective, error) {
	timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	coll := i.mongo.Database("lobster-db").Collection("objectives")
	_, err := coll.InsertOne(timeout, created)
	if err != nil {
		return nil, err
	}

	return created, nil
}

func (i *impl) QueryByID(ctx contextx.Contextx, id string) (*okr.Objective, error) {
	timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	coll := i.mongo.Database("lobster-db").Collection("objectives")
	res := coll.FindOne(timeout, bson.D{{"id", id}})
	if res.Err() != nil {
		return nil, res.Err()
	}

	var ret *okr.Objective
	if err := res.Decode(&ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func (i *impl) List(ctx contextx.Contextx, offset, limit int) ([]*okr.Objective, error) {
	timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	coll := i.mongo.Database("lobster-db").Collection("objectives")
	cur, err := coll.Find(timeout, bson.D{}, options.Find().SetSkip(int64(offset)).SetLimit(int64(limit)))
	if err != nil {
		return nil, err
	}
	defer cur.Close(timeout)

	var ret []*okr.Objective
	for cur.Next(timeout) {
		var ele *okr.Objective
		if err := cur.Decode(&ele); err != nil {
			return nil, err
		}

		ret = append(ret, ele)
	}

	if cur.Err() != nil {
		return nil, cur.Err()
	}

	return ret, nil
}

func (i *impl) Count(ctx contextx.Contextx) (int, error) {
	timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	coll := i.mongo.Database("lobster-db").Collection("objectives")
	ret, err := coll.CountDocuments(timeout, bson.D{})
	if err != nil {
		return 0, err
	}

	return int(ret), nil
}

func (i *impl) Update(ctx contextx.Contextx, updated *okr.Objective) (*okr.Objective, error) {
	timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	coll := i.mongo.Database("lobster-db").Collection("objectives")
	_, err := coll.UpdateOne(timeout, bson.D{{"id", updated.ID}}, bson.D{{"$set", updated}})
	if err != nil {
		return nil, err
	}

	return updated, nil
}

func (i *impl) Delete(ctx contextx.Contextx, id string) (int, error) {
	timeout, cancel := contextx.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	coll := i.mongo.Database("lobster-db").Collection("objectives")
	_, err := coll.DeleteOne(timeout, bson.D{{"id", id}})
	if err != nil {
		return 0, err
	}

	return 1, nil
}

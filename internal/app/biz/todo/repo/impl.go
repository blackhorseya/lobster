package repo

import (
	"github.com/blackhorseya/lobster/internal/pkg/contextx"
	"github.com/blackhorseya/lobster/internal/pkg/entities/biz/todo"
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

func (i *impl) QueryByID(ctx contextx.Contextx, id string) (*todo.Task, error) {
	coll := i.mongo.Database("lobster-db").Collection("tasks")
	find := coll.FindOne(ctx, bson.D{{"id", id}})
	if find.Err() != nil {
		return nil, find.Err()
	}

	var ret *todo.Task
	if err := find.Decode(&ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func (i *impl) Create(ctx contextx.Contextx, task *todo.Task) (*todo.Task, error) {
	coll := i.mongo.Database("lobster-db").Collection("tasks")
	if _, err := coll.InsertOne(ctx, task); err != nil {
		return nil, err
	}

	return task, nil
}

func (i *impl) List(ctx contextx.Contextx, offset, limit int) ([]*todo.Task, error) {
	coll := i.mongo.Database("lobster-db").Collection("tasks")
	cur, err := coll.Find(ctx, bson.D{}, options.Find().SetLimit(int64(limit)).SetSkip(int64(offset)))
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var ret []*todo.Task
	for cur.Next(ctx) {
		var task *todo.Task
		err := cur.Decode(&task)
		if err != nil {
			return nil, err
		}

		ret = append(ret, task)
	}

	if cur.Err() != nil {
		return nil, cur.Err()
	}

	return ret, nil
}

func (i *impl) Count(ctx contextx.Contextx) (int, error) {
	coll := i.mongo.Database("lobster-db").Collection("tasks")
	ret, err := coll.CountDocuments(ctx, bson.D{})
	if err != nil {
		return 0, err
	}

	return int(ret), nil
}

func (i *impl) Update(ctx contextx.Contextx, updated *todo.Task) (*todo.Task, error) {
	coll := i.mongo.Database("lobster-db").Collection("tasks")
	_, err := coll.UpdateOne(ctx, bson.D{{"id", updated.ID}}, bson.D{{"$set", updated}})
	if err != nil {
		return nil, err
	}

	return updated, nil
}

func (i *impl) Delete(ctx contextx.Contextx, id string) (int, error) {
	coll := i.mongo.Database("lobster-db").Collection("tasks")
	ret, err := coll.DeleteOne(ctx, bson.D{{"id", id}})
	if err != nil {
		return 0, err
	}

	return int(ret.DeletedCount), nil
}

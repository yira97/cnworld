package user

import (
	"context"

	"github.com/yira97/cnworld/app/blog/storage"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateTempUser(ctx context.Context, v TempUserView) (*TempUserView, error) {
	collection := storage.GetCollection(storage.CollectionNameTempUser)

	v.Init4Create()
	insertResult, err := collection.InsertOne(ctx, v)
	if err != nil {
		return nil, err
	}

	filter := bson.D{{Key: storage.CommonColumn__ID, Value: insertResult.InsertedID}}
	createdRecord := collection.FindOne(ctx, filter)

	// decode the Mongo record
	created := &storage.UserContactModel{}
	createdRecord.Decode(created)
	// overwrite
	v = TempUserView_From(created)
	return &v, nil
}

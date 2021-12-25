package message

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/yira97/cnworld/app/blog/storage"
	"github.com/yira97/cnworld/lib/base/short/s_mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateMessage(ctx context.Context, m storage.MessageModel) (*storage.MessageModel, error) {
	collection := storage.GetCNWorldInstance().DB.Collection(storage.CollectionNameMessage)

	now := time.Now()
	m.CreatedAt = now
	m.UID = uuid.New().String()

	insertionResult, err := collection.InsertOne(ctx, m)
	if err != nil {
		return nil, err
	}

	filter := bson.D{{Key: storage.CommonColumn__ID, Value: insertionResult.InsertedID}}
	createdRecord := collection.FindOne(ctx, filter)

	createdMessage := &storage.MessageModel{}
	createdRecord.Decode(createdMessage)

	return createdMessage, nil
}

func NewMetaMessage(ctx context.Context, messageUID string) storage.MetaMessageModel {
	now := time.Now()
	m := storage.MetaMessageModel{}
	m.UpdatedAt = now
	m.MessageUID = messageUID
	return m
}

func SetMessageHasRead(ctx context.Context, messageUID string) error {
	collection := storage.GetCNWorldInstance().DB.Collection(storage.CollectionNameMessage)

	filter := bson.M{
		storage.CommonColumn_MessageUID: messageUID,
	}

	findRes := collection.FindOne(ctx, filter)

	editingMetaMessage := storage.MetaMessageModel{}

	if findRes.Err() != nil {
		if findRes.Err() != mongo.ErrNoDocuments {
			return findRes.Err()
		}
		// create and insert
		editingMetaMessage = NewMetaMessage(ctx, messageUID)
		editingMetaMessage.HasRead = true
		_, err := collection.InsertOne(ctx, collection)
		if err != nil {
			return err
		}
	}

	opt := s_mongo.OptionFindOneAndUpdateThenReturn()

	update := bson.M{
		"$set": bson.M{
			storage.CommonColumn_HasRead: true,
		},
	}

	// Find one result and update it
	updateRes := collection.FindOneAndUpdate(ctx, filter, update, &opt)
	if updateRes.Err() != nil {
		return updateRes.Err()
	}

	return nil
}

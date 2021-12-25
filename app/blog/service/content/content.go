package content

import (
	"context"

	"github.com/yira97/cnworld/app/blog/storage"
	"github.com/yira97/cnworld/lib/base/short/s_mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ListContents(ctx context.Context, count int, offset int) ([]ContentView, error) {
	collection := storage.GetCollection(storage.CollectionNameContent)

	results := make([]storage.ContentModel, 0)

	opts := s_mongo.OptionFindDesc(count, offset, storage.CommonColumn_CreatedAt)

	views := make([]ContentView, 0)

	cursor, err := collection.Find(ctx, bson.D{}, opts)
	if err != nil {
		return views, err
	}

	if err := cursor.All(ctx, &results); err != nil {
		return views, err

	}
	for _, result := range results {
		views = append(views, ContentView_From(&result))
	}
	return views, nil
}

// 参数用的是下层storage的模型, 把本层的view模型, 返回给上层
func CreateContent(ctx context.Context, v ContentView) (*ContentView, error) {
	collection := storage.GetCollection(storage.CollectionNameContent)

	v.Init4Create()

	insertionResult, err := collection.InsertOne(ctx, v.Model())
	if err != nil {
		return nil, err
	}

	filter := bson.D{{Key: storage.CommonColumn__ID, Value: insertionResult.InsertedID}}
	createdRecord := collection.FindOne(ctx, filter)

	created := &storage.ContentModel{}
	createdRecord.Decode(created)

	// overwrite
	v = ContentView_From(created)

	return &v, nil
}

func GetContent(ctx context.Context, uID string) *ContentView {
	collection := storage.GetCollection(storage.CollectionNameContent)
	opt := options.FindOne()
	filter := bson.M{
		"uid": uID,
	}
	findResult := collection.FindOne(ctx, filter, opt)
	if findResult == nil {
		return nil
	}
	m := &storage.ContentModel{}
	findResult.Decode(m)
	v := ContentView_From(m)
	return &v
}

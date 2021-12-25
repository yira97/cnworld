package content

import (
	"context"
	"math"

	"github.com/yira97/cnworld/app/blog/storage"
	"github.com/yira97/cnworld/lib/base/short/s_mongo"
	"go.mongodb.org/mongo-driver/bson"
)

// count == -1 代表全部取出
func ListComment(ctx context.Context, contentUID string, count int, offset int) ([]storage.ContentCommentModel, error) {
	collection := storage.GetCollection(storage.CollectionNameComment)

	result := make([]storage.ContentCommentModel, 0)

	if count == -1 {
		// 应该足够
		count = math.MaxInt16
	}

	opts := s_mongo.OptionFindAsc(count, offset, storage.CommonColumn_CreatedAt)

	filter := bson.D{
		{Key: storage.CommonColumn_ContentUID, Value: contentUID},
	}

	cursor, err := collection.Find(ctx, filter, opts)
	if err != nil {
		return result, err
	}

	// TODO: 不一次提取全部, 而是提取摘要.
	if err := cursor.All(ctx, &result); err != nil {
		return result, err

	}
	return result, nil
}

func CreateComment(ctx context.Context, content storage.ContentCommentModel) (*storage.ContentCommentModel, error) {
	// TODO
	return nil, nil
}

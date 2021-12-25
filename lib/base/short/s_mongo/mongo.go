package s_mongo

import (
	"github.com/yira97/cnworld/lib/base/verify"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func OptionFindOneAndUpdateThenReturn() options.FindOneAndUpdateOptions {
	after := options.After
	upsert := false
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	return opt
}

// sortType: 1 ascending, -1 descending
func OptionFind(count int, offset int, sortBy string, sortType int) *options.FindOptions {
	opts := options.Find()
	opts.SetLimit(int64(count))
	opts.SetSkip(int64(offset))

	// https://stackoverflow.com/a/58223839
	opts.SetSort(bson.D{
		// https://stackoverflow.com/questions/54548441/composite-literal-uses-unkeyed-fields
		primitive.E{Key: sortBy, Value: sortType},
	})
	return opts
}

func OptionFindDesc(count int, offset int, sortBy string) *options.FindOptions {
	return OptionFind(count, offset, sortBy, -1)
}

func OptionFindAsc(count int, offset int, sortBy string) *options.FindOptions {
	return OptionFind(count, offset, sortBy, 1)
}

// 假设在根节点存在deleted_at, 就可以适用.
// 包含字段不存在, 以及字段存在但为null 两种情况
func DefaultDeletedOrderedFilter(fs ...bson.E) bson.D {
	arr := bson.D{
		{
			Key: "$or",
			Value: []bson.M{
				{
					verify.RawEntityDeletedAt: nil,
				},
				{
					verify.RawEntityDeletedAt: bson.M{"$exists": false},
				},
			},
		},
	}

	for _, f := range fs {
		arr = append(arr, f)
	}

	return arr
}

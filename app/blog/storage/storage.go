package storage

import (
	"github.com/yira97/cnworld/app/blog/config"
	"github.com/yira97/cnworld/app/blog/logger"
	"github.com/yira97/cnworld/lib/base/defawt"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	baseDb string
	mc     *mongo.Client
)

const (
	CollectionNameTempUser    = "temp_user"
	CollectionNameMessage     = "message"
	CollectionNameMessageMeta = "message_meta"
	CollectionNameContent     = "content"
	CollectionNameComment     = "comment"
)

func Setup() {
	logger.Log("正在加载 BLOG - Module Storage ...")

	cfg := config.GetMongoDB()
	defawt.InitMongoDB(cfg)
	baseDb = cfg.Database
}

func GetCollection(collection string) *mongo.Collection {
	return mc.Database(baseDb).Collection(collection)
}

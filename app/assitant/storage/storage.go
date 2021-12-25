package storage

import (
	"fmt"

	"github.com/yira97/cnworld/app/assitant/options"
	"github.com/yira97/cnworld/lib/base/defawt"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	// under BaseDatabase
	ApplicationCollection = "application"
	UserCollection        = "user"
)

var (
	mc     *mongo.Client
	baseDb string
)

func Setup() {
	fmt.Println("正在加载 BLOG - Module Storage ...")

	cfg := options.GetMongoDB()
	mc = defawt.InitMongoDB(cfg)
	baseDb = cfg.Database
}

func GetBaseDB() *mongo.Database {
	return mc.Database(baseDb)
}

package defawt

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB_Option struct {
	URI           string  `json:"uri"`
	Database      string  `json:"database"`
	LogCollection *string `json:"log_collection"`
}

const DEFAULT_MONGODB_LogCollection = "start_log"

func InitMongoDB(opts MongoDB_Option) *mongo.Client {
	logCollection := DEFAULT_MONGODB_LogCollection
	if opts.LogCollection != nil && *opts.LogCollection != "" {
		logCollection = *opts.LogCollection
	}

	if opts.URI == "" {
		panic("uri is null")
	}
	tryDuration := 10 * time.Second
	clientOptions := options.Client().
		ApplyURI(opts.URI)
	ctx, cancel := context.WithTimeout(context.Background(), tryDuration)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	startLogCollection := client.Database(opts.Database).Collection(logCollection)

	startRecord := struct {
		Time time.Time `bson:"time"`
	}{Time: time.Now()}

	_, err = startLogCollection.InsertOne(ctx, startRecord)
	if err != nil {
		panic(err)
	}

	return client
}

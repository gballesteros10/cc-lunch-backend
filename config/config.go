package config

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dbName            string        = "cclunch"
	dbHost            string        = "mongodb://localhost:27017"
	DBRequestDuration time.Duration = 10 * time.Second

	Port string = ":3001"
)

func GetMongoDB() (*mongo.Database, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), DBRequestDuration)
	defer cancelFunc()

	dbClient, err := mongo.Connect(ctx, options.Client().ApplyURI(dbHost))
	if err != nil {
		return nil, err
	}

	return dbClient.Database(dbName), nil
}

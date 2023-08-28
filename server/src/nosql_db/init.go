package nosql_db

import (
	"context"
	"fmt"
	"log"
	"server/src/logger"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const logFileName = "server.log"
const logPrefix = "nosql_db: "

const (
	db      = "sports-near-me"
	address = "mongodb://localhost:27017"
)

type Client struct {
	l         *log.Logger
	client    *mongo.Client
	issueColl *mongo.Collection
}

type ClientI interface {
}

func Initialize(ctx context.Context) ClientI {
	l := logger.InitLogger(logFileName, logPrefix, nil)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(address))
	if err != nil {
		panic(fmt.Errorf("failed to connect to mongo database. error: %v", err))
		return nil
	}

	issueColl := client.Database(db).Collection("issues")
	seedData(issueColl)

	return &Client{l, client, issueColl}
}

func seedData(issueColl *mongo.Collection) {}

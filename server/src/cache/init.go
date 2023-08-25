package cache

import (
	"context"
	"github.com/go-redis/redis/v9"
	"log"
	"server/src/env"
	"server/src/logger"
)

const (
	logFileName = "server.log"
	logPrefix   = "cache: "
	address     = "localhost:6379"
	apiTokenKey = "api-token"
)

type Client struct {
	l      *log.Logger
	client *redis.Client
}

type ClientI interface {
	GetSession(ctx context.Context, key string) (string, error)
	CreateSession(ctx context.Context, key, val string) error
	DeleteSession(ctx context.Context, key string) error
}

func Initialize() *Client {
	l := logger.InitLogger(logFileName, logPrefix, nil)
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	seedData(client)

	return &Client{l, client}
}

func seedData(client *redis.Client) {
	ctx := context.Background()
	err := client.Set(ctx, apiTokenKey, env.ApiToken, 0).Err()
	if err != nil {
		panic(err)
	}
}

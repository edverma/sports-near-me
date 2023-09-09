package jobs

import (
	"context"
	"log"
	"server/src/cache"
	"server/src/logger"
	"server/src/nosql_db"
	"server/src/sql_db"
	"sync"
)

type job struct {
	l           *log.Logger
	nosqlClient nosql_db.ClientI
	sqlClient   sql_db.ClientI
	cacheClient cache.ClientI
}

func RunAllJobs(ctx context.Context, once *sync.Once) {
	go RunSportsNearMeJob(ctx, once)
}

func newJob(ctx context.Context, logFileName, logPrefix string, once *sync.Once) *job {
	nosqlClient := nosql_db.Initialize(ctx)
	sqlClient := sql_db.Initialize(once)
	cacheClient := cache.Initialize()

	l := logger.InitLogger(logFileName, logPrefix, nil)
	return &job{l, nosqlClient, sqlClient, cacheClient}
}

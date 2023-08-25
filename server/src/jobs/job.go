package jobs

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"server/src/cache"
	"server/src/logger"
	"server/src/nosql_db"
	"server/src/sql_db"
)

type job struct {
	l           *log.Logger
	nosqlClient nosql_db.ClientI
	sqlClient   sql_db.ClientI
	cacheClient cache.ClientI
}

func RunAllJobs(ctx context.Context) {
	go RunTemplateJob(ctx)
}

func newJob(ctx context.Context, logFileName, logPrefix string) *job {
	nosqlClient := nosql_db.Initialize(ctx)
	sqlClient := sql_db.Initialize()
	cacheClient := cache.Initialize()

	l := logger.InitLogger(logFileName, logPrefix, gin.DefaultWriter)
	return &job{l, nosqlClient, sqlClient, cacheClient}
}

package api

import (
	"context"
	"fmt"
	"log"
	"server/src/cache"
	"server/src/logger"
	"server/src/nosql_db"
	"server/src/sql_db"
	"sync"

	"github.com/gin-gonic/gin"
)

const logFileName = "server.log"
const logPrefix = "server: "

type server struct {
	l           *log.Logger
	nosqlClient nosql_db.ClientI
	sqlClient   sql_db.ClientI
	cacheClient cache.ClientI
}

func Run(ctx context.Context, once *sync.Once) {
	nosqlClient := nosql_db.Initialize(ctx)
	sqlClient := sql_db.Initialize(once)
	cacheClient := cache.Initialize()

	l := logger.InitLogger(logFileName, logPrefix, gin.DefaultWriter)
	s := &server{l, nosqlClient, sqlClient, cacheClient}

	r := gin.Default()
	s.initializeRoutes(r)
	err := r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		panic(fmt.Errorf("gin error: %v", err))
	}
}

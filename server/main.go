package main

import (
	"context"
	"fmt"
	"os"
	"server/src/api"
	"server/src/jobs"
	"sync"
)

func main() {
	err := os.Setenv("TZ", "")
	if err != nil {
		panic(fmt.Sprintf("failed to set timezone. error: %v", err))
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var once sync.Once
	jobs.RunAllJobs(ctx, &once)
	api.Run(ctx, &once)
}

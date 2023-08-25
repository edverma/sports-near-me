package main

import (
	"context"
	"fmt"
	"os"
	"server/src/api"
	"server/src/jobs"
)

func main() {
	err := os.Setenv("TZ", "")
	if err != nil {
		panic(fmt.Sprintf("failed to set timezone. error: %v", err))
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	jobs.RunAllJobs(ctx)
	api.Run(ctx)
}

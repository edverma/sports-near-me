package jobs

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/go-co-op/gocron"
)

const logFileName = "sports-near-me_job.log"
const logPrefix = "sports-near-me_job: "
const pace = 1 * time.Second

func RunSportsNearMeJob(parentCtx context.Context, once *sync.Once) {
	ctx, cancel := context.WithCancel(parentCtx)
	defer cancel()
	j := newJob(ctx, logFileName, logPrefix, once)
	j.l.Printf("sports-near-me cronjob is setting up...\n")

	s := gocron.NewScheduler(time.UTC)
	_, err := s.StartAt(time.Now().Add(time.Second * 1)).Every(1).Hour().DoWithJobDetails(j.sportsNearMeJob)
	if err != nil {
		panic(fmt.Sprintf("failed to initialize sports-near-me job. error: %v", err))
	}
	s.StartBlocking()
}

func (j *job) sportsNearMeJob(cron gocron.Job) {
	j.l.Printf("running sports-near-me job....")

	resp, err := http.Get("https://www.milb.com/gwinnett/schedule/2023-09")
	if err != nil {
		j.l.Printf("failed to get HTTP. error: %v", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		j.l.Printf("failed to read all HTTP. error: %v", err)
	}
	j.l.Printf(string(body))
}

package jobs

import (
	"context"
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
)

const logFileName = "sports-near-me_job.log"
const logPrefix = "sports-near-me_job: "
const pace = 1 * time.Second

func RunSportsNearMeJob(parentCtx context.Context) {
	ctx, cancel := context.WithCancel(parentCtx)
	defer cancel()
	j := newJob(ctx, logFileName, logPrefix)
	j.l.Printf("sports-near-me cronjob is setting up...\n")

	s := gocron.NewScheduler(time.UTC)
	_, err := s.Every(1).Hour().StartAt(time.Now().Add(1 * time.Hour).Truncate(time.Hour)).DoWithJobDetails(j.SportsNearMeJob)
	if err != nil {
		panic(fmt.Sprintf("failed to initialize sports-near-me job. error: %v", err))
	}
	s.StartBlocking()
}

func (j *job) SportsNearMeJob(cron gocron.Job) {
	j.l.Printf("running sports-near-me job....")

	for i := 0; i < 100; i++ {
		j.l.Printf("iteration %d", i+1)

		select {
		case <-cron.Context().Done():
			j.l.Printf("Job canceled")
			return
		case <-time.After(pace):
		}
	}
	j.l.Printf("sports-near-me job complete")
}

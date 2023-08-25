package jobs

import (
	"context"
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
)

const logFileName = "template_job.log"
const logPrefix = "template_job: "
const pace = 1 * time.Second

func RunTemplateJob(parentCtx context.Context) {
	ctx, cancel := context.WithCancel(parentCtx)
	defer cancel()
	j := newJob(ctx, logFileName, logPrefix)
	j.l.Printf("template cronjob is setting up...\n")

	s := gocron.NewScheduler(time.UTC)
	_, err := s.Every(1).Hour().StartAt(time.Now().Add(1 * time.Hour).Truncate(time.Hour)).DoWithJobDetails(j.templateJob)
	if err != nil {
		panic(fmt.Sprintf("failed to initialize template job. error: %v", err))
	}
	s.StartBlocking()
}

func (j *job) templateJob(cron gocron.Job) {
	j.l.Printf("running template job....")

	for i := 0; i < 100; i++ {
		j.l.Printf("iteration %d", i+1)

		select {
		case <-cron.Context().Done():
			j.l.Printf("Job canceled")
			return
		case <-time.After(pace):
		}
	}
	j.l.Printf("template job complete")
}

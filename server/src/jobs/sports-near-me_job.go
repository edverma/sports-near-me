package jobs

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"server/src/sql_db"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/google/uuid"
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

type HomeTeamName struct {
	Name string `json:"name"`
}

type AwayTeamName struct {
	Name string `json:"name"`
}

type Location struct {
	Address1   string `json:"address1"`
	City       string `json:"city"`
	Country    string `json:"USA"`
	PostalCode string `json:"postalCode"`
	State      string `json:"state"`
}

type Venue struct {
	Location Location `json:"location"`
	Name     string   `json:"name"`
}

type Away struct {
	AwayTeamName AwayTeamName `json:"team"`
}

type Home struct {
	HomeTeamName HomeTeamName `json:"team"`
}

type Team struct {
	Away Away `json:"away"`
	Home Home `json:"home"`
}

type Games struct {
	Teams Team  `json:"teams"`
	Venue Venue `json:"venue"`
}

type Date struct {
	Date       string  `json:"date"`
	TotalGames int     `json:"totalGames"`
	Games      []Games `json:"games"`
}

type ScheduleResponse struct {
	Dates []Date `json:"dates"`
}

func parseDate(str string) []int {
	dateArr := strings.Split(str, "-")
	dateInt := []int{}
	for i := 0; i < 3; i++ {
		intArr, err := strconv.Atoi(dateArr[i])
		if err != nil {
			panic(err)
		}
		dateInt = append(dateInt, intArr)
	}
	return dateInt
}

func (j *job) sportsNearMeJob(cron gocron.Job) {
	j.l.Printf("running sports-near-me job....")

	resp, err := http.Get("https://statsapi.mlb.com/api/v1/schedule?lang=en&sportId=11,12,13,14,15,16,5442&hydrate=team(venue(timezone,location)),venue(timezone,location),game(seriesStatus,seriesSummary,tickets,promotions,sponsorships,content(summary,media(epg))),seriesStatus,seriesSummary,decisions,person,linescore,broadcasts(all)&season=2023&startDate=2023-07-01&endDate=2023-07-31&teamId=431&eventTypes=primary&scheduleTypes=games,events,xref")
	if err != nil {
		j.l.Printf("failed to get HTTP. error: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		j.l.Printf("failed to read all HTTP. error: %v", err)
	}

	var res ScheduleResponse
	jsonErr := json.Unmarshal(body, &res)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	dateString := (res.Dates[0].Date)

	t := time.Date(parseDate(dateString)[0], time.Month(parseDate(dateString)[1]), parseDate(dateString)[2], 0, 0, 0, 0, time.UTC)

	j.sqlClient.CreateGame(&sql_db.Game{
		Id:       uuid.NewString(),
		Date:     t,
		HomeTeam: res.Dates[0].Games[0].Teams.Home.HomeTeamName.Name,
		AwayTeam: res.Dates[0].Games[0].Teams.Away.AwayTeamName.Name,
		Venue:    res.Dates[0].Games[0].Venue.Name,
		Address:  res.Dates[0].Games[0].Venue.Location.Address1,
		State:    res.Dates[0].Games[0].Venue.Location.State,
		City:     res.Dates[0].Games[0].Venue.Location.City,
		Zipcode:  res.Dates[0].Games[0].Venue.Location.PostalCode,
	})
}

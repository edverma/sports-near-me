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

func parseDate(str string) time.Time {
	dateArr := strings.Split(str, "-")
	dateArrInt := []int{}
	for i := 0; i < len(dateArr); i++ {
		intArr, err := strconv.Atoi(dateArr[i])
		if err != nil {
			panic(err)
		}
		dateArrInt = append(dateArrInt, intArr)
	}
	t := time.Date(dateArrInt[0], time.Month(dateArrInt[1]), dateArrInt[2], 0, 0, 0, 0, time.UTC)
	return t
}

func (jb *job) sportsNearMeJob(cron gocron.Job) {
	jb.l.Printf("running sports-near-me job....")
	year := 2023
	month := 07
	for loop := 0; loop < 24; loop++ {
		month = month + 1
		if month == 13 {
			month = 1
			year = year + 1
		}
		monthstr := strconv.Itoa(month)
		yearstr := strconv.Itoa(year)
		resp, err := http.Get("https://statsapi.mlb.com/api/v1/schedule?lang=en&sportId=11,12,13,14,15,16,5442&hydrate=team(venue(timezone,location)),venue(timezone,location),game(seriesStatus,seriesSummary,tickets,promotions,sponsorships,content(summary,media(epg))),seriesStatus,seriesSummary,decisions,person,linescore,broadcasts(all)&season=" + yearstr + "&startDate=" + yearstr + "-" + monthstr + "-01&endDate=" + yearstr + "-" + monthstr + "-31&teamId=431&eventTypes=primary&scheduleTypes=games,events,xref")
		if err != nil {
			jb.l.Printf("failed to get HTTP. error: %v", err)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			jb.l.Printf("failed to read all HTTP. error: %v", err)
		}

		var res ScheduleResponse
		jsonErr := json.Unmarshal(body, &res)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}

		for i := range res.Dates {
			lenGames := len(res.Dates[i].Games)
			date := res.Dates[i].Date
			for j := 0; j < lenGames; j++ {
				game := res.Dates[i].Games[j]
				jb.sqlClient.CreateGame(&sql_db.Game{
					Id:       uuid.NewString(),
					Date:     parseDate(date),
					HomeTeam: game.Teams.Home.HomeTeamName.Name,
					AwayTeam: game.Teams.Away.AwayTeamName.Name,
					Venue:    game.Venue.Name,
					Address:  game.Venue.Location.Address1,
					State:    game.Venue.Location.State,
					City:     game.Venue.Location.City,
					Zipcode:  game.Venue.Location.PostalCode,
				})
			}
		}
	}
	//startDate should be what user inputs end date should be 24 months afte
}

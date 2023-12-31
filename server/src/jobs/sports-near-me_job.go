package jobs

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
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
const numMonthsToSetScheduleInfo = 24

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
	Teams  Team   `json:"teams"`
	Venue  Venue  `json:"venue"`
	GameId string `json:"gameGuid"`
}

type Date struct {
	Date       string  `json:"date"`
	TotalGames int     `json:"totalGames"`
	Games      []Games `json:"games"`
}

type ScheduleResponse struct {
	Dates []Date `json:"dates"`
}

func (jb *job) sportsNearMeJob(cron gocron.Job) {
	jb.l.Printf("running sports-near-me job....")
	now := time.Now()
	year := now.Year()
	month := int(now.Month())
	for loop := 0; loop < numMonthsToSetScheduleInfo; loop++ {
		month, year = incrementMonth(month, year)
		daysInMonth := daysIn(month, year)
		for dayLoop := 1; dayLoop <= daysInMonth; dayLoop++ {
			url := createRequestUrl(month, year, dayLoop)
			resp, err := http.Get(url)
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
				jb.l.Print(jsonErr)
			}

			jb.insertGamesdb(res)
		}
	}

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

func daysIn(month, year int) int {
	return time.Date(year, time.Month(month)+1, 0, 0, 0, 0, 0, time.UTC).Day()
}

func createRequestUrl(month, year, day int) string {
	monthstr := strconv.Itoa(month)
	yearstr := strconv.Itoa(year)
	dayStr := strconv.Itoa(day)
	urlStr := "https://bdfed.stitch.mlbinfra.com/bdfed/transform-milb-schedule?stitch_env=prod&sortTemplate=5&sportId=11&&sportId=12&&sportId=13&&sportId=14&&sportId=16&startDate=" + yearstr + "-" + monthstr + "-" + dayStr + "&endDate=" + yearstr + "-" + monthstr + "-" + dayStr + "&gameType=E&&gameType=S&&gameType=R&&gameType=F&&gameType=D&&gameType=L&&gameType=W&&gameType=A&&gameType=C&language=en&leagueId=&contextTeamId=milb&teamId=&orgId="
	return urlStr
}

func (jb *job) insertGamesdb(res ScheduleResponse) {
	for i := range res.Dates {
		lenGames := len(res.Dates[i].Games)
		date := res.Dates[i].Date
		for j := 0; j < lenGames; j++ {
			game := res.Dates[i].Games[j]
			jb.sqlClient.CreateGame(&sql_db.Game{
				Id:       uuid.NewString(),
				GameId:   game.GameId,
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

func incrementMonth(month, year int) (int, int) {
	if month == 13 {
		month = 1
		year = year + 1
	}
	month++
	return month, year
}

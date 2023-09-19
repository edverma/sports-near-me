package sql_db

import (
	"fmt"
	"log"
	"server/src/logger"
	"sync"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const logFileName = "server.log"
const logPrefix = "sql_db: "

var (
	db = "sports-near-me.db"
)

type Client struct {
	l      *log.Logger
	client *gorm.DB
}

type ClientI interface {
	CreateCredential(tx *gorm.DB, credential *Credential) (credentialId string, err error)
	GetCredential(credential *Credential) (*Credential, error)
	CreateUser(tx *gorm.DB, user *User) error
	GetUser(user *User) (*User, error)
	CreateUserCredential(userCredential *UserCredential) error
	CreateGame(game *Game) error
	GetGame(game *Game) (*Game, error)
}

func Initialize(once *sync.Once) *Client {
	l := logger.InitLogger(logFileName, logPrefix, nil)
	client, err := gorm.Open(sqlite.Open(db), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("failed to connect to database. error: %v", err))
	}

	once.Do(func() {
		err = client.AutoMigrate(&Credential{}, &User{}, &Game{})
		if err != nil {
			panic(fmt.Errorf("automigrate failed. error: %v", err))
		}
	})

	seedData(client)

	return &Client{l, client}
}

func seedData(dbConn *gorm.DB) {
	credentials := setupCredentials()
	users := setupUsers()
	game := setupGame()

	err := dbConn.Clauses(clause.OnConflict{DoNothing: true}).Create(credentials).Error
	if err != nil {
		panic(err)
	}
	err = dbConn.Clauses(clause.OnConflict{DoNothing: true}).Create(users).Error
	if err != nil {
		panic(err)
	}
	err = dbConn.Clauses(clause.OnConflict{DoNothing: true}).Create(game).Error
	if err != nil {
		panic(err)
	}
}

func RunSportsNearMeJob() {
	panic("unimplemented")
}

//func sportsNearMeJob() {
//	panic("unimplemented")
//}

func setupCredentials() []*Credential {
	return []*Credential{
		{
			Id:       "1",
			Username: "test-user-1",
			Hash:     "test-user-1-hash",
		},
		{
			Id:       "2",
			Username: "test-user-2",
			Hash:     "test-user-2-hash",
		},
	}
}

func setupUsers() []*User {
	return []*User{
		{
			Id:           "1",
			CredentialId: "1",
			Username:     "test-user-1",
			Email:        "test-user-1@test.com",
			FirstName:    "First1",
			LastName:     "Last1",
		},
		{
			Id:           "2",
			CredentialId: "2",
			Username:     "test-user-2",
			Email:        "test-user-2@test.com",
			FirstName:    "First2",
			LastName:     "Last2",
		},
	}
}

func setupGame() []*Game {
	return []*Game{
		{
			Id:       "1",
			Date:     time.Now(),
			HomeTeam: "Gwinnett Stripers",
			AwayTeam: "Beep Boops",
			Venue:    "Truist Park",
			Address:  "360 N. Broadway Street St. Paul",
			State:    "Georgia",
			City:     "Atlanta",
			Zipcode:  "30314",
		},
	}
}

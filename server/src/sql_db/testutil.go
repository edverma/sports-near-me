package sql_db

import (
	"fmt"
	"log"
	"os"
	"server/src/logger"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	logFileName_test = "../../server-test.log"
	logPrefix_test   = "../../sql_db-test: "
	db_test          = "../../sports-near-me-test.db"
)

var test_client *Client

// TODO: get this working
func TestMain(m *testing.M) {
	test_client = initTestDb()
	code := m.Run()
	teardownTestDb()
	os.Exit(code)
}

func initTestDb() *Client {
	l := logger.InitLogger(logFileName_test, logPrefix_test, nil)
	client, err := gorm.Open(sqlite.Open(db_test), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("failed to connect to database. error: %v", err))
	}

	err = client.AutoMigrate(&Credential{}, &User{}, &Game{})
	if err != nil {
		panic(fmt.Errorf("automigrate failed. error: %v", err))
	}

	seedData(client)
	return &Client{l, client}
}

func teardownTestDb() {
	e := os.Remove(db_test)
	if e != nil {
		log.Fatal(e)
	}
}

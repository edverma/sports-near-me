package sql_db

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateGame(t *testing.T) {
	// setup
	test_client = initTestDb()
	game := setupGame()
	// run test

	err := test_client.CreateGame(game[0])
	if err != nil {
		t.Fatal(err)
	}

	gameToQuery := *game[0]
	gameToQuery.CreatedAt = time.Time{}
	gameToQuery.UpdatedAt = time.Time{}
	gamePrint, err := test_client.GetGame(&gameToQuery)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(gamePrint)

	err = test_client.CreateGame(&Game{GameId: "1", HomeTeam: "test"})
	if err != nil {
		t.Fatal(err)
	}
	updatedGamePrint, err := test_client.GetGame(&Game{Id: game[0].Id, GameId: "1", HomeTeam: "test"})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(updatedGamePrint)

	// teardown
	teardownTestDb()
}

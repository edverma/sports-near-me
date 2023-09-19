package sql_db

import (
	"time"

	"github.com/google/uuid"
)

type Game struct {
	Id       string    `gorm:"primaryKey" json:"id"`
	Date     time.Time `gorm:"unique" json:"date"`
	HomeTeam string    `json:"home_team"`
	AwayTeam string    `json:"away_team"`
	Venue    string    `json:"venue"`
	Address  string    `json:"address"`
	State    string    `json:"state"`
	City     string    `json:"city"`
	Zipcode  string    `json:"postalCode"`
}

func (c *Client) CreateGame(game *Game) error {
	game.Id = uuid.NewString()
	err := c.client.Create(game).Error
	if err != nil {
		c.l.Printf("failed to create new row in games table. error: %v", err)
		return err
	}
	return nil
}

func (c *Client) GetGame(game *Game) (*Game, error) {
	var dbGame Game
	err := c.client.Where(game).First(&dbGame).Error
	if err != nil {
		c.l.Printf("failed to get game. error: %v", err)
		return nil, err
	}
	return &dbGame, nil
}

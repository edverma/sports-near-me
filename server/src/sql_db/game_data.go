package sql_db

import (
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Game struct {
	Id        string         `gorm:"primaryKey" json:"NotGameid"`
	GameId    string         `gorm:"unique" json:"gameGuid"`
	Date      time.Time      `json:"date"`
	HomeTeam  string         `json:"home_team"`
	AwayTeam  string         `json:"away_team"`
	Venue     string         `json:"venue"`
	Address   string         `json:"address"`
	State     string         `json:"state"`
	City      string         `json:"city"`
	Zipcode   string         `json:"postalCode"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (c *Client) CreateGame(game *Game) error {
	err := c.client.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "game_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"date", "home_team", "away_team", "venue", "address", "state", "city", "zipcode", "updated_at"}),
	}).Create(game).Error
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

package sql_db

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Game struct {
	Id       string    `gorm:"primaryKey" json:"id"`
	Date     time.Time `gorm:"unique" json:"date"`
	HomeTeam string    `gorm:"unique" json:"home_team"`
	AwayTeam string    `gorm:"unique" json:"away_team"`
	Venue    string    `gorm:"unique" json:"venue"`
	Address  string    `gorm:"unique" json:"address"`
}

func (c *Client) CreateGameData(tx *gorm.DB, game_data *Game) error {
	game_data.Id = uuid.NewString()
	err := tx.Create(game_data).Error
	if err != nil {
		c.l.Printf("failed to create new row in users games table. error: %v", err)
		return err
	}
	return nil
}

func (c *Client) GetGame(game_data *Game) (*Game, error) {
	var dbGameData Game
	err := c.client.Where(game_data).First(&dbGameData).Error
	if err != nil {
		c.l.Printf("failed to get user. error: %v", err)
		return nil, err
	}
	return &dbGameData, nil
}

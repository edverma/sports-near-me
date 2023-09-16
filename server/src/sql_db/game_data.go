package sql_db

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Game struct {
	Id       string    `gorm:"primaryKey" json:"id"`
	Date     time.Time `gorm:"unique" json:"date"`
	HomeTeam string    `gorm:"column: HomeTeam" json:"home_team"`
	AwayTeam string    `gorm:"column: AwayTeam" json:"away_team"`
	Venue    string    `gorm:"column: Venue" json:"venue"`
	Address  string    `gorm:"column: Address" json:"address"`
}

func (c *Client) CreateGame(tx *gorm.DB, game *Game) error {
	game.Id = uuid.NewString()
	err := tx.Create(game).Error
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

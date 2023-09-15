package sql_db

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GameData struct {
	Game_Id  string `gorm:"primaryKey" json:"id"`
	Date     string `gorm:"unique" json:"date"`
	HomeTeam string `json:"home_team"`
	AwayTeam string `json:"away_team"`
	Venue    string `json:"venue"`
	Address  string `json:"address"`
}

func (c *Client) CreateGameData(tx *gorm.DB, game_data *GameData) error {
	game_data.Game_Id = uuid.NewString()
	err := tx.Create(game_data).Error
	if err != nil {
		c.l.Printf("failed to create new row in users table. error: %v", err)
		return err
	}
	return nil
}

func (c *Client) GetGameData(game_data *GameData) (*GameData, error) {
	var dbGameData GameData
	err := c.client.Where(game_data).First(&dbGameData).Error
	if err != nil {
		c.l.Printf("failed to get user. error: %v", err)
		return nil, err
	}
	return &dbGameData, nil
}

package sql_db

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id           string         `gorm:"primaryKey" json:"id"`
	CredentialId string         `gorm:"foreignKey;unique" json:"credential_id"`
	Email        string         `gorm:"unique" json:"email"`
	Username     string         `gorm:"unique" json:"username"`
	FirstName    string         `json:"first_name"`
	LastName     string         `json:"last_name"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (c *Client) CreateUser(tx *gorm.DB, user *User) error {
	user.Id = uuid.NewString()
	err := tx.Create(user).Error
	if err != nil {
		c.l.Printf("failed to create new row in users table. error: %v", err)
		return err
	}
	return nil
}

func (c *Client) GetUser(user *User) (*User, error) {
	var dbUser User
	err := c.client.Where(user).First(&dbUser).Error
	if err != nil {
		c.l.Printf("failed to get user. error: %v", err)
		return nil, err
	}
	return &dbUser, nil
}

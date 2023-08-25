package sql_db

import (
	"gorm.io/gorm"
	"time"
)

// UserCredential is not a gorm table
type UserCredential struct {
	User       `json:"user"`
	Credential `json:"credential"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (c *Client) CreateUserCredential(userCredential *UserCredential) error {
	err := c.client.Transaction(func(tx *gorm.DB) error {
		credentialId, err := c.CreateCredential(tx, &userCredential.Credential)
		if err != nil {
			return err
		}
		userCredential.User.CredentialId = credentialId
		err = c.CreateUser(tx, &userCredential.User)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		c.l.Printf("gorm transaction failed. error: %v", err)
		return err
	}
	return nil
}

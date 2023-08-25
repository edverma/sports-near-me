package sql_db

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type Credential struct {
	Id        string         `gorm:"primaryKey" json:"id"`
	Username  string         `gorm:"unique" json:"username"`
	Hash      string         `json:"hash"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (c *Client) CreateCredential(tx *gorm.DB, credential *Credential) (credentialId string, err error) {
	var hash []byte
	hash, err = bcrypt.GenerateFromPassword([]byte(credential.Hash), bcrypt.DefaultCost)
	if err != nil {
		c.l.Printf("failed to generate hash from password. error: %v", err)
		return "", err
	}
	credential.Hash = string(hash)
	credential.Id = uuid.NewString()
	err = tx.Create(credential).Error
	if err != nil {
		c.l.Printf("failed to create new row in credentials table. error: %v", err)
		return "", err
	}
	return credential.Id, nil
}

func (c *Client) GetCredential(credential *Credential) (*Credential, error) {
	var dbCredential Credential
	err := c.client.Where("username = ?", credential.Username).First(&dbCredential).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			c.l.Printf("error retrieving credential from username. error: %v", err)
		}
		return nil, err
	}
	return &dbCredential, nil
}

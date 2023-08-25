package cache

import (
	"context"
	"time"
)

const sessionTokenExpirationSeconds = 2592000 // 30 days

func (c *Client) GetSession(ctx context.Context, key string) (string, error) {
	session, err := c.client.Get(ctx, key).Result()
	if err != nil {
		c.l.Printf("failed to get session token from cache. error: %v", err)
		return "", err
	}
	return session, nil
}

func (c *Client) CreateSession(ctx context.Context, key, val string) error {
	err := c.client.Set(ctx, key, val, sessionTokenExpirationSeconds*time.Second).Err()
	if err != nil {
		c.l.Printf("failed to set session token in cache. error: %v", err)
		return err
	}
	return nil
}

func (c *Client) DeleteSession(ctx context.Context, key string) error {
	err := c.client.Del(ctx, key).Err()
	if err != nil {
		c.l.Printf("failed to delete session token. error: %v", err)
		return err
	}
	return nil
}

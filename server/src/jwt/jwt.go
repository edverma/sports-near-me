package jwt

import (
	"server/src/env"
	"server/src/sql_db"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

const (
	sessionTokenExpirationSeconds = 2592000 // 30 days
	appName                       = "sports-near-me"
)

type Claims struct {
	User sql_db.User `json:"user"`
	jwt.StandardClaims
}

func CreateSessionToken(user sql_db.User) (string, error) {
	sessionId := uuid.NewString()
	token := jwt.NewWithClaims(env.SigningMethod, &Claims{
		user,
		jwt.StandardClaims{
			Id:        sessionId,
			Subject:   user.Id,
			ExpiresAt: time.Now().Add(sessionTokenExpirationSeconds * time.Second).Unix(),
			IssuedAt:  time.Now().Unix(),
			NotBefore: time.Now().Unix(),
			Issuer:    appName,
			Audience:  appName,
		},
	})
	sessionToken, err := token.SignedString(env.JwtSecret)
	if err != nil {
		return "", err
	}
	return sessionToken, nil
}

func ParseSessionToken(sessionToken string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(sessionToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return env.JwtSecret, nil
	})
	if err != nil {
		return nil, err
	}

	claims := token.Claims.(*Claims)
	return claims, nil
}

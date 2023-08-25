package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"server/src/env"
	"server/src/jwt"
	"server/src/sql_db"
)

func (s *server) validateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokens := c.Request.Header["Api-Token"]

		if len(tokens) == 0 {
			s.l.Printf("request failed because authorization token not provided")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
			return
		}
		if tokens[0] != env.ApiToken {
			s.l.Printf("request failed because authorization token invalid\n%s\n%s", tokens[0], env.ApiToken)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
			return
		}
		c.Next()
	}
}

func (s *server) validateSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokens := c.Request.Header["Session-Token"]

		if len(tokens) == 0 || tokens[0] == "" {
			s.l.Printf("request failed because authorization token not provided")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
			return
		}
		requestSessionToken := tokens[0]

		claims, err := jwt.ParseSessionToken(requestSessionToken)
		if err != nil {
			s.l.Printf("failed to parse jwt. error: %v", err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		cacheSessionToken, err := s.cacheClient.GetSession(c, claims.Subject)
		if err != nil {
			s.l.Printf("failed to get session token from cache. error: %v", err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		cacheClaims, err := jwt.ParseSessionToken(cacheSessionToken)
		if err != nil {
			s.l.Printf("failed to parse localjwt. error: %v", err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if cacheClaims.Id != claims.Id {
			s.l.Printf("sent session token does not equal session token in cache.")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		err = setUserOnContext(c, &cacheClaims.User)
		if err != nil {
			s.l.Printf("failed to set user onto context. error: %v", err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		c.Next()
	}
}

func setUserOnContext(c *gin.Context, user *sql_db.User) error {
	encodedUser, err := json.Marshal(user)
	if err != nil {
		return err
	}

	c.Set("user", encodedUser)
	return nil
}

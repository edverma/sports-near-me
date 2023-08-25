package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/src/jwt"
)

const sessionTokenHeader = "Session-Token"

func (s *server) logoutHandler(c *gin.Context) {
	if len(c.Request.Header[sessionTokenHeader]) == 0 {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	sessionToken := c.Request.Header["Session-Token"][0]
	claims, err := jwt.ParseSessionToken(sessionToken)
	if err != nil {
		s.l.Printf("failed to parse jwt session token. error: %v", err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	err = s.cacheClient.DeleteSession(c, claims.Subject)
	if err != nil {
		s.l.Printf("failed to delete session from cache. error: %v", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}

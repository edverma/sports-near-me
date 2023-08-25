package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/src/jwt"
)

func (s *server) getUserHandler(c *gin.Context) {
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

	c.JSONP(http.StatusOK, claims.User)
}

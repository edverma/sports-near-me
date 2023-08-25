package api

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"server/src/jwt"
	"server/src/sql_db"
)

func (s *server) loginHandler(c *gin.Context) {
	var req sql_db.Credential
	err := c.BindJSON(&req)
	if err != nil {
		s.l.Printf("failed to bind request to model. error: %v", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if !s.loginRequestIsValid(&req) {
		s.l.Printf("request is invalid. req: %v", req)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	credential, err := s.sqlClient.GetCredential(&req)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.AbortWithStatus(http.StatusUnauthorized)
		} else {
			s.l.Printf("database read to credential table failed. error: %v", err)
			c.AbortWithStatus(http.StatusInternalServerError)
		}
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(credential.Hash), []byte(req.Hash))
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	user, err := s.sqlClient.GetUser(&sql_db.User{CredentialId: credential.Id})
	if err != nil {
		s.l.Printf("database read to user table failed. error: %v", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	sessionToken, err := jwt.CreateSessionToken(*user)
	if err != nil {
		s.l.Printf("failed to created jwt session token. error: %v", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	err = s.cacheClient.CreateSession(c, user.Id, sessionToken)
	if err != nil {
		s.l.Printf("failed to set session token in cache. error: %v", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{"session_token": sessionToken})
}

func (s *server) loginRequestIsValid(req *sql_db.Credential) bool {
	if req.Username == "" {
		s.l.Println("username in request is empty")
		return false
	}
	if req.Hash == "" {
		s.l.Println("password in request is empty")
		return false
	}
	return true
}

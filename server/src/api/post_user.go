package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/src/sql_db"
)

func (s *server) postUserHandler(c *gin.Context) {
	var req sql_db.UserCredential
	err := c.BindJSON(&req)
	if err != nil {
		s.l.Printf("failed to bind request to model. error: %v", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if !s.createUserRequestIsValid(&req) {
		s.l.Printf("request is invalid. request: %v", req)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = s.sqlClient.CreateUserCredential(&req)
	if err != nil {
		s.l.Printf("failed to create user credential in database. error: %v", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}

func (s *server) createUserRequestIsValid(req *sql_db.UserCredential) bool {
	if req.Credential.Username == "" {
		s.l.Printf("username in request is empty")
		return false
	}
	if req.Credential.Hash == "" {
		s.l.Printf("hash in request is empty")
		return false
	}
	if req.User.Email == "" {
		s.l.Printf("email in request is empty")
		return false
	}
	if req.User.FirstName == "" {
		s.l.Printf("first name in request is empty")
		return false
	}
	if req.User.LastName == "" {
		s.l.Printf("last name in request is empty")
		return false
	}
	return true
}

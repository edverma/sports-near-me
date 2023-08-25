package api

import (
	"github.com/gin-gonic/gin"
)

func (s *server) initializeRoutes(r *gin.Engine) {
	authProtect := r.Group("")
	noProtect := r.Group("")

	authProtect.Use(s.validateSession())

	noProtect.GET("/ping", s.pingHandler)

	noProtect.POST("/login", s.loginHandler)
	authProtect.POST("/logout", s.logoutHandler)

	noProtect.POST("/user", s.postUserHandler)
	authProtect.GET("/user", s.getUserHandler)
}

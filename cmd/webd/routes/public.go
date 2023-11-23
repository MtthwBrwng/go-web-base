package routes

import (
	"github.com/gin-gonic/gin"
)

func (s *Service) LoadPublicRoutes(router *gin.Engine) {
	/*
		Home
	*/

	router.GET("", s.TodoHandler.PublicIndexPageHandler)
}

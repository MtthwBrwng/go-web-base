package routes

import (
	"github.com/gin-gonic/gin"
)

func (s *Service) LoadAppRoutes(router *gin.Engine) {
	app := router.Group("/app")

	/*
		Home
	*/

	app.GET("", s.TodoHandler.AppIndexPageHandler)

	app.GET("/todos", s.TodoHandler.AppIndexPageHandler)
	app.GET("/todos/new", s.TodoHandler.AppNewPageHandler)

	app.POST("/todos/new", s.TodoHandler.AppNewActionHandler)

	todo := app.Group("/todos/:id")
	todo.Use(s.TodoHandler.FetcherMiddleware)

	todo.DELETE("", s.TodoHandler.AppDeleteActionHandler)
}

package routes

import (
	middlewares "github.com/event-backend-api/midlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.POST("/login", login)
	//USERS ROUTES
	server.POST("/signup", signUp)

	//EVENTS ROUTES
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEventsById)
	protected := server.Group("/")
	protected.Use(middlewares.Authenticate)
	{
		server.POST("/events", middlewares.Authenticate, createEvent)
		server.PUT("/events/:id", updateEvent)
		server.DELETE("/events/:id", deleteEvent)
	}
}

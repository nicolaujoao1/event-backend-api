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

	server.POST("/events", middlewares.Authenticate, createEvent)
	server.PUT("/events/:id", middlewares.Authenticate, updateEvent)
	server.DELETE("/events/:id", middlewares.Authenticate, deleteEvent)
	server.POST("/events/:id/register", middlewares.Authenticate, registerForEvent)
	server.DELETE("/events/:id/register", middlewares.Authenticate, cancelRegistration)

}

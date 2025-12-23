package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	//EVENTS ROUTES
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEventsById)
	server.POST("/events", createEvent)
	server.PUT("/events/:id", updateEvent)
	server.DELETE("/events/:id", deleteEvent)

	//USERS ROUTES
	server.POST("/signup", signUp)
	server.POST("/login", login)
}

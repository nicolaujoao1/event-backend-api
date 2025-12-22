package main

import (
	"github.com/event-backend-api/db"
	"github.com/event-backend-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080") //localhost:8080
}

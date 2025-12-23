package routes

import (
	"net/http"

	"github.com/event-backend-api/models"
	"github.com/gin-gonic/gin"
)

func signUp(context *gin.Context) {
	var user models.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := user.Save(); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{
		"message": "user created successfully",
	})
}

func login(context *gin.Context) {
	var credentials models.User
	if err := context.ShouldBindJSON(&credentials); err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := credentials.ValidateCredentials(); err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "login successful",
	})

}

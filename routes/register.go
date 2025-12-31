package routes

import (
	"net/http"
	"strconv"

	"github.com/event-backend-api/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	userId := ctx.GetInt64("userId")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	event, err := models.GetEventById(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event!!"})
		return
	}

	if err = event.Register(userId); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "created succesfully",
	})

}

func cancelRegistration(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	userId := ctx.GetInt64("userId")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var event models.Event
	event.ID = id

	if err = event.CancelRegistration(userId); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Registration canceled!",
	})

}

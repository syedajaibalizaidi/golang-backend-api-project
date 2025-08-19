package routes

import (
	"net/http"
	"rest-api/m/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func registerEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64) // to get a path parameter. id is set to int64 so we need to convert it to that by using str.conv which used to convert the val.
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "parsing event id failed! "})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "failed to fetch event."})
		return
	}

	// newly created registered table
	err = event.Register(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "user registration failed."}) // map gin.H
		return
	}
	// success respone
	context.JSON(http.StatusCreated, gin.H{"message": "user registered successfully!"})
}

func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	var event models.Event
	event.ID = eventId

	err = event.CancelRegistration(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "user cancelling failed."}) // map gin.H
		return
	}
	// success respone
	context.JSON(http.StatusCreated, gin.H{"message": "Cancelled"})
}

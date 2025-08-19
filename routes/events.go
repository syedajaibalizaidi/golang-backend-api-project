// file responsible for containing all request handler functions that deal with events.
package routes

import (
	"net/http"
	"rest-api/m/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// request handlers
func getEvents(context *gin.Context) {
	events, err := models.GetAllWEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "events fetching failed try again bro!"})
		return
	}
	context.JSON(http.StatusOK, events) // writing the status code here. then passing the obj with key value pairs. map
}

// getting single event.
func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64) // to get a path parameter. id is set to int64 so we need to convert it to that by using str.conv which used to convert the val.
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "parsing event id failed! "})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "event fetching failed."})
		return
	}
	// success response
	context.JSON(http.StatusOK, event)
}

func createEvents(context *gin.Context) {
	// checking whether it contains a valid token extract token from incoming request

	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "sorry parsing problem!"})
		return
	}

	// retrieving userid from context
	userId := context.GetInt64("userId") // gives us the val converted to right type. using the same key that we set with the set method in the auth.go

	// event.ID = 1
	event.UserID = userId

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "error in creating events. try again bro!"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "event created!", "event": event})
}

// for update

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64) // to get a path parameter. id is set to int64 so we need to convert it to that by using str.conv which used to convert the val.
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "parsing event id failed! "})
		return
	}

	userId := context.GetInt64("userId")
	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch the event."})
		return
	}

	// logic ->> only the authorized users can update that event not random user.
	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized work."})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not fetch the event."})
		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not update the event."})
		return
	}
	// for success
	context.JSON(http.StatusOK, gin.H{"message": "success"})
}

// for deletion.
func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64) // to get a path parameter. id is set to int64 so we need to convert it to that by using str.conv which used to convert the val.
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "parsing event id failed! "})
		return
	}

	userId := context.GetInt64("userId")
	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch the event."})
		return
	}

	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized to delete the event."})
		return
	}

	err = event.Delete()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not delete the event."})
		return
	}

	// for success
	context.JSON(http.StatusOK, gin.H{"message": "deletion success"})

}

// for routes responsible for registering these routes in our app.

package routes

import (
	"rest-api/m/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)    // relative path and handlers
	server.GET("/events/:id", getEvent) // will be unique val dynamic path variable 1/2/3/4/5/ by using gin we can define that path by adding a :

	// Grouping all of them
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate) // to protect them
	authenticated.POST("/events", createEvents)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerEvent) // routes for registration tables
	authenticated.DELETE("/events/:id/register", cancelRegistration)

	// server.POST("/events", middlewares.Authenticate, createEvents) // post method ->> HTTP METHOD >>> GET, POST, PUT, PATCH, DELETE
	// server.PUT("/events/:id", updateEvent)                         // HTTP METHOD FOR UPDATEING THE DATA. dynamic path parameter used.
	// server.DELETE("/events/:id", deleteEvent)                      // HTTP METHOD FOR POST DELETION
	server.POST("/signup", signup) // logic for user signup db.
	server.POST("/login", login)   // logic for login
}

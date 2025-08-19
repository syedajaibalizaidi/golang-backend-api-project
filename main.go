// imported gin library for req methods which is specially for golang developers.

package main

import (
	"rest-api/m/db"
	"rest-api/m/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server) // passing pointer to register route
	server.Run(":8000")           // localhost address as string. we call these endpoints as routes.
}

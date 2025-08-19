package middlewares

import (
	"net/http"
	"rest-api/m/utils"

	"github.com/gin-gonic/gin"
)

// for updateevent ->> put will run in the middle of the request. req handler will be executed in the middle of the req
func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization") // header commonly used for tranforming such tokens. that then is a token sent by the client.

	// its also possible client didnt send anything
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"messsage": "Not authorized."})
		return
	}

	// its also possible we do have a token which is not empty but invalid.
	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"messsage": "Not authorized."})
		return
	}
	// method that allows to add some data to the context value. in event go we can extract it from that context.
	context.Set("userId", userId)
	// this will ensure the next request handler.
	context.Next()
}

package routes

import (
	"net/http"
	"rest-api/m/models"
	"rest-api/m/utils"

	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "request parsing failed"})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not save user!"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully."})

}

// login logic

func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not login."})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "user authentication failed."})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "user authentication failed."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login Sucessful!", "token": token})
}

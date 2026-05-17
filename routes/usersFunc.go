package routes

import (
	"net/http"

	"example.com/restapi/models"
	"example.com/restapi/utils"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data."})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created!"})
}

func login(context *gin.Context){
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data."})
		return
	}

	err = user.Validate()
	if err != nil{
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorize"})
		return
	}

	
	token, err := utils.GenerateToken(user.Email, user.ID, user.Role)
	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Unauthorize"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Log in succesful", "token": token})
}
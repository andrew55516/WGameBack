package controllers

import (
	"WGameBack/database"
	helper "WGameBack/helpers"
	"WGameBack/helpers/structConv"
	"WGameBack/models"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

var DB = database.Connection()
var validate = validator.New()

func HashPassword() {

}

func VerifyPassword() {

}

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		isExisting, err := DB.CheckUserByEmail(context.Background(), user.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if isExisting {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user with this email already exists"})
			return
		}

		dbUser := structConv.ToCreateUserParams(user)

		if err := DB.CreateUser(context.Background(), dbUser); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "success"})
		// userCollection 1

	}
}

func Login() {

}

func GetUsers() {

}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("user_id")

		if err := helper.MatchUserTypeToUid(c, userId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user, err := DB.GetUser(context.Background(), userId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		userModel := structConv.ToUserModel(user)

		c.JSON(http.StatusOK, userModel)
	}
}

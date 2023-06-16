package controllers

import (
	"net/http"
	"time"

	"github.com/GA-Marketing/service-viability/db"
	"github.com/GA-Marketing/service-viability/entities"
	"github.com/GA-Marketing/service-viability/helpers"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	entities.User
	Password string `gorm:"->:false"`
}

func UsersList(c *gin.Context) {
	var users []User
	db := db.GetCurrentConnection()
	db.Find(&users)

	c.JSON(200, users)
}

type userInput struct {
	Name         string
	Email        string
	MemberNumber string
	Password     string
	UserType     string
}

func UsersCreate(c *gin.Context) {
	db := db.GetCurrentConnection()

	var userInput userInput
	error := c.ShouldBindJSON(&userInput)

	user := entities.User{
		Name:         userInput.Name,
		Email:        userInput.Email,
		MemberNumber: userInput.MemberNumber,
		UserType:     entities.UserType(userInput.UserType),
		Password:     string(helpers.Encrypt(userInput.Password)),
	}

	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Is not possible get request data",
			"error":   error,
		})
		return
	}

	if user.UserType == "" {
		user.UserType = entities.USER_TYPE_ADMINISTRATOR
	}

	result := db.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": "Not is possible add new user with this data",
			"error":   result.Error,
		})
		return
	}

	c.JSON(201, user)
}

func UsersTypes(c *gin.Context) {

	c.JSON(200, gin.H{
		"types": []string{
			string(entities.USER_TYPE_ADMINISTRATOR),
			string(entities.USER_TYPE_CONSULTANT),
		},
	})
}

type inputLogin struct {
	Email    string
	Password string
}

func UsersAuthLogin(c *gin.Context) {

	var inputLoginData inputLogin
	err := c.BindJSON(&inputLoginData)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid data request",
			"error":   err,
		})
		return
	}

	db := db.GetCurrentConnection()

	user := entities.User{Email: inputLoginData.Email}
	result := db.First(&user)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": "User is not found",
		})

		return
	}

	if !helpers.EncryptVerify(inputLoginData.Password, user.Password) {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"message": "Email or password is invalid",
		})
		return
	}

	key := helpers.GetSecurityKey()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"exp": time.Now().Unix() + 3600,
			"iss": "https://gamarketing.com",
			"sub": "auth",
			"uid": user.ID,
		})
	tokenString, err := token.SignedString([]byte(key))

	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	user.Password = ""

	c.JSON(200, gin.H{
		"user":  user,
		"token": tokenString,
	})
}

func UserMe(c *gin.Context) {
	user := c.MustGet("loggedUser").(entities.User)
	c.JSON(200, user)
}

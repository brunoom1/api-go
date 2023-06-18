package controllers

import (
	"net/http"

	"github.com/GA-Marketing/service-viability/entities"
	"github.com/GA-Marketing/service-viability/helpers"
	"github.com/GA-Marketing/service-viability/repository"
	"github.com/GA-Marketing/service-viability/services"
	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (uc UserController) List(c *gin.Context) {
	var userRepository repository.UsersRepository
	users, err := userRepository.Find()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

type userInput struct {
	Name         string
	Email        string
	MemberNumber string
	Password     string
	UserType     string
}

func (uc UserController) Create(c *gin.Context) {

	var userInput userInput
	error := c.ShouldBindJSON(&userInput)

	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Is not possible get request data",
			"error":   error,
		})
		return
	}

	user := entities.User{
		Name:         userInput.Name,
		Email:        userInput.Email,
		MemberNumber: userInput.MemberNumber,
		UserType:     entities.UserType(userInput.UserType),
		Password:     string(helpers.Encrypt(userInput.Password)),
	}

	var userRepository repository.UsersRepository
	newUser, error := userRepository.Create(&user)

	if error != nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": "Not is possible add new user with this data",
			"error":   error.Error(),
		})
		return
	}

	c.JSON(201, newUser)
}

func (uc UserController) UserTypes(c *gin.Context) {

	var userRepository repository.UsersRepository

	c.JSON(200, gin.H{
		"types": userRepository.FindTypes(),
	})
}

type inputLogin struct {
	Email    string
	Password string
}

func (uc UserController) Login(c *gin.Context) {

	var inputLoginData inputLogin
	err := c.BindJSON(&inputLoginData)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid data request",
			"error":   err,
		})
		return
	}

	var loginService services.LoginService
	token, err := loginService.Login(inputLoginData.Email, inputLoginData.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}

func (uc UserController) Me(c *gin.Context) {
	user := c.MustGet("loggedUser").(entities.User)
	c.JSON(200, user)
}

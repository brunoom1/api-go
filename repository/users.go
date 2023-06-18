package repository

import (
	"github.com/GA-Marketing/service-viability/db"
	"github.com/GA-Marketing/service-viability/entities"
)

type UsersRepository struct{}

type User struct {
	entities.User
	Password string `gorm:"->:false"`
}

func (u UsersRepository) Find() ([]User, error) {
	var users []User
	result := db.GetCurrentConnection().Find(&users)

	return users, result.Error
}

func (u UsersRepository) Create(user *entities.User) (entities.User, error) {
	if user.UserType == "" {
		user.UserType = entities.USER_TYPE_ADMINISTRATOR
	}

	result := db.GetCurrentConnection().Create(&user)

	return *user, result.Error
}

func (u UsersRepository) FindTypes() []entities.UserType {
	var user entities.User
	return user.GetTypes()
}

func (u UsersRepository) ByEmail(email string) (entities.User, error) {
	user := entities.User{Email: email}
	result := db.GetCurrentConnection().First(&user)

	return user, result.Error
}

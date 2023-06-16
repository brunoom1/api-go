package repository

import (
	"github.com/GA-Marketing/service-viability/db"
	"github.com/GA-Marketing/service-viability/entities"
)

type UsersRepository struct{}

func (u UsersRepository) Get() []entities.User {
	var users []entities.User
	db.GetCurrentConnection().Find(&users)

	return users
}

package migration

import (
	"github.com/GA-Marketing/service-viability/db"
	"github.com/GA-Marketing/service-viability/entities"
)

func Run() {

	db := db.GetCurrentConnection()

	user := entities.User{}
	if !db.Migrator().HasTable(user) {
		db.Migrator().CreateTable(user)
	}
}

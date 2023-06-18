package entities

import (
	"time"

	"gorm.io/gorm"
)

type UserType string

const (
	USER_TYPE_CONSULTANT    UserType = "consultant"
	USER_TYPE_ADMINISTRATOR UserType = "administrator"
	USER_TYPE_APPLICATION   UserType = "application"
)

type User struct {
	gorm.Model
	Name         string
	Email        string `gorm:"uniqueIndex"`
	MemberNumber string
	UserType     UserType
	Password     string
	ActivatedAt  time.Time
}

func (u User) GetTypes() []UserType {
	return []UserType{
		USER_TYPE_ADMINISTRATOR,
		USER_TYPE_APPLICATION,
		USER_TYPE_CONSULTANT,
	}
}

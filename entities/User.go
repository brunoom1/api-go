package entities

import (
	"time"

	"gorm.io/gorm"
)

type UserType string

const (
	USER_TYPE_CONSULTANT    UserType = "consultant"
	USER_TYPE_ADMINISTRATOR UserType = "administrator"
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

package services

import (
	"errors"
	"time"

	"github.com/GA-Marketing/service-viability/helpers"
	"github.com/GA-Marketing/service-viability/repository"
	"github.com/golang-jwt/jwt"
)

type LoginService struct{}

func (ls LoginService) Login(email string, password string) (string, error) {

	var userRepository repository.UsersRepository
	user, err := userRepository.ByEmail(email)

	if err != nil {
		return "", err
	}

	if !helpers.EncryptVerify(password, user.Password) {
		return "", errors.New("Email or password is invalid")
	}

	key := helpers.GetSecurityKey()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"exp": time.Now().Unix() + 3600,
			"iss": "https://gamarketing.com",
			"sub": "auth",
			"uid": user.ID,
		})

	return token.SignedString([]byte(key))
}

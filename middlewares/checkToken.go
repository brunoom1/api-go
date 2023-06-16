package middlewares

import (
	"net/http"
	"strings"

	"github.com/GA-Marketing/service-viability/db"
	"github.com/GA-Marketing/service-viability/entities"
	"github.com/GA-Marketing/service-viability/helpers"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func CheckToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.Request.Header.Get("Authorization")

		if authorizationHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{})
			c.Abort()
			return
		}

		tokenString := strings.Split(authorizationHeader, " ")[1]

		type MyCustomClaims struct {
			Uid uint `json:"uid,omitempty"`
			jwt.RegisteredClaims
		}

		t, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(_ *jwt.Token) (interface{}, error) {
			return []byte(helpers.GetSecurityKey()), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
			c.Abort()
			return
		}

		claims := t.Claims.(*MyCustomClaims)

		db := db.GetCurrentConnection()

		user := entities.User{}
		result := db.First(&user, "id = ?", claims.Uid)

		if result.Error != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": result.Error.Error(),
			})
		}

		c.Set("loggedUser", user)
	}
}

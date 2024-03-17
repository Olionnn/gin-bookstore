package middleware

import (
	"fmt"
	"time"

	"github.com/Olionnn/gin-bookstore/config"
	"github.com/Olionnn/gin-bookstore/http/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func AuthMiddleware(c *gin.Context) {
	tokenString, errC := c.Cookie("Authorization")
	if errC != nil {
		c.JSON(401, gin.H{"error": "unauthorized"})
		c.Abort()
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error : %v", token.Header["alg"])
		}

		return []byte(viper.GetString("token.secret")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.JSON(401, gin.H{"error": "token expired"})
			c.Abort()
			return
		}

		var user models.User
		config.DB.First(&user, "id = ?", claims["sub"])

		if user.ID == 0 {
			c.JSON(401, gin.H{"error": "user not found "})
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()

	} else {
		c.JSON(401, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

}

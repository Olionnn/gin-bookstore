package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Olionnn/gin-bookstore/config"
	"github.com/Olionnn/gin-bookstore/http/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var input models.RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Users := models.Users{}

	Users.Username = input.Username
	Users.Email = input.Email
	Users.Password = input.Password

	errHash := Users.HashPassword()
	if errHash != nil {
		fmt.Println(errHash)
		c.JSON(http.StatusInternalServerError, gin.H{"error": errHash.Error()})
		return
	}

	_, errSave := Users.SaveUser()
	if errSave != nil {
		fmt.Println(errSave)
		c.JSON(http.StatusInternalServerError, gin.H{"error": errSave.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "register success"})
}

func Login(c *gin.Context) {
	var input models.LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	config.DB.First(&user, "email = ?", input.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	errHash := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if errHash != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "password is incorrect"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  user.ID,
		"name": user.Username,
		"exp":  time.Now().Add(time.Hour * 1).Unix(),
	})

	tokenString, errSignToken := token.SignedString([]byte(viper.GetString("token.secret")))
	if errSignToken != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errSignToken.Error()})
		return
	}

	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie("Authorization", tokenString, 3600 * 1, "", "", false, true)



	c.JSON(http.StatusOK, gin.H{
		"message": "login success",
		// "token":   tokenString,
	})
}

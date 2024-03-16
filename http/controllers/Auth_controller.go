package controllers

import (
	"fmt"
	"net/http"

	"github.com/Olionnn/gin-bookstore/http/models"
	"github.com/gin-gonic/gin"
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

	Users := models.Users{}

	Users.Email = input.Email
	Users.Password = input.Password

	token, errLogin := models.LoginCheck(Users.Email, Users.Password)
	if errLogin != nil {
		fmt.Println(errLogin)
		c.JSON(http.StatusInternalServerError, gin.H{"error": errLogin.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "login success",
		"token":   token,
	})
}

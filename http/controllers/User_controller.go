package controllers

import (
	"net/http"

	"github.com/Olionnn/gin-bookstore/config"
	"github.com/Olionnn/gin-bookstore/http/models"
	"github.com/gin-gonic/gin"
)

func FindAllUsers(c *gin.Context) {
	data := []models.User{}
	result := config.DB.Find(&data)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": data})
}

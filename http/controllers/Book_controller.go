package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindAllBooks(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "all books"})
}

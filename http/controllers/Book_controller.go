package controllers

import (
	"net/http"

	"github.com/Olionnn/gin-bookstore/config"
	"github.com/Olionnn/gin-bookstore/http/models"
	"github.com/gin-gonic/gin"
)

func FindAllBooks(c *gin.Context) {
	data := []models.Book{}
	result := config.DB.Find(&data)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": data})
}

func AddBook(c *gin.Context) {
	type Request struct {
		Title  string `json:"title" binding:"required"`
		Author string `json:"author" binding:"required"`
		Price  int    `json:"price" binding:"required"`
		Prev   int    `json:"prev"`
		Thumb  string `json:"thumb"`
	}

	var request Request
	var book models.Books

	book.Title = request.Title
	book.Author = request.Author
	book.Price = request.Price
	book.Prev = request.Prev
	book.Thumb = request.Thumb

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := config.DB.Model(&book).Create(&request)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": request})
}

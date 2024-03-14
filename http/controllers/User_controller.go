package controllers

import (
	"fmt"
	"net/http"

	"github.com/Olionnn/gin-bookstore/config"
	"github.com/Olionnn/gin-bookstore/http/models"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	db := config.Connection()

	defer db.Close()

	res, errSelectQuery := db.Query("SELECT id, username, email FROM users")
	if errSelectQuery != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": errSelectQuery.Error()})
		return
	}

	var users []models.Users
	for res.Next() {
		var user models.Users
		errScn := res.Scan(&user.ID, &user.Username, &user.Email)
		if errScn != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": errScn.Error()})
			return
		}
		users = append(users, user)
	}

	c.IndentedJSON(http.StatusOK, users)
}

func AddUsers(c *gin.Context) {
	db := config.Connection()

	defer db.Close()

	var newUser models.Users

	if err := c.ShouldBind(&newUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, errInsertQuery := db.Query("INSERT INTO users (username, email, password) VALUES (?, ?, ?)", newUser.Username, newUser.Email, newUser.Password)
	if errInsertQuery != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": errInsertQuery.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, newUser)
}

func UpdateUsers(c *gin.Context) {
	db := config.Connection()

	defer db.Close()

	id := c.Param("id")

	var updatedUser models.Users
	if err := c.ShouldBind(&updatedUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, errUpdateQuery := db.Query("UPDATE users SET username = ?, email = ?, password = ? WHERE id = ?", updatedUser.Username, updatedUser.Email, updatedUser.Password, id)
	if errUpdateQuery != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": errUpdateQuery.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, updatedUser)
}

func GetUsersByID(c *gin.Context) {
	db := config.Connection()

	defer db.Close()

	id := c.Param("id")

	var user models.Users
	err := db.QueryRow("SELECT * FROM users WHERE id = ?", id).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}

func DeleteUsers(c *gin.Context) {
	db := config.Connection()

	defer db.Close()

	id := c.Param("id")

	_, errDeleteQuery := db.Query("DELETE FROM users WHERE id = ?", id)
	if errDeleteQuery != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": errDeleteQuery.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "User has been deleted"})
}

func Login(c *gin.Context) {
	db := config.Connection()

	defer db.Close()

	var user models.Users
	if err := c.ShouldBind(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(user.Username)

	if user.Username == "" || user.Password == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Username and Password are required"})
		return
	}

	userNotFound, errNotFound := db.Query("SELECT username FROM users WHERE username = ?", user.Username)
	if errNotFound != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": errNotFound.Error()})
		return
	}

	for userNotFound.Next() {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Username not found"})
		return
	}

	err := db.QueryRow("SELECT * FROM users WHERE username = ? AND password = ?", user.Username, user.Password).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, user)
}

func Register(c *gin.Context) {
	db := config.Connection()

	defer db.Close()

	var newUser models.Users

	if err := c.ShouldBind(&newUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if newUser.Username == "" || newUser.Email == "" || newUser.Password == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Username, Email, and Password are required"})
		return
	}

	userNameSame, errSame := db.Query("SELECT username FROM users WHERE username = ?", newUser.Username)
	if errSame != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": errSame.Error()})
		return
	}

	for userNameSame.Next() {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
		return
	}

	_, errInsertQuery := db.Query("INSERT INTO users (username, email, password) VALUES (?, ?, ?)", newUser.Username, newUser.Email, newUser.Password)
	if errInsertQuery != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": errInsertQuery.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, newUser)
}

func Logout(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "You have been logged out"})
}

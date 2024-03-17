package models

import (
	"html"
	"strings"

	"github.com/Olionnn/gin-bookstore/config"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Users struct {
	gorm.Model
	Username string `gorm:"type:varchar(100);unique_index"`
	Email    string `gorm:"type:varchar(100);unique_index"`
	Password string `gorm:"size:255"`
}

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *Users) SaveUser() (*Users, error) {
	err := config.DB.Create(&u).Error
	if err != nil {
		return &Users{}, err
	}
	return u, nil
}

func (u *Users) HashPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	return nil

}

// func verifHashPass(password, hash string) error {
// 	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
// }

// func LoginCheck(email, password string) (string, error) {
// 	var user User
// 	config.DB.First(&user, "email = ?", email)
// 	if user.ID == 0 {
// 		return "", gorm.ErrRecordNotFound
// 	}
// 	errHash := verifHashPass(password, user.Password)
// 	if errHash != nil && errHash == bcrypt.ErrMismatchedHashAndPassword {
// 		return "", errHash
// 	}
// 	return email, nil
// }

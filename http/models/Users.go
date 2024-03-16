package models

import (
	"html"
	"strings"

	"github.com/Olionnn/gin-bookstore/config"
	"github.com/Olionnn/gin-bookstore/utils/token"
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

func (u *Users) SaveUser() (*Users, error) {
	DB, errCN := config.ConnectGorm()
	if errCN != nil {
		return nil, errCN
	}
	defer DB.Close()

	err := DB.Create(&u).Error
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

func verifHashPass(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func LoginCheck(email, password string) (string, error) {
	DB, errCN := config.ConnectGorm()
	if errCN != nil {
		return "", errCN
	}
	defer DB.Close()
	u := Users{}

	errGetUser := DB.Model(Users{}).Where("email = ?", email).Take(&u).Error
	if errGetUser != nil {
		return "", errGetUser
	}

	errVerif := verifHashPass(password, u.Password)
	if errVerif != nil && errVerif == bcrypt.ErrMismatchedHashAndPassword {
		return "", errVerif
	}

	token, errGenToken := token.GenerateToken(u.ID)
	if errGenToken != nil {
		return "", errGenToken
	}

	return token, nil
}

package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Books struct {
	gorm.Model
	Title  string `json:"title" gorm:"type:text;not null"`
	Author string `json:"author" gorm:"type:text;not null"`
	Price  int    `json:"price" gorm:"type:int;not null"`
	Prev   int    `json:"prev" gorm:"type:int;null"`
	Thumb  string `json:"thumb" gorm:"type:text;null"`
}

type Book struct {
	ID       int
	Title    string
	Author   string
	Price    int
	Prev     int
	Thumb    string
	createAt *time.Time
}

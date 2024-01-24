package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model `json:"-"`
	Id         int       `json:"id" gorm:"unique;not null"`
	UserID     int       `json:"userID"`
	Text       string    `json:"name"`
	Place      string    `json:"string"`
	Amount     int       `json:"amount"`
	Collected  int       `json:"collected"`
	Image      string    `json:"image"`
	Date       time.Time `json:"time"`
	CategoryId int       `json:"categoryId"`
}

type Reported struct {
	Id        int    `json:"id" gorm:"unique;not null"`
	UserID    int    `json:"userID"`
	Reason    string `json:"reason"`
	PostID    int    `json:"postID" `
	CommentID int    `json:"commentID"`
	Category  string `json:"category"`
}

type Category struct {
	Id       int    `json:"id" gorm:"unique;not null"`
	Category string `json:"category" gorm:"unique;not null"`
}

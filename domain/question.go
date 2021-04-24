package domain

import "gorm.io/gorm"

type Question struct {
	gorm.Model
	Content string
	Likes int
	Dislikes int
	UserID uint
	User User
	Answers []Answer
}
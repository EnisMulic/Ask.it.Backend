package domain

import "gorm.io/gorm"

type Answer struct {
	gorm.Model
	Content string
	Likes int
	Dislikes int
	UserID uint
	User User
	QuestionID uint
}
package domain

import "gorm.io/gorm"

type Answer struct {
	gorm.Model
	Content string
	UserID uint
	QuestionID uint
}
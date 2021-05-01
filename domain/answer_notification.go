package domain

import "gorm.io/gorm"

type AnswerNotification struct {
	gorm.Model
	UserID uint
	User User `gorm:"foreignkey:UserID"`
	QuestionID uint
	Question Question `gorm:"foreignkey:QuestionID"`
	AnswerID uint
	Answer Answer `gorm:"foreignkey:AnswerID"`
	Content string
	IsRead bool
}
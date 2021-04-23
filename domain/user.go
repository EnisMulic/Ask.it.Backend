package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string `gorm:"uniqueIndex;size:320"`
	Password  string
	Questions []Question
	Answers []Answer
	QuestionRatings []Question `gorm:"many2many:user_question_ratings;"`
	AnswerRatings []Answer `gorm:"many2many:user_answer_ratings;"`
}
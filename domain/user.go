package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string `gorm:"uniqueIndex;size:320"`
	PasswordSalt  string
	PasswordHash  string
	Questions []Question
	Answers []Answer
	AnswerCount int
	QuestionRatings []Question `gorm:"many2many:user_question_ratings;"`
	AnswerRatings []Answer `gorm:"many2many:user_answer_ratings;"`
	UserQuestionRatings []UserQuestionRating `gorm:"->"`
	UserAnswerRatings []UserAnswerRating `gorm:"->"`
}
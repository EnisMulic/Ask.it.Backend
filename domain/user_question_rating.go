package domain

type UserQuestionRating struct {
	UserID     uint `gorm:"primaryKey"`
	QuestionID uint `gorm:"primaryKey"`
	IsLiked    bool
}
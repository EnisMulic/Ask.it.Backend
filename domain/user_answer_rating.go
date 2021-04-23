package domain

type UserAnswerRating struct {
	UserID   uint `gorm:"primaryKey"`
	AnswerID uint `gorm:"primaryKey"`
	IsLiked  bool
}
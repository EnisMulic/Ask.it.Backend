package domain

import "gorm.io/gorm"

type UserQuestionRating struct {
	UserID     uint `gorm:"primaryKey"`
	QuestionID uint `gorm:"primaryKey"`
	IsLiked    bool
}

func (uqr *UserQuestionRating) AfterCreate(tx *gorm.DB) (err error) {
	var question Question
	tx.Find(&question, uqr.QuestionID)

	var likeIncrement int
	var dislikeIncrement int

	if uqr.IsLiked {
		likeIncrement = 1
		dislikeIncrement = 0
	} else {
		likeIncrement = 0
		dislikeIncrement = 1
	}

	result := tx.Model(&question).Updates(Question{
		Likes:    question.Likes + likeIncrement,
		Dislikes: question.Dislikes + dislikeIncrement,
	})

	return result.Error
}
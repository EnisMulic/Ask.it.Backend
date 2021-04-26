package domain

import "gorm.io/gorm"

type UserAnswerRating struct {
	UserID   uint `gorm:"primaryKey"`
	AnswerID uint `gorm:"primaryKey"`
	IsLiked  bool
}

func (uar *UserAnswerRating) AfterCreate(tx *gorm.DB) (err error) {
	var like int
	var dislike int

	if uar.IsLiked {
		like = 1
		dislike = 0
	} else {
		like = 0
		dislike = 1
	}

	var answer Answer
	tx.Find(&answer, uar.AnswerID)

	result := tx.Model(&answer).Where("id = ?", uar.AnswerID).Updates(Answer{
		Likes:    answer.Likes + like,
		Dislikes: answer.Dislikes + dislike,
	})

	return result.Error
}

func (uar *UserAnswerRating) AfterUpdate(tx *gorm.DB) (err error) {
	var like int
	var dislike int

	if uar.IsLiked {
		like = 1
		dislike = -1
	} else {
		like = -1
		dislike = 1
	}

	var answer Answer
	tx.Find(&answer, uar.AnswerID)

	result := tx.Model(&answer).Where("id = ?", uar.AnswerID).Updates(Answer{
		Likes:    answer.Likes + like,
		Dislikes: answer.Dislikes + dislike,
	})

	return result.Error
}

func (uar *UserAnswerRating) BeforeDelete(tx *gorm.DB) (err error) {
	var like int
	var dislike int
	

	if uar.IsLiked {
		like = 1
		dislike = 0
	} else {
		like = 0
		dislike = 1
	}

	var answer Answer
	tx.Find(&answer, uar.AnswerID)

	result := tx.Model(&answer).Where("id = ?", uar.AnswerID).Updates(Answer{
		Likes:    answer.Likes - like,
		Dislikes: answer.Dislikes - dislike,
	})

	return result.Error
}
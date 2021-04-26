package domain

import "gorm.io/gorm"

type UserQuestionRating struct {
	UserID     uint `gorm:"primaryKey"`
	QuestionID uint `gorm:"primaryKey"`
	IsLiked    bool
}

func (uqr *UserQuestionRating) AfterCreate(tx *gorm.DB) (err error) {
	var like int
	var dislike int

	if uqr.IsLiked {
		like = 1
		dislike = 0
	} else {
		like = 0
		dislike = 1
	}

	var question Question
	tx.Find(&question, uqr.QuestionID)
	
	result := tx.Model(&question).Where("id = ?", uqr.QuestionID).Updates(Question{
		Likes:    question.Likes + like,
		Dislikes: question.Dislikes + dislike,
	})

	return result.Error
}

func (uqr *UserQuestionRating) AfterUpdate(tx *gorm.DB) (err error) {
	var like int
	var dislike int

	if uqr.IsLiked {
		like = 1
		dislike = -1
	} else {
		like = -1
		dislike = 1
	}

	var question Question
	tx.Find(&question, uqr.QuestionID)

	result := tx.Model(&question).Where("id = ?", uqr.QuestionID).Updates(Question{
		Likes:    question.Likes + like,
		Dislikes: question.Dislikes + dislike,
	})

	return result.Error
}

func (uqr *UserQuestionRating) BeforeDelete(tx *gorm.DB) (err error) {
	
	var like int
	var dislike int

	if uqr.IsLiked {
		like = 1
		dislike = 0
	} else {
		like = 0
		dislike = 1
	}

	var question Question
	tx.Find(&question, uqr.QuestionID)

	result := tx.Model(&question).Where("id = ?", uqr.QuestionID).Updates(Question{
		Likes:    question.Likes - like,
		Dislikes: question.Dislikes - dislike,
	})

	return result.Error
}
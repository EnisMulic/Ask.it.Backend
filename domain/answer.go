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

func (a *Answer) AfterCreate(tx *gorm.DB) (err error) {
	var user User
	tx.Find(&user, a.UserID)
	
	result := tx.Model(&user).Where("id = ?", a.UserID).Updates(User{
		AnswerCount: user.AnswerCount + 1,
	})

	return result.Error
}

func (a *Answer) BeforeDelete(tx *gorm.DB) (err error) {
	var user User
	tx.Find(&user, a.UserID)
	
	result := tx.Model(&user).Where("id = ?", a.UserID).Updates(User{
		AnswerCount: user.AnswerCount - 1,
	})

	return result.Error
}
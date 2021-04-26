package repositories

import (
	"github.com/EnisMulic/Ask.it.Backend/domain"
	"gorm.io/gorm"
)


type UserQuestionRatingRepository struct {
	db *gorm.DB
}

func NewUserQuestionRatingRepository(db *gorm.DB) *UserQuestionRatingRepository {
	return &UserQuestionRatingRepository{db}
}

func (uqrr *UserQuestionRatingRepository) Get (questionId uint, userId uint) (domain.UserQuestionRating, error) {
	var rating domain.UserQuestionRating
	result := uqrr.db.Where("user_id = ?", userId).Where("question_id = ?", questionId).First(&rating)
	return rating, result.Error
}

func (uqrr *UserQuestionRatingRepository) Create (rating domain.UserQuestionRating) (domain.UserQuestionRating, error) {
	result := uqrr.db.Create(&rating)
	return rating, result.Error
}

func (uqrr *UserQuestionRatingRepository) Update (rating domain.UserQuestionRating, newRating domain.UserQuestionRating) (domain.UserQuestionRating, error) {
	result := uqrr.db.Model(&rating).Updates(domain.UserQuestionRating{
		IsLiked: newRating.IsLiked,
	})

	return rating, result.Error
}

func (uqrr *UserQuestionRatingRepository) Delete (rating domain.UserQuestionRating) error {
	result := uqrr.db.Delete(&rating)
	return result.Error
}
package repositories

import (
	"github.com/EnisMulic/Ask.it.Backend/domain"
	"gorm.io/gorm"
)


type UserAnswerRatingRepository struct {
	db *gorm.DB
}

func NewUserAnswerRatingRepository(db *gorm.DB) *UserAnswerRatingRepository {
	return &UserAnswerRatingRepository{db}
}

func (r *UserAnswerRatingRepository) Get (answerId uint, userId uint) (domain.UserAnswerRating, error) {
	var rating domain.UserAnswerRating
	result := r.db.Where("user_id = ?", userId).Where("answer_id = ?", answerId).First(&rating)
	return rating, result.Error
}

func (r *UserAnswerRatingRepository) Create (rating domain.UserAnswerRating) (domain.UserAnswerRating, error) {
	result := r.db.Create(&rating)
	return rating, result.Error
}

func (r *UserAnswerRatingRepository) Update (rating domain.UserAnswerRating, newRating domain.UserAnswerRating) (domain.UserAnswerRating, error) {
	result := r.db.Model(&rating).Updates(domain.UserAnswerRating{
		IsLiked: newRating.IsLiked,
	})

	return rating, result.Error
}

func (r *UserAnswerRatingRepository) Delete (rating domain.UserAnswerRating) error {
	result := r.db.Delete(&rating)
	return result.Error
}
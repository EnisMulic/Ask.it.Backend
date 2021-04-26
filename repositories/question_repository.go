package repositories

import (
	"github.com/EnisMulic/Ask.it.Backend/domain"
	"gorm.io/gorm"
)


type QuestionRepository struct {
	db *gorm.DB
}

func NewQuestionRepository(db *gorm.DB) *QuestionRepository {
	return &QuestionRepository{db}
}

type QuestionFilter struct {
	PageNumber int
	PageSize int
	UserID uint
}

func (ur *QuestionRepository) GetPaged (filter QuestionFilter) []domain.Question {
	var questions []domain.Question
	query := ur.db

	if (QuestionFilter{} != filter) {
		
		if filter.UserID != 0 {
			query = query.Where("user_id = ?", filter.UserID)
		}

		if filter.PageNumber > 0 && filter.PageSize > 0 {
			query = query.Limit(filter.PageSize).Offset((filter.PageNumber - 1) * filter.PageSize)
		}
	}
	

	query.Joins("User").Find(&questions)

	return questions
}

func (ur *QuestionRepository) GetById (id uint) domain.Question {
	var question domain.Question
	ur.db.Joins("User").First(&question, id)
	return question
}

func (ur *QuestionRepository) Create (question domain.Question) (domain.Question, error) {
	result := ur.db.Create(&question)
	return question, result.Error
}

func (ur *QuestionRepository) Delete (question domain.Question) {
	ur.db.Delete(&question)
}
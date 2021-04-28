package repositories

import (
	"github.com/EnisMulic/Ask.it.Backend/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)


type QuestionRepository struct {
	db *gorm.DB
}

func NewQuestionRepository(db *gorm.DB) *QuestionRepository {
	return &QuestionRepository{db}
}

type PaginationFilter struct {
	PageNumber int
	PageSize int
}

type QuestionFilter struct {
	UserID uint
}

func (ur *QuestionRepository) GetPaged (filter QuestionFilter, sorting []SortFilter, pagination PaginationFilter) ([]domain.Question, int64) {
	var questions []domain.Question
	query := ur.db.Model(&domain.Question{})

	var sortStr string
	for i, sort := range sorting {
		sortStr = sort.Column + " " + sort.Order
		if i < len(sorting) - 1 {
			sortStr = sortStr + ","
		}
	}
	
	if sortStr != "" {
		query = query.Order(sortStr)
	}

	if (QuestionFilter{} != filter) {
		
		if filter.UserID != 0 {
			query = query.Where("user_id = ?", filter.UserID)
		}
	}

	var count int64
	query.Count(&count)

	if (PaginationFilter{} != pagination) {
		if pagination.PageNumber > 0 && pagination.PageSize > 0 {
			query = query.Limit(pagination.PageSize).Offset((pagination.PageNumber - 1) * pagination.PageSize)
		}
	}
	

	query.Joins("User").Find(&questions)
	
	return questions, count
}

func (ur *QuestionRepository) GetById (id uint) (domain.Question, error) {
	var question domain.Question
	result := ur.db.Preload("Answers.User").Preload(clause.Associations).First(&question, id)
	return question, result.Error
}

func (ur *QuestionRepository) Create (question domain.Question) (domain.Question, error) {
	result := ur.db.Create(&question)
	return question, result.Error
}

func (ur *QuestionRepository) Delete (question domain.Question) {
	ur.db.Delete(&question)
}
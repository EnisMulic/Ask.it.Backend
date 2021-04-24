package repositories

import (
	"github.com/EnisMulic/Ask.it.Backend/contracts/requests"
	"github.com/EnisMulic/Ask.it.Backend/domain"
	"gorm.io/gorm"
)


type QuestionRepository struct {
	db *gorm.DB
}

func NewQuestionRepository(db *gorm.DB) *QuestionRepository {
	return &QuestionRepository{db}
}

func (ur *QuestionRepository) GetPaged (search requests.QuestionSearchRequest) []domain.Question {
	var questions []domain.Question
	query := ur.db
	if (requests.QuestionSearchRequest{} != search) && search.PageNumber > 0 && search.PageSize > 0 {
		query = query.Limit(search.PageNumber).Offset((search.PageNumber - 1) * search.PageSize)
	}

	query.Find(&questions)

	return questions
}
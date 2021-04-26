package repositories

import (
	"github.com/EnisMulic/Ask.it.Backend/domain"
	"gorm.io/gorm"
)


type AnswerRepository struct {
	db *gorm.DB
}

func NewAnswerRepository(db *gorm.DB) *AnswerRepository {
	return &AnswerRepository{db}
}

func (ur *AnswerRepository) GetById (id uint) (domain.Answer, error) {
	var answer domain.Answer
	result := ur.db.Joins("User").First(&answer, id)
	return answer, result.Error  
}

func (ar *AnswerRepository) Create(answer domain.Answer) (domain.Answer, error) {
	result := ar.db.Create(&answer)
	return answer, result.Error
}

func (ar *AnswerRepository) Update(answer domain.Answer, updatedAnswer domain.Answer) (domain.Answer, error) {
	result := ar.db.Model(&answer).Updates(domain.Answer{
		Content: updatedAnswer.Content,
	})

	return answer, result.Error
}
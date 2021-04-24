package services

import (
	"github.com/EnisMulic/Ask.it.Backend/contracts/requests"
	"github.com/EnisMulic/Ask.it.Backend/contracts/responses"
	"github.com/EnisMulic/Ask.it.Backend/domain"
	"github.com/EnisMulic/Ask.it.Backend/repositories"
)


type QuestionService struct {
	repo *repositories.QuestionRepository
}

func NewQuestionService(repo *repositories.QuestionRepository) *QuestionService {
	return &QuestionService{repo}
} 

func convertToQuestionResponseModel(question domain.Question) responses.QuestionResponseModel {
	return responses.QuestionResponseModel{
		ID: question.ID,
		Content: question.Content,
		CreatedAt: question.CreatedAt,
		Likes: question.Likes,
		Dislikes: question.Dislikes,
	}
}

func (qs *QuestionService) Get (search requests.QuestionSearchRequest) responses.QuestionsReponse {
	questions := qs.repo.GetPaged(search)

	var response []responses.QuestionResponseModel
	for _, question := range questions {
		questionResponse := convertToQuestionResponseModel(question)
		response = append(response, questionResponse)
	}

	return responses.QuestionsReponse{Data: response}
}
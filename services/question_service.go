package services

import (
	"errors"

	"github.com/EnisMulic/Ask.it.Backend/contracts/requests"
	"github.com/EnisMulic/Ask.it.Backend/contracts/responses"
	"github.com/EnisMulic/Ask.it.Backend/repositories"
	"github.com/EnisMulic/Ask.it.Backend/utils"
)

var ErrorQuestionNotFound error = errors.New("question not found")

type QuestionService struct {
	repo *repositories.QuestionRepository
}

func NewQuestionService(repo *repositories.QuestionRepository) *QuestionService {
	return &QuestionService{repo}
} 

func (qs *QuestionService) Get (search requests.QuestionSearchRequest) responses.QuestionsReponse {
	questions := qs.repo.GetPaged(search)

	var response []responses.QuestionResponseModel
	for _, question := range questions {
		questionResponse := utils.ConvertToQuestionResponseModel(question)
		response = append(response, questionResponse)
	}

	return responses.QuestionsReponse{Data: response}
}

func (qs *QuestionService) GetById (id uint) (*responses.QuestionResponse, *responses.ErrorResponse) {
	question := qs.repo.GetById(id)

	if question.ID == 0 {
		err := responses.ErrorResponseModel{
			FieldName: "",
			Message: ErrorQuestionNotFound.Error(),
		}

		errors := responses.NewErrorResponse(err)	

		return nil, errors
	}

	response := utils.ConvertToQuestionResponseModel(question)

	return &responses.QuestionResponse{Data: response}, nil
}
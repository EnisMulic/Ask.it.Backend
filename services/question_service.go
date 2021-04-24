package services

import (
	"errors"

	"github.com/EnisMulic/Ask.it.Backend/contracts/requests"
	"github.com/EnisMulic/Ask.it.Backend/contracts/responses"
	"github.com/EnisMulic/Ask.it.Backend/domain"
	"github.com/EnisMulic/Ask.it.Backend/repositories"
)

var ErrorQuestionNotFound error = errors.New("question not found")

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
		User: convertToUserResponseModel(question.User),
	}
}

func convertToUserResponseModel(user domain.User) responses.UserResponseModel {
	return responses.UserResponseModel{
		ID: user.ID,
		FirstName: user.FirstName,
		LastName: user.LastName,
		Email: user.Email,
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

	response := convertToQuestionResponseModel(question)

	return &responses.QuestionResponse{Data: response}, nil
}
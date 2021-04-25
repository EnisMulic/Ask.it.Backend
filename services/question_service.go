package services

import (
	"errors"

	"github.com/EnisMulic/Ask.it.Backend/contracts/requests"
	"github.com/EnisMulic/Ask.it.Backend/contracts/responses"
	"github.com/EnisMulic/Ask.it.Backend/domain"
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
	var filter repositories.QuestionFilter

	if (search != requests.QuestionSearchRequest{}) {
		filter = repositories.QuestionFilter{
			PageNumber: search.PageNumber,
			PageSize: search.PageSize,
		}
	}

	questions := qs.repo.GetPaged(filter)

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

func (qs *QuestionService) Create (userId uint, req requests.QuestionInsertRequest) (*responses.QuestionResponse, *responses.ErrorResponse) {
	question := domain.Question{
		Content: req.Content,
		UserID: userId,
	}

	newQuestion, err := qs.repo.Create(question)
	if err != nil {
		err := responses.ErrorResponseModel{
			FieldName: "",
			Message: err.Error(),
		}

		errors := responses.NewErrorResponse(err)	

		return nil, errors
	}

	newQuestion = qs.repo.GetById(newQuestion.ID)
	response := utils.ConvertToQuestionResponseModel(newQuestion)
	
	return &responses.QuestionResponse{Data: response}, nil
}

func (qs *QuestionService) Delete (questionId uint, userId uint) *responses.ErrorResponse {
	question := qs.repo.GetById(questionId)

	if question.UserID != userId {
		err := responses.ErrorResponseModel{
			FieldName: "",
			Message: "You do not have permission to delete this question",
		}

		errors := responses.NewErrorResponse(err)	

		return errors
	}

	qs.repo.Delete(question)
	return nil
}
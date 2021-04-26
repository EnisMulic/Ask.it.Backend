package services

import (
	"errors"

	"github.com/EnisMulic/Ask.it.Backend/contracts/requests"
	"github.com/EnisMulic/Ask.it.Backend/contracts/responses"
	"github.com/EnisMulic/Ask.it.Backend/domain"
	"github.com/EnisMulic/Ask.it.Backend/repositories"
	"github.com/EnisMulic/Ask.it.Backend/utils"
)

var ErrorAnswerDeletePermission = errors.New("you do not have permission to delete this answer")

type AnswerService struct {
	answerRepo *repositories.AnswerRepository
}

func NewAnswerRepository(ar *repositories.AnswerRepository) *AnswerService {
	return &AnswerService{ar}
}

func (as *AnswerService) Update (
	answerId uint,
	userId uint,
	req requests.AnswerUpdateRequest,
) (*responses.AnswerResponse, *responses.ErrorResponse) {
	answer, _ := as.answerRepo.GetById(answerId)

	if answer.ID == 0 {
		err := responses.ErrorResponseModel{
			FieldName: "",
			Message: ErrorAnswerNotFound.Error(),
		}

		errors := responses.NewErrorResponse(err)	

		return nil, errors
	}

	if answer.UserID != userId {
		err := responses.ErrorResponseModel{
			FieldName: "",
			Message: ErrorAnswerEditPermission.Error(),
		}

		errors := responses.NewErrorResponse(err)	

		return nil, errors
	}

	updatedAnswer := domain.Answer{
		Content: req.Content,
	}

	answer, err := as.answerRepo.Update(answer, updatedAnswer)

	if err != nil {
		err := responses.ErrorResponseModel{
			FieldName: "",
			Message: err.Error(),
		}

		errors := responses.NewErrorResponse(err)	

		return nil, errors
	}

	response := utils.ConvertToAnswerResponseModel(answer)
	return &responses.AnswerResponse{
		Data: response,
	}, nil
}

func (as *AnswerService) Delete(answerId uint, userId uint ) *responses.ErrorResponse {
	answer, err := as.answerRepo.GetById(answerId)

	if err != nil {
		err := responses.ErrorResponseModel{
			FieldName: "",
			Message: err.Error(),
		}

		errors := responses.NewErrorResponse(err)	

		return errors
	}

	if answer.UserID != userId {
		err := responses.ErrorResponseModel{
			FieldName: "",
			Message: ErrorAnswerDeletePermission.Error(),
		}

		errors := responses.NewErrorResponse(err)	

		return errors
	}

	err = as.answerRepo.Delete(answer)

	if err != nil {
		err := responses.ErrorResponseModel{
			FieldName: "",
			Message: err.Error(),
		}

		errors := responses.NewErrorResponse(err)	

		return errors
	}

	return nil
}
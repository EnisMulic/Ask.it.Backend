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
	ratingRepo *repositories.UserAnswerRatingRepository
}

func NewAnswerRepository(ar *repositories.AnswerRepository, arr *repositories.UserAnswerRatingRepository) *AnswerService {
	return &AnswerService{ar, arr}
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

func (as *AnswerService) Like (answerId uint, userId uint) *responses.ErrorResponse {
	rating, err := as.ratingRepo.Get(answerId, userId)
	
	if err != nil {
		_, err := as.ratingRepo.Create(domain.UserAnswerRating{
			UserID: userId,
			AnswerID: answerId,
			IsLiked: true,
		})

		if err != nil {
			resErr := responses.ErrorResponseModel{
				FieldName: "",
				Message: "An error occurred",
			}

			errors := responses.NewErrorResponse(resErr)	

			return errors
		}

		return nil
	}

	if !rating.IsLiked {
		_, err = as.ratingRepo.Update(rating, domain.UserAnswerRating{
			IsLiked: true,
		})

		if err != nil {
			resErr := responses.ErrorResponseModel{
				FieldName: "",
				Message: "An error occurred",
			}

			errors := responses.NewErrorResponse(resErr)	

			return errors
		}
	}

	return nil
}

func (as *AnswerService) LikeUndo (answerId uint, userId uint) *responses.ErrorResponse {
	rating, err := as.ratingRepo.Get(answerId, userId)

	if err != nil {
		err := responses.ErrorResponseModel{
			FieldName: "",
			Message: "An error occurred",
		}

		errors := responses.NewErrorResponse(err)	

		return errors
	}

	if rating.IsLiked {
		err := as.ratingRepo.Delete(rating)
		
		if err != nil {
			err := responses.ErrorResponseModel{
				FieldName: "",
				Message: "An error occurred",
			}

			errors := responses.NewErrorResponse(err)	

			return errors
		}
	}

	return nil
}

func (as *AnswerService) Dislike (answerId uint, userId uint) *responses.ErrorResponse {
	rating, err := as.ratingRepo.Get(answerId, userId)
	
	if err != nil {
		_, err := as.ratingRepo.Create(domain.UserAnswerRating{
			UserID: userId,
			AnswerID: answerId,
			IsLiked: false,
		})

		if err != nil {
			resErr := responses.ErrorResponseModel{
				FieldName: "",
				Message: "An error occurred",
			}

			errors := responses.NewErrorResponse(resErr)	

			return errors
		}

		return nil
	}

	if rating.IsLiked {
		_, err = as.ratingRepo.Update(rating, domain.UserAnswerRating{
			IsLiked: false,
		})

		if err != nil {
			resErr := responses.ErrorResponseModel{
				FieldName: "",
				Message: "An error occurred",
			}

			errors := responses.NewErrorResponse(resErr)	

			return errors
		}
	}

	return nil
}

func (as *AnswerService) DislikeUndo (answerId uint, userId uint) *responses.ErrorResponse {
	rating, err := as.ratingRepo.Get(answerId, userId)

	if err != nil {
		err := responses.ErrorResponseModel{
			FieldName: "",
			Message: "An error occurred",
		}

		errors := responses.NewErrorResponse(err)	

		return errors
	}

	if !rating.IsLiked {
		
		err := as.ratingRepo.Delete(rating)

		if err != nil {
			err := responses.ErrorResponseModel{
				FieldName: "",
				Message: "An error occurred",
			}

			errors := responses.NewErrorResponse(err)	

			return errors
		}
	}

	return nil
}
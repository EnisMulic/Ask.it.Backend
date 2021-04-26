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
var ErrorAnswerNotFound error = errors.New("answer not found")
var ErrorAnswerEditPermission error = errors.New("you do not have permission to edit this answer")

type QuestionService struct {
	repo *repositories.QuestionRepository
	ratingRepo *repositories.UserQuestionRatingRepository
	answerRepo *repositories.AnswerRepository
}

func NewQuestionService(
	repo *repositories.QuestionRepository, 
	ratingRepo *repositories.UserQuestionRatingRepository,
	answerRepo *repositories.AnswerRepository,
) *QuestionService {
	return &QuestionService{repo, ratingRepo, answerRepo}
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

func (qs *QuestionService) Like (questionId uint, userId uint) *responses.ErrorResponse {
	rating, err := qs.ratingRepo.Get(questionId, userId)
	
	if err != nil {
		_, err := qs.ratingRepo.Create(domain.UserQuestionRating{
			UserID: userId,
			QuestionID: questionId,
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
		_, err = qs.ratingRepo.Update(rating, domain.UserQuestionRating{
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

func (qs *QuestionService) LikeUndo (questionId uint, userId uint) *responses.ErrorResponse {
	rating, err := qs.ratingRepo.Get(questionId, userId)

	if err != nil {
		err := responses.ErrorResponseModel{
			FieldName: "",
			Message: "An error occurred",
		}

		errors := responses.NewErrorResponse(err)	

		return errors
	}

	if rating.IsLiked {
		err := qs.ratingRepo.Delete(rating)
		
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

func (qs *QuestionService) Dislike (questionId uint, userId uint) *responses.ErrorResponse {
	rating, err := qs.ratingRepo.Get(questionId, userId)
	
	if err != nil {
		_, err := qs.ratingRepo.Create(domain.UserQuestionRating{
			UserID: userId,
			QuestionID: questionId,
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
		_, err = qs.ratingRepo.Update(rating, domain.UserQuestionRating{
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

func (qs *QuestionService) DislikeUndo (questionId uint, userId uint) *responses.ErrorResponse {
	rating, err := qs.ratingRepo.Get(questionId, userId)

	if err != nil {
		err := responses.ErrorResponseModel{
			FieldName: "",
			Message: "An error occurred",
		}

		errors := responses.NewErrorResponse(err)	

		return errors
	}

	if !rating.IsLiked {
		
		err := qs.ratingRepo.Delete(rating)

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

func (qs *QuestionService) CreateAnswer (questionId uint, userId uint, req requests.AnswerInsertRequest) (*responses.AnswerResponse, *responses.ErrorResponse) {
	answer := domain.Answer{
		QuestionID: questionId,
		UserID: userId,
		Content: req.Content,
	}

	newAnswer, err := qs.answerRepo.Create(answer)
	if err != nil {
		err := responses.ErrorResponseModel{
			FieldName: "",
			Message: err.Error(),
		}

		errors := responses.NewErrorResponse(err)	

		return nil, errors
	}

	newAnswer, _ = qs.answerRepo.GetById(newAnswer.ID)
	response := utils.ConvertToAnswerResponseModel(newAnswer)
	
	return &responses.AnswerResponse{Data: response}, nil
}
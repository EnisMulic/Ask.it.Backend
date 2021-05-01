package services

import (
	"encoding/json"
	"errors"

	"github.com/EnisMulic/Ask.it.Backend/contracts/requests"
	"github.com/EnisMulic/Ask.it.Backend/contracts/responses"
	"github.com/EnisMulic/Ask.it.Backend/domain"
	"github.com/EnisMulic/Ask.it.Backend/repositories"
	"github.com/EnisMulic/Ask.it.Backend/utils"
	"github.com/EnisMulic/Ask.it.Backend/websockets"
)

var ErrorQuestionNotFound error = errors.New("question not found")
var ErrorAnswerNotFound error = errors.New("answer not found")
var ErrorAnswerEditPermission error = errors.New("you do not have permission to edit this answer")

type QuestionService struct {
	repo *repositories.QuestionRepository
	ratingRepo *repositories.UserQuestionRatingRepository
	answerRepo *repositories.AnswerRepository
	pool *websockets.Pool
}

func NewQuestionService(
	repo *repositories.QuestionRepository, 
	ratingRepo *repositories.UserQuestionRatingRepository,
	answerRepo *repositories.AnswerRepository,
	pool *websockets.Pool,
) *QuestionService {
	return &QuestionService{repo, ratingRepo, answerRepo, pool}
} 

func (qs *QuestionService) Get (search requests.QuestionSearchRequest) *responses.QuestionsReponse {
	var filter repositories.QuestionFilter
	var pagination repositories.PaginationFilter

	if (search.PaginationQuery != nil) {
		pagination = repositories.PaginationFilter{
			PageNumber: search.PageNumber,
			PageSize: search.PageSize,
		}
	}

	sort := repositories.SortFilter{
		Column: "created_at",
		Order: "desc",
	}

	questions, count := qs.repo.GetPaged(filter, []repositories.SortFilter{sort}, pagination)

	var response []responses.QuestionResponseModel
	for _, question := range questions {
		questionResponse := utils.ConvertToQuestionResponseModel(question)
		response = append(response, questionResponse)
	}

	if (search != requests.QuestionSearchRequest{}) {
		return utils.CreateQuestionPagedResponse(response, count, int64(search.PageNumber), int64(search.PageSize))
	} else {
		return utils.CreateQuestionPagedResponse(response, count, 0, 0)
	}
}

func (qs *QuestionService) GetHot (search requests.QuestionSearchRequest) *responses.QuestionsReponse {
	var filter repositories.QuestionFilter
	var pagination repositories.PaginationFilter

	if (search.PaginationQuery != nil) {
		pagination = repositories.PaginationFilter{
			PageNumber: search.PageNumber,
			PageSize: search.PageSize,
		}
	}

	sort := repositories.SortFilter{
		Column: "likes",
		Order: "desc",
	}

	questions, count := qs.repo.GetPaged(filter, []repositories.SortFilter{sort}, pagination)

	var response []responses.QuestionResponseModel
	for _, question := range questions {
		questionResponse := utils.ConvertToQuestionResponseModel(question)
		response = append(response, questionResponse)
	}

	if (search != requests.QuestionSearchRequest{}) {
		return utils.CreateQuestionPagedResponse(response, count, int64(search.PageNumber), int64(search.PageSize))
	} else {
		return utils.CreateQuestionPagedResponse(response, count, 0, 0)
	}
}

func (qs *QuestionService) GetById (id uint) (*responses.QuestionResponse, *responses.ErrorResponse) {
	question, _ := qs.repo.GetById(id)

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

	newQuestion, _ = qs.repo.GetById(newQuestion.ID)
	response := utils.ConvertToQuestionResponseModel(newQuestion)
	
	return &responses.QuestionResponse{Data: response}, nil
}

func (qs *QuestionService) Delete (questionId uint, userId uint) *responses.ErrorResponse {
	question, _ := qs.repo.GetById(questionId)

	if question.ID == 0 {
		err := responses.ErrorResponseModel{
			FieldName: "",
			Message: ErrorQuestionNotFound.Error(),
		}

		errors := responses.NewErrorResponse(err)	

		return errors
	}

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
	question, _ := qs.repo.GetById(newAnswer.QuestionID)
	response := utils.ConvertToAnswerResponseModel(newAnswer)

	if question.UserID != userId {
		notification := responses.AnswerNotification{
			QuestionID: question.ID,
			Question: question.Content,
			User: newAnswer.User.Email,
		}

		msg, _ := json.Marshal(notification)
		
		qs.pool.Broadcast <- websockets.Message{
			ClientID: uint64(userId), 
			Body: string(msg),
		}
	}
	
	return &responses.AnswerResponse{Data: response}, nil
}
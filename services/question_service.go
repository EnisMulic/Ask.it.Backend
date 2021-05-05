package services

import (
	"encoding/json"

	"github.com/EnisMulic/Ask.it.Backend/constants"
	"github.com/EnisMulic/Ask.it.Backend/contracts/requests"
	"github.com/EnisMulic/Ask.it.Backend/contracts/responses"
	"github.com/EnisMulic/Ask.it.Backend/domain"
	"github.com/EnisMulic/Ask.it.Backend/repositories"
	"github.com/EnisMulic/Ask.it.Backend/utils"
	"github.com/EnisMulic/Ask.it.Backend/websockets"
)

type QuestionService struct {
	repo *repositories.QuestionRepository
	ratingRepo *repositories.UserQuestionRatingRepository
	answerRepo *repositories.AnswerRepository
	answerNotificationRepo *repositories.AnswerNotificationRepository
	pool *websockets.Pool
}

func NewQuestionService(
	repo *repositories.QuestionRepository, 
	ratingRepo *repositories.UserQuestionRatingRepository,
	answerRepo *repositories.AnswerRepository,
	answerNotificationRepo *repositories.AnswerNotificationRepository,
	pool *websockets.Pool,
) *QuestionService {
	return &QuestionService{repo, ratingRepo, answerRepo, answerNotificationRepo, pool}
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

func (qs *QuestionService) GetById (id uint) (*responses.QuestionResponse, error) {
	question, _ := qs.repo.GetById(id)

	if question.ID == 0 {
		return nil, constants.ErrQuestionNotFound
	}

	response := utils.ConvertToQuestionResponseModel(question)

	return &responses.QuestionResponse{Data: response}, nil
}

func (qs *QuestionService) Create (userId uint, req requests.QuestionInsertRequest) (*responses.QuestionResponse, error) {
	question := domain.Question{
		Content: req.Content,
		UserID: userId,
	}

	newQuestion, err := qs.repo.Create(question)
	if err != nil {
		return nil, constants.ErrGeneric
	}

	newQuestion, _ = qs.repo.GetById(newQuestion.ID)
	response := utils.ConvertToQuestionResponseModel(newQuestion)
	
	return &responses.QuestionResponse{Data: response}, nil
}

func (qs *QuestionService) Delete (questionId uint, userId uint) error {
	question, _ := qs.repo.GetById(questionId)

	if question.ID == 0 {
		return constants.ErrQuestionNotFound
	}

	if question.UserID != userId {
		return constants.ErrForbidden
	}

	qs.repo.Delete(question)
	return nil
}

func (qs *QuestionService) Like (questionId uint, userId uint) error {
	rating, err := qs.ratingRepo.Get(questionId, userId)
	
	if err != nil {
		_, err := qs.ratingRepo.Create(domain.UserQuestionRating{
			UserID: userId,
			QuestionID: questionId,
			IsLiked: true,
		})

		if err != nil {
			return constants.ErrGeneric
		}

		return nil
	}

	if !rating.IsLiked {
		_, err = qs.ratingRepo.Update(rating, domain.UserQuestionRating{
			IsLiked: true,
		})

		if err != nil {
			return constants.ErrGeneric
		}
	}

	return nil
}

func (qs *QuestionService) LikeUndo (questionId uint, userId uint) error {
	rating, err := qs.ratingRepo.Get(questionId, userId)

	if err != nil {
		return constants.ErrGeneric
	}

	if rating.IsLiked {
		err := qs.ratingRepo.Delete(rating)
		
		if err != nil {
			return constants.ErrGeneric
		}
	}

	return nil
}

func (qs *QuestionService) Dislike (questionId uint, userId uint) error {
	rating, err := qs.ratingRepo.Get(questionId, userId)
	
	if err != nil {
		_, err := qs.ratingRepo.Create(domain.UserQuestionRating{
			UserID: userId,
			QuestionID: questionId,
			IsLiked: false,
		})

		if err != nil {
			return constants.ErrGeneric
		}

		return nil
	}

	if rating.IsLiked {
		_, err = qs.ratingRepo.Update(rating, domain.UserQuestionRating{
			IsLiked: false,
		})

		if err != nil {
			return constants.ErrGeneric
		}
	}

	return nil
}

func (qs *QuestionService) DislikeUndo (questionId uint, userId uint) error {
	rating, err := qs.ratingRepo.Get(questionId, userId)

	if err != nil {
		return constants.ErrGeneric
	}

	if !rating.IsLiked {
		
		err := qs.ratingRepo.Delete(rating)

		if err != nil {
			return constants.ErrGeneric
		}
	}

	return nil
}

func (qs *QuestionService) CreateAnswer (questionId uint, userId uint, req requests.AnswerInsertRequest) (*responses.AnswerResponse, error) {
	answer := domain.Answer{
		QuestionID: questionId,
		UserID: userId,
		Content: req.Content,
	}

	newAnswer, err := qs.answerRepo.Create(answer)
	if err != nil {
		return nil, constants.ErrGeneric
	}

	newAnswer, _ = qs.answerRepo.GetById(newAnswer.ID)
	question, _ := qs.repo.GetById(newAnswer.QuestionID)
	response := utils.ConvertToAnswerResponseModel(newAnswer)

	if question.UserID != userId {
		notification := domain.AnswerNotification{
			UserID: question.UserID,
			AnswerID: newAnswer.ID,
			QuestionID: question.ID,
			Content: newAnswer.User.Email + " answerd you question: " + question.Content,
			IsRead: false,
		}

		notification, _ = qs.answerNotificationRepo.Create(notification)

		n := utils.ConvertToAnswerNotification(notification)

		msg, _ := json.Marshal(n)
		
		qs.pool.Broadcast <- websockets.Message{
			ClientID: uint64(question.UserID), 
			Body: string(msg),
			Type: 0,
		}
	}
	
	return &responses.AnswerResponse{Data: response}, nil
}
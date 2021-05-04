package services

import (
	"errors"

	"github.com/EnisMulic/Ask.it.Backend/constants"
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
) (*responses.AnswerResponse, error) {
	answer, _ := as.answerRepo.GetById(answerId)

	if answer.ID == 0 {
		return nil, constants.ErrAnswerNotFound
	}

	if answer.UserID != userId {
		return nil, constants.ErrUnauthorized
	}

	updatedAnswer := domain.Answer{
		Content: req.Content,
	}

	answer, err := as.answerRepo.Update(answer, updatedAnswer)

	if err != nil {
		return nil, constants.ErrGeneric
	}

	response := utils.ConvertToAnswerResponseModel(answer)
	return &responses.AnswerResponse{
		Data: response,
	}, nil
}

func (as *AnswerService) Delete(answerId uint, userId uint ) error {
	answer, err := as.answerRepo.GetById(answerId)

	if answer.ID == 0 {
		return constants.ErrAnswerNotFound
	}

	if err != nil {
		return constants.ErrGeneric
	}

	if answer.UserID != userId {
		return constants.ErrUnauthorized
	}

	err = as.answerRepo.Delete(answer)

	if err != nil {
		return constants.ErrGeneric
	}

	return nil
}

func (as *AnswerService) Like (answerId uint, userId uint) error {
	rating, err := as.ratingRepo.Get(answerId, userId)
	
	if err != nil {
		_, err := as.ratingRepo.Create(domain.UserAnswerRating{
			UserID: userId,
			AnswerID: answerId,
			IsLiked: true,
		})

		if err != nil {
			return constants.ErrGeneric
		}

		return nil
	}

	if !rating.IsLiked {
		_, err = as.ratingRepo.Update(rating, domain.UserAnswerRating{
			IsLiked: true,
		})

		if err != nil {
			return constants.ErrGeneric
		}
	}

	return nil
}

func (as *AnswerService) LikeUndo (answerId uint, userId uint) error {
	rating, err := as.ratingRepo.Get(answerId, userId)

	if err != nil {
		return constants.ErrGeneric
	}

	if rating.IsLiked {
		err := as.ratingRepo.Delete(rating)
		
		if err != nil {
			return constants.ErrGeneric
		}
	}

	return nil
}

func (as *AnswerService) Dislike (answerId uint, userId uint) error {
	rating, err := as.ratingRepo.Get(answerId, userId)
	
	if err != nil {
		_, err := as.ratingRepo.Create(domain.UserAnswerRating{
			UserID: userId,
			AnswerID: answerId,
			IsLiked: false,
		})

		if err != nil {
			return constants.ErrGeneric
		}

		return nil
	}

	if rating.IsLiked {
		_, err = as.ratingRepo.Update(rating, domain.UserAnswerRating{
			IsLiked: false,
		})

		if err != nil {
			return constants.ErrGeneric
		}
	}

	return nil
}

func (as *AnswerService) DislikeUndo (answerId uint, userId uint) error {
	rating, err := as.ratingRepo.Get(answerId, userId)

	if err != nil {
		return constants.ErrGeneric
	}

	if !rating.IsLiked {
		
		err := as.ratingRepo.Delete(rating)

		if err != nil {
			return constants.ErrGeneric
		}
	}

	return nil
}
package utils

import (
	"github.com/EnisMulic/Ask.it.Backend/contracts/responses"
	"github.com/EnisMulic/Ask.it.Backend/domain"
)

func ConvertToUserResponseModel(user domain.User) responses.UserResponseModel {
	return responses.UserResponseModel{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		AnswerCount: user.AnswerCount,
	}
}

func ConvertToUserPersonalInfoResponseModel(user domain.User) responses.UserPersonalInfoResponseModel {
	return responses.UserPersonalInfoResponseModel{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		AnswerCount: user.AnswerCount,
		QuestionRatings: convertToUserQuestionRatingsModel(user.UserQuestionRatings),
		AnswerRatings: convertToUserAnswerRatingsModel(user.UserAnswerRatings),
	}
}

func convertToUserQuestionRatingsModel(ratings []domain.UserQuestionRating) []responses.UserQuestionRatingModel {
	var ratingModels []responses.UserQuestionRatingModel
	for _, rating := range ratings {
		ratingModel := responses.UserQuestionRatingModel{
			QuestionID: rating.QuestionID,
			IsLiked: rating.IsLiked,
		}

		ratingModels = append(ratingModels, ratingModel)
	}

	return ratingModels
}

func convertToUserAnswerRatingsModel(ratings []domain.UserAnswerRating) []responses.UserAnswerRatingModel {
	var ratingModels []responses.UserAnswerRatingModel
	for _, rating := range ratings {
		ratingModel := responses.UserAnswerRatingModel{
			AnswerID: rating.AnswerID,
			IsLiked: rating.IsLiked,
		}

		ratingModels = append(ratingModels, ratingModel)
	}

	return ratingModels
}

func ConvertToQuestionResponseModel(question domain.Question) responses.QuestionResponseModel {
	return responses.QuestionResponseModel{
		ID: question.ID,
		Content: question.Content,
		CreatedAt: question.CreatedAt,
		Likes: question.Likes,
		Dislikes: question.Dislikes,
		User: ConvertToUserResponseModel(question.User),
		Answers: ConvertToAnswerResponseModels(question.Answers),
	}
}

func ConvertToAnswerResponseModel(answer domain.Answer) responses.AnswerResponseModel {
	return responses.AnswerResponseModel{
		ID: answer.ID,
		User: ConvertToUserResponseModel(answer.User),
		Content: answer.Content,
		CreatedAt: answer.CreatedAt,
		UpdatedAt: answer.UpdatedAt,
		Likes: answer.Likes,
		Dislikes: answer.Dislikes,
	}
}

func ConvertToAnswerResponseModels(answers []domain.Answer) []responses.AnswerResponseModel {
	var list []responses.AnswerResponseModel
	for _, answer := range answers {
		answerModel := ConvertToAnswerResponseModel(answer)
		list = append(list, answerModel)
	}
	
	return list
}
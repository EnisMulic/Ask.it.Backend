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
	}
}

func ConvertToQuestionResponseModel(question domain.Question) responses.QuestionResponseModel {
	return responses.QuestionResponseModel{
		ID: question.ID,
		Content: question.Content,
		CreatedAt: question.CreatedAt,
		Likes: question.Likes,
		Dislikes: question.Dislikes,
		User: ConvertToUserResponseModel(question.User),
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
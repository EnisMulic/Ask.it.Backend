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
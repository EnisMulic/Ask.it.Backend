package services

import (
	"github.com/EnisMulic/Ask.it.Backend/contracts/requests"
	"github.com/EnisMulic/Ask.it.Backend/contracts/responses"
	"github.com/EnisMulic/Ask.it.Backend/repositories"
)


type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo}
} 

func (us *UserService) Get (search requests.UserSearchRequest) responses.UsersResponse {
	users := us.repo.GetPaged(search)

	var response []responses.UserResponseModel
	for _, user := range users {
		userResponse := responses.UserResponseModel{
			ID: user.ID,
			FirstName: user.FirstName,
			LastName: user.LastName,
			Email: user.Email,
		}

		response = append(response, userResponse)
	}

	return responses.UsersResponse{Data: response}
}
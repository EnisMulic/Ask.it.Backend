package services

import (
	"errors"
	"strings"

	"github.com/EnisMulic/Ask.it.Backend/constants"
	"github.com/EnisMulic/Ask.it.Backend/contracts/requests"
	"github.com/EnisMulic/Ask.it.Backend/contracts/responses"
	"github.com/EnisMulic/Ask.it.Backend/domain"
	"github.com/EnisMulic/Ask.it.Backend/repositories"
	"github.com/EnisMulic/Ask.it.Backend/utils"
)

var ErrorUserNotFound error = errors.New("user not found")
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
		userResponse := utils.ConvertToUserResponseModel(user)

		response = append(response, userResponse)
	}

	return responses.UsersResponse{Data: response}
}

func (us *UserService) GetById (id uint) (*responses.UserResponse, *responses.ErrorResponse) {
	user := us.repo.GetById(id)

	if user.ID == 0 {
		err := responses.ErrorResponseModel{
			FieldName: "",
			Message: ErrorUserNotFound.Error(),
		}

		errors := responses.NewErrorResponse(err)	

		return nil, errors
	}

	response := utils.ConvertToUserResponseModel(user)

	return &responses.UserResponse{Data: response}, nil
}

func (us *UserService) Update(id uint, req requests.UserUpdateRequest) (*responses.UserResponse, *responses.ErrorResponse){

	updatedUser := domain.User{
		FirstName: req.FirstName,
		LastName: req.LastName,
		Email: req.Email,
	}

	user, err := us.repo.Update(id, updatedUser)

	if err != nil {
		strErr := err.Error()
		if strings.Contains(strErr, "Duplicate entry") && strings.Contains(strErr, "email") {
			err := responses.ErrorResponseModel{
				FieldName: "email",
				Message: constants.EmailIsTakenError,
			}

			errors := responses.NewErrorResponse(err)	

			return nil, errors
		}
	}

	response := utils.ConvertToUserResponseModel(user)

	return &responses.UserResponse{Data: response}, nil
}
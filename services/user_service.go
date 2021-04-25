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
	userRepo *repositories.UserRepository
	questionRepo *repositories.QuestionRepository
}

func NewUserService(userRepo *repositories.UserRepository, questionRepo *repositories.QuestionRepository) *UserService {
	return &UserService{userRepo, questionRepo}
} 

func (us *UserService) Get (search requests.UserSearchRequest) responses.UsersResponse {
	users := us.userRepo.GetPaged(search)

	var response []responses.UserResponseModel
	for _, user := range users {
		userResponse := utils.ConvertToUserResponseModel(user)

		response = append(response, userResponse)
	}

	return responses.UsersResponse{Data: response}
}

func (us *UserService) GetById (id uint) (*responses.UserResponse, *responses.ErrorResponse) {
	user := us.userRepo.GetById(id)

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
	user := us.userRepo.GetById(id)

	if user.ID == 0 {
		err := responses.ErrorResponseModel{
			FieldName: "",
			Message: ErrorUserNotFound.Error(),
		}

		errors := responses.NewErrorResponse(err)	

		return nil, errors
	}

	updatedUser := domain.User{
		FirstName: req.FirstName,
		LastName: req.LastName,
		Email: req.Email,
	}

	user, err := us.userRepo.Update(user, updatedUser)

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

func (us *UserService) ChangePassword(id uint, req requests.ChangePasswordRequest) *responses.ErrorResponse {
	user := us.userRepo.GetById(id)

	if user.ID == 0 {
		err := responses.ErrorResponseModel{
			FieldName: "",
			Message: ErrorUserNotFound.Error(),
		}

		errors := responses.NewErrorResponse(err)	

		return errors
	}

	if !doPasswordsMatch(user.PasswordHash, req.Password, user.PasswordSalt) {
		err := responses.ErrorResponseModel{
			FieldName: "password",
			Message: constants.ErrorWrongPassword,
		}

		errors := responses.NewErrorResponse(err)	

		return errors
	}

	salt := generateRandomSalt(saltSize)
	hash := hashPassword(req.NewPassword, salt)


	updatedUser := domain.User{
		PasswordSalt: salt,
		PasswordHash: hash,
	}

	user, err := us.userRepo.ChangePassword(user, updatedUser)

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

func (us *UserService) GetQuestions (userId uint, search requests.QuestionSearchRequest) (*responses.QuestionsReponse, *responses.ErrorResponse) {
	user := us.userRepo.GetById(userId)

	if user.ID == 0 {
		err := responses.ErrorResponseModel{
			FieldName: "",
			Message: ErrorUserNotFound.Error(),
		}

		errors := responses.NewErrorResponse(err)	

		return nil, errors
	}

	var filter repositories.QuestionFilter

	if (search != requests.QuestionSearchRequest{}) {
		filter = repositories.QuestionFilter{
			PageNumber: search.PageNumber,
			PageSize: search.PageSize,
		}
	}

	filter.UserID = userId

	questions := us.questionRepo.GetPaged(filter)

	var response []responses.QuestionResponseModel
	for _, question := range questions {
		questionResponse := utils.ConvertToQuestionResponseModel(question)
		response = append(response, questionResponse)
	}

	return &responses.QuestionsReponse{Data: response}, nil
}
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
	var filter repositories.UserFilter
	var pagination repositories.PaginationFilter

	if (search.PaginationQuery != nil) {
		pagination = repositories.PaginationFilter{
			PageNumber: search.PageNumber,
			PageSize: search.PageSize,
		}
	}

	users, count := us.userRepo.GetPaged(filter, pagination)

	var response []responses.UserResponseModel
	for _, user := range users {
		userResponse := utils.ConvertToUserResponseModel(user)

		response = append(response, userResponse)
	}

	if (search != requests.UserSearchRequest{}) {
		return utils.CreateUserPagedResponse(response, count, int64(search.PageNumber), int64(search.PageSize))
	} else {
		return utils.CreateUserPagedResponse(response, count, 0, 0)
	}
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

	filter := repositories.QuestionFilter{
		UserID: userId,
	}

	var pagination repositories.PaginationFilter

	if (search.PaginationQuery != nil) {
		pagination = repositories.PaginationFilter{
			PageNumber: search.PageNumber,
			PageSize: search.PageSize,
		}
	}

	questions, count := us.questionRepo.GetPaged(filter, pagination)

	var response []responses.QuestionResponseModel
	for _, question := range questions {
		questionResponse := utils.ConvertToQuestionResponseModel(question)
		response = append(response, questionResponse)
	}

	if (search != requests.QuestionSearchRequest{}) {
		return utils.CreateQuestionPagedResponse(response, count, int64(search.PageNumber), int64(search.PageSize)), nil
	} else {
		return utils.CreateQuestionPagedResponse(response, count, 0, 0), nil
	}
}
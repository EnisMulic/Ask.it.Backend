package services

import (
	"strings"

	"github.com/EnisMulic/Ask.it.Backend/constants"
	"github.com/EnisMulic/Ask.it.Backend/contracts/requests"
	"github.com/EnisMulic/Ask.it.Backend/contracts/responses"
	"github.com/EnisMulic/Ask.it.Backend/domain"
	"github.com/EnisMulic/Ask.it.Backend/repositories"
	"github.com/EnisMulic/Ask.it.Backend/utils"
)


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

	users, count := us.userRepo.GetPaged(filter, []repositories.SortFilter{}, pagination)

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

func (us *UserService) GetTop (search requests.UserSearchRequest) responses.UsersResponse {
	sort := repositories.SortFilter{
		Column: "answer_count",
		Order: "desc",
	}

	var filter repositories.UserFilter
	var pagination repositories.PaginationFilter

	if (search.PaginationQuery != nil) {
		pagination = repositories.PaginationFilter{
			PageNumber: search.PageNumber,
			PageSize: search.PageSize,
		}
	}

	users, count := us.userRepo.GetPaged(filter, []repositories.SortFilter{sort}, pagination)

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

func (us *UserService) GetById (id uint) (*responses.UserResponse, error) {
	user := us.userRepo.GetById(id)

	if user.ID == 0 {
		// err := responses.ErrorResponseModel{
		// 	FieldName: "",
		// 	Message: ErrorUserNotFound.Error(),
		// }

		// errors := responses.NewErrorResponse(err)	

		return nil, constants.ErrUserNotFound
	}

	response := utils.ConvertToUserResponseModel(user)

	return &responses.UserResponse{Data: response}, nil
}

func (us *UserService) GetPersonalInfo (id uint) (*responses.UserPersonalInfoResponse, error) {
	user := us.userRepo.GetPersonalInfo(id)

	if user.ID == 0 {
		return nil, constants.ErrUserNotFound
	}
	response := utils.ConvertToUserPersonalInfoResponseModel(user)

	return &responses.UserPersonalInfoResponse{Data: response}, nil
}

func (us *UserService) Update(id uint, req requests.UserUpdateRequest) (*responses.UserResponse, error){
	user := us.userRepo.GetById(id)

	if user.ID == 0 {
		return nil, constants.ErrUserNotFound
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
			return nil, constants.ErrEmailIsTaken
		}
	}

	response := utils.ConvertToUserResponseModel(user)

	return &responses.UserResponse{Data: response}, nil
}

func (us *UserService) ChangePassword(id uint, req requests.ChangePasswordRequest) error {
	user := us.userRepo.GetById(id)

	if user.ID == 0 {
		return constants.ErrUserNotFound
	}

	if !doPasswordsMatch(user.PasswordHash, req.Password, user.PasswordSalt) {
		return constants.ErrWrongPassword
	}

	salt := generateRandomSalt(saltSize)
	hash := hashPassword(req.NewPassword, salt)


	updatedUser := domain.User{
		PasswordSalt: salt,
		PasswordHash: hash,
	}

	user, err := us.userRepo.ChangePassword(user, updatedUser)

	if err != nil {
		return constants.ErrGeneric
	}

	return nil
}

func (us *UserService) GetQuestions (userId uint, search requests.QuestionSearchRequest) (*responses.QuestionsReponse, error) {
	user := us.userRepo.GetById(userId)

	if user.ID == 0 {
		return nil, constants.ErrUserNotFound
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

	sort := repositories.SortFilter{
		Column: "created_at",
		Order: "desc",
	}

	questions, count := us.questionRepo.GetPaged(filter, []repositories.SortFilter{sort}, pagination)

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
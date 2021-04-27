package responses

type UserResponseModel struct {
	ID          uint
	FirstName   string
	LastName    string
	Email       string
	AnswerCount int
}

// swagger:response UserResponse
type UserResponse struct {
	// in: body
	Data UserResponseModel
}

// swagger:response UsersResponse
type UsersResponse struct {
	// in: body
	Data []UserResponseModel
	// in: body
	PagedResponse
}

type UserQuestionRatingModel struct {
	QuestionID uint
	IsLiked    bool
}

type UserAnswerRatingModel struct {
	AnswerID uint
	IsLiked  bool
}

type UserPersonalInfoResponseModel struct {
	ID              uint
	FirstName       string
	LastName        string
	Email           string
	AnswerCount     int
	QuestionRatings []UserQuestionRatingModel
	AnswerRatings   []UserAnswerRatingModel
}

type UserPersonalInfoResponse struct {
	Data UserPersonalInfoResponseModel
}
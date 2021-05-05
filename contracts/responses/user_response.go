package responses

type UserResponseModel struct {
	ID          uint   `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	AnswerCount int    `json:"answerCount"`
}

// swagger:response UserResponse
type UserResponse struct {
	// in: body
	Data UserResponseModel `json:"data"`
}

// swagger:response UsersResponse
type UsersResponse struct {
	// in: body
	Data []UserResponseModel `json:"data"`
	// in: body
	PagedResponse
}

type UserQuestionRatingModel struct {
	QuestionID uint `json:"questionId"`
	IsLiked    bool `json:"isLiked"`
}

type UserAnswerRatingModel struct {
	AnswerID uint `json:"answerId"`
	IsLiked  bool `json:"isLiked"`
}

type UserPersonalInfoResponseModel struct {
	ID                  uint                      `json:"id"`
	FirstName           string                    `json:"firstName"`
	LastName            string                    `json:"lastName"`
	Email               string                    `json:"email"`
	AnswerCount         int                       `json:"answerCount"`
	QuestionRatings     []UserQuestionRatingModel `json:"questionRatings"`
	AnswerRatings       []UserAnswerRatingModel   `json:"answerRatings"`
	AnswerNotifications []AnswerNotification      `json:"answerNotifications"`
}

type UserPersonalInfoResponse struct {
	Data UserPersonalInfoResponseModel `json:"data"`
}
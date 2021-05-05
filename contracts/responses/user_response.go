package responses

// swagger:model UserResponseModel
type UserResponseModel struct {
	ID          uint   `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	AnswerCount int    `json:"answerCount"`
}

// swagger:model UserResponse
type UserResponse struct {
	// in: body
	Data UserResponseModel `json:"data"`
}

// swagger:model UsersResponse
type UsersResponse struct {
	// in: body
	Data []UserResponseModel `json:"data"`
	// in: body
	PagedResponse
}

// swagger:model UserQuestionRatingModel
type UserQuestionRatingModel struct {
	QuestionID uint `json:"questionId"`
	IsLiked    bool `json:"isLiked"`
}

// swagger:model UserAnswerRatingModel
type UserAnswerRatingModel struct {
	AnswerID uint `json:"answerId"`
	IsLiked  bool `json:"isLiked"`
}

// swagger:model UserPersonalInfoResponseModel
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

// swagger:model UserPersonalInfoResponse
type UserPersonalInfoResponse struct {
	Data UserPersonalInfoResponseModel `json:"data"`
}
package responses

import "time"

// swagger:model AnswerResponseModel
type AnswerResponseModel struct {
	ID        uint `json:"id"`
	Content   string `json:"content"`
	Likes     int `json:"likes"`
	Dislikes  int `json:"dislikes"`
	User UserResponseModel `json:"user"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// swagger:model AnswerResponse
type AnswerResponse struct {
	// in: body
	Data AnswerResponseModel `json:"data"`
}

// swagger:model AnswersResponse
type AnswersResponse struct {
	// in: body
	Date []AnswerResponseModel `json:"data"`
	// in: body
	PagedResponse
}


package responses

import "time"

type AnswerResponseModel struct {
	ID        uint `json:"id"`
	Content   string `json:"content"`
	Likes     int `json:"likes"`
	Dislikes  int `json:"dislikes"`
	User UserResponseModel `json:"user"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// swagger:response AnswerResponse
type AnswerResponse struct {
	// in: body
	Data AnswerResponseModel `json:"data"`
}

// swagger:response AnswersResponse
type AnswersResponse struct {
	// in: body
	Date []AnswerResponseModel `json:"data"`
	// in: body
	PagedResponse
}


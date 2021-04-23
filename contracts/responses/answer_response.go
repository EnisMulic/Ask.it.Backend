package responses

import "time"

type AnswerResponseModel struct {
	ID        uint
	Content   string
	Likes     int
	Dislikes  int
	User UserResponseModel
	CreatedAt time.Time
	UpdatedAt time.Time
}

// swagger:response AnswerResponse
type AnswerResponse struct {
	// in: body
	Data AnswerResponseModel
}

// swagger:response AnswersResponse
type AnswersResponse struct {
	// in: body
	Date []AnswerResponseModel
	// in: body
	PagedResponse
}


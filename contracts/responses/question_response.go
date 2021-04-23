package responses

import "time"

type QuestionResponseModel struct {
	ID        uint
	Content   string
	Likes     int
	Dislikes  int
	CreatedAt time.Time
	UpdatedAt time.Time
	User UserResponseModel
	Answers []AnswerResponseModel
}


// swagger:response QuestionResponse
type QuestionResponse struct {
	// in: body
	Data QuestionResponseModel
}

// swagger:response QuestionsResponse
type QuestionsReponse struct {
	// in: body
	Data []QuestionResponseModel
	// in: body
	PagedResponse
}
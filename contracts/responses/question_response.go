package responses

import "time"

type QuestionResponseModel struct {
	ID        uint `json:"id"`
	Content   string `json:"content"`
	Likes     int `json:"likes"`
	Dislikes  int `json:"dislikes"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	User UserResponseModel `json:"user"`
	Answers []AnswerResponseModel `json:"answers"`
}


// swagger:response QuestionResponse
type QuestionResponse struct {
	// in: body
	Data QuestionResponseModel `json:"data"`
}

// swagger:response QuestionsResponse
type QuestionsReponse struct {
	// in: body
	Data []QuestionResponseModel `json:"data"`
	// in: body
	PagedResponse
}
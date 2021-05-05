package responses

import "time"

// swagger:model QuestionResponseModel
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


// swagger:model QuestionResponse
type QuestionResponse struct {
	// in: body
	Data QuestionResponseModel `json:"data"`
}

// swagger:model QuestionsResponse
type QuestionsReponse struct {
	// in: body
	Data []QuestionResponseModel `json:"data"`
	// in: body
	PagedResponse
}
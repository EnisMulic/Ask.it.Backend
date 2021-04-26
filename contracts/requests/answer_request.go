package requests

type AnswerInsertRequest struct {
	Content string `schema:"content"`
}

type AnswerUpdateRequest struct {
	Content string `schema:"content"`
}
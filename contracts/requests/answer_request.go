package requests

type AnswerInsertRequest struct {
	Content string `json:"content"`
}

// swagger:parameters answerInsert
type AnswerInsertRequestWrapper struct {
	// in: body
	Body AnswerInsertRequest
}

type AnswerUpdateRequest struct {
	Content string `json:"content"`
}

// swagger:parameters answerUpdate
type AnswerUpdateRequestWrapper struct {
	// in: body
	Body AnswerUpdateRequest
}
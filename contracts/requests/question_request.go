package requests

// swagger:parameters questionSearch
type QuestionSearchRequest struct {
	*PaginationQuery
}

type QuestionInsertRequest struct {
	Content string `json:"content"`
}

// swagger:parameters questionInsert
type QuestionInsertRequestWrapper struct {
	// in: body
	Body QuestionInsertRequest
}
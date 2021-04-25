package requests

type QuestionSearchRequest struct {
	*PaginationQuery
}

type QuestionInsertRequest struct {
	Content string `schema:"content"`
}
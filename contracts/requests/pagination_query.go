package requests

// swagger:model PaginationQuery
type PaginationQuery struct {
	PageNumber int `schema:"pageNumber"`
	PageSize   int `schame:"pageSize"`
}
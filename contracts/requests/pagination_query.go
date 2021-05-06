package requests

// swagger:parameters pagination
type PaginationQuery struct {
	PageNumber int `json:"pageNumber"`
	PageSize   int `json:"pageSize"`
}
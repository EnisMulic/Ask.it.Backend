package responses

// swagger:model PagedResponse
type PagedResponse struct {
	// in: body
	PageNumber   int64 `json:"pageNumber"`
	// in: body
	PageSize     int64 `json:"pageSize"`
	// in: body
	NextPage     *int64 `json:"nextPage"`
	// in: body 
	PreviousPage *int64 `json:"previousPage"`
	// in: body
	FirstPage    int64 `json:"firstPage"`
	// in: body
	LastPage     int64 `json:"lastPage"`
}
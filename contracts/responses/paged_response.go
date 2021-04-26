package responses

// swagger:response PagedResponse
type PagedResponse struct {
	// in: body
	PageNumber   int64
	// in: body
	PageSize     int64
	// in: body
	NextPage     *int64
	// in: body
	PreviousPage *int64
	// in: body
	FirstPage    int64
	// in: body
	LastPage     int64
}
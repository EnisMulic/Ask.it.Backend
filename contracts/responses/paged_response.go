package responses

// swagger:response PagedResponse
type PagedResponse struct {
	// in: body
	PageNumber   int
	// in: body
	PageSize     int
	// in: body
	NextPage     int
	// in: body
	PreviousPage int
	// in: body
	FirstPage    int
	// in: body
	LastPage     int
}
package responses

type UserResponseModel struct {
	ID        uint
	FirstName string
	LastName  string
	Email     string
}

// swagger:response UserResponse
type UserResponse struct {
	// in: body
	Data UserResponseModel
}

// swagger:response UsersResponse
type UsersResponse struct {
	// in: body
	Data []UserResponseModel
	// in: body
	PagedResponse
}
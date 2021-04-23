package requests

// swagger:parameters UserSearchRequest
type UserSearchRequest struct {
	*PaginationQuery
}

// swagger:model RegisterRequest
type RegisterRequest struct {
	FirstName string `schema:"firstName"`
	LastName string `schema:"lastName"`
	Email string `schema:"email"`
	Password string `schema:"password"`
}

type LoginRequest struct {
	Email string `schema:"email"`
	Password string `schema:"password"`
}
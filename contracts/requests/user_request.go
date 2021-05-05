package requests

// swagger:parameters UserSearchRequest
type UserSearchRequest struct {
	*PaginationQuery
}

type UserUpdateRequest struct {
	FirstName string `schema:"firstName"`
	LastName string `schema:"lastName"`
	Email string `schema:"email"`
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

// swagger:parameters changePassword
type ChangePasswordRequest struct {
	// in: body
	Password string `json:"password"`
	// in: body
	NewPassword string `json:"newPassword"`
}
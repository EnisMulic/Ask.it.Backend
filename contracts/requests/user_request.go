package requests

// swagger:parameters usersSearch topUsersSearch
type UserSearchRequest struct {
	*PaginationQuery
}

type UserUpdateRequest struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Email string `json:"email"`
}

// swagger:parameters userUpdate
type UserUpdateRequestWrapper struct {
	// in: body
	Body UserUpdateRequest
}

type RegisterRequest struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Email string `json:"email"`
	Password string `json:"password"`
}

// swagger:parameters register
type RegisterRequestWrapper struct {
	// in: body
	Body RegisterRequest
}

type LoginRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

// swagger:parameters login
type LoginRequestWrapper struct {
	// in: body
	Body LoginRequest
}

type ChangePasswordRequest struct {
	Password string `json:"password"`
	NewPassword string `json:"newPassword"`
}

// swagger:parameters changePassword
type ChangePasswordRequestWrapper struct {
	// in: body
	Body ChangePasswordRequest
}
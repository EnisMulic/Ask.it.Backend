package controllers

import (
	"log"
	"net/http"

	"github.com/EnisMulic/Ask.it.Backend/services"
)

type UserController struct {
	l *log.Logger
	us *services.UserService
}

func NewUserController(l *log.Logger, us *services.UserService) *UserController {
	return &UserController{l, us}
}

// swagger:route GET /api/users users users
// Returns a list of users
//
// responses:
//	200: UsersResponse
func (uc *UserController) Get(rw http.ResponseWriter, r *http.Request) {

}

// swagger:route GET /api/users/{id} users user
// Returns a single user
//
// responses:
//	200: UserResponse
func (uc *UserController) GetById(rw http.ResponseWriter, r *http.Request) {

}

// swagger:route GET /api/users/me users user
// Returns information of the user that called the route
//
// responses:
//	200: UserResponse
func (uc *UserController) GetMe(rw http.ResponseWriter, r *http.Request) {

}

// swagger:route POST /api/users/change-password users bool
// Change users password
func (uc *UserController) ChangePassword(rw http.ResponseWriter, r *http.Request) {

}

// swagger:route PUT /api/users users user
//
// responses:
//	200: UserResponse
func (uc *UserController) Update(rw http.ResponseWriter, r *http.Request) {

}

// swagger:route GET /api/users/{id}/questions users questions
//
// responses:
//	200: QuestionsResponse
func (uc *UserController) GetQuestions(rw http.ResponseWriter, r *http.Request) {

}
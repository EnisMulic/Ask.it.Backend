package controllers

import (
	"log"
	"net/http"
)

type UserController struct {
	l *log.Logger
}

func NewUserController(l *log.Logger) *UserController {
	return &UserController{l}
}

// swagger:route GET /api/users users users
// Returns a list of users
func (uc *UserController) Get(rw http.ResponseWriter, r *http.Request) {

}

// swagger:route GET /api/users/{id} users user
// Returns a single user
func (uc *UserController) GetById(rw http.ResponseWriter, r *http.Request) {

}

// swagger:route GET /api/users/me users user
// Returns information of the user that called the route
func (uc *UserController) GetMe(rw http.ResponseWriter, r *http.Request) {

}

// swagger:route POST /api/users/change-password users bool
// Change users password
func (uc *UserController) ChangePassword(rw http.ResponseWriter, r *http.Request) {

}

// swagger:route PUT /api/users users user
//
func (uc *UserController) Update(rw http.ResponseWriter, r *http.Request) {

}

// swagger:route GET /api/users/{id}/questions users questions
//
func (uc *UserController) GetQuestions(rw http.ResponseWriter, r *http.Request) {

}
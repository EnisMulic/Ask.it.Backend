package controllers

import (
	"log"
	"net/http"
)

type AuthController struct {
	l *log.Logger
}

func NewAuthController(l *log.Logger) *AuthController {
	return &AuthController{l}
}

// swagger:route POST /api/auth/login auth jwt
//
func (ac *AuthController) Login(rw http.ResponseWriter, r *http.Request) {

}

// swagger:route POST /api/auth/register auth jwt
//
func (ac *AuthController) Register(rw http.ResponseWriter, r *http.Request) {
	
}

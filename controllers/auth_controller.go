package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/EnisMulic/Ask.it.Backend/constants"
	"github.com/EnisMulic/Ask.it.Backend/contracts/requests"
	"github.com/EnisMulic/Ask.it.Backend/services"
)

type AuthController struct {
	l *log.Logger
	as *services.AuthService
}

func NewAuthController(l *log.Logger, as *services.AuthService) *AuthController {
	return &AuthController{l, as}
}

// swagger:route POST /api/auth/login auth jwt
//
// responses:
//	200: AuthResponse
func (ac *AuthController) Login(rw http.ResponseWriter, r *http.Request) {
	var req requests.LoginRequest

	decoder := json.NewDecoder(r.Body)
	
	err := decoder.Decode(&req)
    if err != nil {
		http.Error(rw, constants.UnableToParseJSONBody, http.StatusBadRequest)
		return
    } 

	res, resErr := ac.as.Login(req)
	if resErr != nil {
		http.Error(rw, "", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(rw).Encode(res)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

// swagger:route POST /api/auth/register auth jwt
//
// parameters:
// + name: request
//   in: body
//   schema:
//    "$ref": "#/definitions/RegisterRequest"
//
// responses:
//	200: AuthResponse
func (ac *AuthController) Register(rw http.ResponseWriter, r *http.Request) {
	var req requests.RegisterRequest

	decoder := json.NewDecoder(r.Body)
	
	err := decoder.Decode(&req)
    if err != nil {
		http.Error(rw, constants.UnableToParseJSONBody, http.StatusBadRequest)
		return
    } 

	res, resErr := ac.as.Register(req)
	if resErr != nil {
		http.Error(rw, "", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(rw).Encode(res)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

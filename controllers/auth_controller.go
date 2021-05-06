package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/EnisMulic/Ask.it.Backend/constants"
	"github.com/EnisMulic/Ask.it.Backend/contracts/requests"
	"github.com/EnisMulic/Ask.it.Backend/contracts/responses"
	"github.com/EnisMulic/Ask.it.Backend/services"
	"github.com/EnisMulic/Ask.it.Backend/utils"
)

type AuthController struct {
	l *log.Logger
	as *services.AuthService
}

func NewAuthController(l *log.Logger, as *services.AuthService) *AuthController {
	return &AuthController{l, as}
}

// swagger:route POST /api/auth/login auth login
//
// responses:
//	200: AuthResponse
//  400: ErrorResponse
//  500: ErrorResponse
func (ac *AuthController) Login(rw http.ResponseWriter, r *http.Request) {
	var req requests.LoginRequest

	decoder := json.NewDecoder(r.Body)
	
	err := decoder.Decode(&req)
    if err != nil {
		errors := responses.NewErrorResponse(responses.ErrorResponseModel{
			Message: constants.ErrMsgUnableToParseJSONBody,
		})

		_ = json.NewEncoder(rw).Encode(errors)
		rw.WriteHeader(http.StatusBadRequest)

		return
    } 

	res, err := ac.as.Login(req)
	if err != nil {
		errRes := utils.ConvertToErrorResponse(err)
		
		_ = json.NewEncoder(rw).Encode(errRes)
		rw.WriteHeader(http.StatusBadRequest)

		return
	}

	err = json.NewEncoder(rw).Encode(res)
	if err != nil {
		errors := responses.NewErrorResponse(responses.ErrorResponseModel{
			Message: ErrorUnableToMarshalJson.Error(),
		})

		_ = json.NewEncoder(rw).Encode(errors)
		rw.WriteHeader(http.StatusInternalServerError)

		return
	}
}

// swagger:route POST /api/auth/register auth register
//
// responses:
//	200: AuthResponse
//  400: ErrorResponse
//  500: ErrorResponse
func (ac *AuthController) Register(rw http.ResponseWriter, r *http.Request) {
	var req requests.RegisterRequest

	decoder := json.NewDecoder(r.Body)
	
	err := decoder.Decode(&req)
    if err != nil {
		http.Error(rw, constants.ErrMsgUnableToParseJSONBody, http.StatusBadRequest)
		return
    } 

	res, err := ac.as.Register(req)
	if err != nil {
		errRes := utils.ConvertToErrorResponse(err)
		
		_ = json.NewEncoder(rw).Encode(errRes)
		rw.WriteHeader(http.StatusBadRequest)

		return
	}

	err = json.NewEncoder(rw).Encode(res)
	if err != nil {
		errors := responses.NewErrorResponse(responses.ErrorResponseModel{
			Message: ErrorUnableToMarshalJson.Error(),
		})

		_ = json.NewEncoder(rw).Encode(errors)
		rw.WriteHeader(http.StatusInternalServerError)

		return
	}
}

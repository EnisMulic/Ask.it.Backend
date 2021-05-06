package controllers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/EnisMulic/Ask.it.Backend/constants"
	"github.com/EnisMulic/Ask.it.Backend/contracts/requests"
	"github.com/EnisMulic/Ask.it.Backend/contracts/responses"
	"github.com/EnisMulic/Ask.it.Backend/services"
	"github.com/EnisMulic/Ask.it.Backend/utils"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

var ErrorUnableToMarshalJson = errors.New("unable to marshal json")

var decoder = schema.NewDecoder()

type UserController struct {
	l *log.Logger
	us *services.UserService
}

func NewUserController(l *log.Logger, us *services.UserService) *UserController {
	return &UserController{l, us}
}

// swagger:route GET /api/users users userSearch
// Returns a list of users
//
// responses:
//	200: UsersResponse
func (uc *UserController) Get(rw http.ResponseWriter, r *http.Request) {
	var request requests.UserSearchRequest

	err := decoder.Decode(&request, r.URL.Query())
    if err != nil {
        log.Println("Error in GET parameters : ", err)
		http.Error(rw, "Unable to parse query parametars.", http.StatusBadRequest)
		return
    } 

	users := uc.us.Get(request)

	err = json.NewEncoder(rw).Encode(users)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

// swagger:route GET /api/users-top users userSearch
// Returns a list of users
//
// responses:
//	200: UsersResponse
func (uc *UserController) GetTop(rw http.ResponseWriter, r *http.Request) {
	var request requests.UserSearchRequest

	err := decoder.Decode(&request, r.URL.Query())
    if err != nil {
        log.Println("Error in GET parameters : ", err)
		http.Error(rw, "Unable to parse query parametars.", http.StatusBadRequest)
		return
    } 

	users := uc.us.GetTop(request)

	err = json.NewEncoder(rw).Encode(users)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

// swagger:route GET /api/users/{id} users user
// Returns a single user
//
// parameters:
// + name: id
//	 in: path
//	 schema: int
//
// responses:
//	200: UserResponse
//	404: ErrorResponse
func (uc *UserController) GetById(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		errors := responses.NewErrorResponse(responses.ErrorResponseModel{
			Message: constants.ErrMsgUnableToConvertId,
		})

		out, _ := json.Marshal(errors)

		http.Error(rw, string(out), http.StatusBadRequest)
		return
	}

	user, errRes := uc.us.GetById(uint(id))

	if errRes != nil {
		out, _ := json.Marshal(errRes)
		http.Error(rw, string(out), http.StatusNotFound)
		return
	}

	err = json.NewEncoder(rw).Encode(user)
	if err != nil {
		errors := responses.NewErrorResponse(responses.ErrorResponseModel{
			Message: ErrorUnableToMarshalJson.Error(),
		})

		out, _ := json.Marshal(errors)

		http.Error(rw, string(out), http.StatusInternalServerError)
	}
}

// swagger:route GET /api/me users user
// Returns information of the user that called the route
//
// security:
//  - Bearer: []
//
// responses:
//	200: UserResponse
//  400: ErrorResponse
func (uc *UserController) GetMe(rw http.ResponseWriter, r *http.Request) {
	sub, err := utils.ExtractSubFromJwt(r)

	if err != nil {
		http.Error(rw, "", http.StatusBadRequest)
		return;
	}
	
	id, err := strconv.ParseUint(sub, 10, 64)
	if err != nil {
		errors := responses.NewErrorResponse(responses.ErrorResponseModel{
			Message: constants.ErrMsgUnableToConvertId,
		})

		out, _ := json.Marshal(errors)

		http.Error(rw, string(out), http.StatusBadRequest)
		return
	}

	user, errRes := uc.us.GetPersonalInfo(uint(id))

	if errRes != nil {
		out, _ := json.Marshal(errRes)
		http.Error(rw, string(out), http.StatusNotFound)
		return
	}

	err = json.NewEncoder(rw).Encode(user)
	if err != nil {
		errors := responses.NewErrorResponse(responses.ErrorResponseModel{
			Message: ErrorUnableToMarshalJson.Error(),
		})

		out, _ := json.Marshal(errors)

		http.Error(rw, string(out), http.StatusInternalServerError)
	}
}

// swagger:route POST /api/users/change-password users changePassword
// Change users password
//
// security:
//  - Bearer: []
//
// responses:
//	200: 
//  400: ErrorResponse
func (uc *UserController) ChangePassword(rw http.ResponseWriter, r *http.Request) {
	sub, err := utils.ExtractSubFromJwt(r)

	if err != nil {
		http.Error(rw, "", http.StatusBadRequest)
		return;
	}

	id, err := strconv.ParseUint(sub, 10, 64)
	if err != nil {
		errors := responses.NewErrorResponse(responses.ErrorResponseModel{
			Message: constants.ErrMsgUnableToConvertId,
		})

		out, _ := json.Marshal(errors)

		http.Error(rw, string(out), http.StatusBadRequest)
		return
	}

	var req requests.ChangePasswordRequest

	decoder := json.NewDecoder(r.Body)
	
	err = decoder.Decode(&req)
    if err != nil {
		errors := responses.NewErrorResponse(responses.ErrorResponseModel{
			Message: constants.ErrMsgUnableToParseJSONBody,
		})

		out, _ := json.Marshal(errors)

		http.Error(rw, string(out), http.StatusBadRequest)
		return
    } 

	errRes := uc.us.ChangePassword(uint(id), req)

	if errRes != nil {
		out, _ := json.Marshal(errRes)

		http.Error(rw, string(out), http.StatusBadRequest)
		return;
	}
}

// swagger:route PUT /api/users users userUpdate
//
// security:
//  - Bearer: []
//
// responses:
//	200: UserResponse
//  400: ErrorResponse
//  404: ErrorResponse
//  500: ErrorResponse
func (uc *UserController) Update(rw http.ResponseWriter, r *http.Request) {
	sub, err := utils.ExtractSubFromJwt(r)

	if err != nil {
		http.Error(rw, "", http.StatusBadRequest)
		return;
	}

	id, err := strconv.ParseUint(sub, 10, 64)
	if err != nil {
		errors := responses.NewErrorResponse(responses.ErrorResponseModel{
			Message: constants.ErrMsgUnableToConvertId,
		})

		out, _ := json.Marshal(errors)

		http.Error(rw, string(out), http.StatusBadRequest)
		return
	}

	var req requests.UserUpdateRequest

	decoder := json.NewDecoder(r.Body)
	
	err = decoder.Decode(&req)
    if err != nil {
		errors := responses.NewErrorResponse(responses.ErrorResponseModel{
			Message: constants.ErrMsgUnableToParseJSONBody,
		})

		out, _ := json.Marshal(errors)

		http.Error(rw, string(out), http.StatusBadRequest)
		return
    } 

	user, err := uc.us.Update(uint(id), req)

	if err != nil {
		resErr := utils.ConvertToErrorResponse(err)
		out, _ := json.Marshal(resErr)

		if err == constants.ErrUserNotFound {
			http.Error(rw, string(out), http.StatusNotFound)
		} else if err == constants.ErrEmailIsTaken {
			http.Error(rw, string(out), http.StatusBadRequest)
		}

		return
	}

	err = json.NewEncoder(rw).Encode(user)
	if err != nil {
		errors := responses.NewErrorResponse(responses.ErrorResponseModel{
			Message: ErrorUnableToMarshalJson.Error(),
		})

		out, _ := json.Marshal(errors)

		http.Error(rw, string(out), http.StatusInternalServerError)
	}
}

// swagger:route GET /api/users/{id}/questions users questions
//
// security:
//  - Bearer: []
//
// parameters:
// + name: id
//	 in: path
//	 schema: int
//
// responses:
//	200: QuestionsResponse
//  400: ErrorResponse
//  500: ErrorResponse
func (uc *UserController) GetQuestions(rw http.ResponseWriter, r *http.Request) {
	var request requests.QuestionSearchRequest

	err := decoder.Decode(&request, r.URL.Query())
    if err != nil {
		errors := responses.NewErrorResponse(responses.ErrorResponseModel{
			Message: constants.ErrMsgUnableToParseQueryParametars,
		})

		out, _ := json.Marshal(errors)

		http.Error(rw, string(out), http.StatusBadRequest)
		return
    } 

	vars := mux.Vars(r)

	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		errors := responses.NewErrorResponse(responses.ErrorResponseModel{
			Message: constants.ErrMsgUnableToConvertId,
		})

		out, _ := json.Marshal(errors)

		http.Error(rw, string(out), http.StatusBadRequest)
		return
	}

	questions, err := uc.us.GetQuestions(uint(id), request)
	if err != nil {
		resErr := utils.ConvertToErrorResponse(err)
		out, _ := json.Marshal(resErr)

		http.Error(rw, string(out), http.StatusBadRequest)
		return;
	}

	err = json.NewEncoder(rw).Encode(questions)
	if err != nil {
		errors := responses.NewErrorResponse(responses.ErrorResponseModel{
			Message: ErrorUnableToMarshalJson.Error(),
		})

		out, _ := json.Marshal(errors)

		http.Error(rw, string(out), http.StatusInternalServerError)
	}
}
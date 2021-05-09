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

// swagger:route GET /api/users users usersSearch
// Returns a list of users
//
// responses:
//	200: UsersResponse
func (uc *UserController) Get(rw http.ResponseWriter, r *http.Request) {
	var request requests.UserSearchRequest

	err := decoder.Decode(&request, r.URL.Query())
    if err != nil {
        log.Println("Error in GET parameters : ", err)
		errors := responses.NewErrorResponse(responses.ErrorResponseModel{
			Message: constants.ErrMsgUnableToParseQueryParametars,
		})

		_ = json.NewEncoder(rw).Encode(errors)
		rw.WriteHeader(http.StatusBadRequest)

		return
    } 

	users := uc.us.Get(request)

	err = json.NewEncoder(rw).Encode(users)
	if err != nil {
		errors := responses.NewErrorResponse(responses.ErrorResponseModel{
			Message: ErrorUnableToMarshalJson.Error(),
		})

		_ = json.NewEncoder(rw).Encode(errors)
		rw.WriteHeader(http.StatusInternalServerError)
	}
}

// swagger:route GET /api/users-top users topUsersSearch
// Returns a list of users
//
// responses:
//	200: UsersResponse
func (uc *UserController) GetTop(rw http.ResponseWriter, r *http.Request) {
	var request requests.UserSearchRequest

	err := decoder.Decode(&request, r.URL.Query())
    if err != nil {
        log.Println("Error in GET parameters : ", err)
		errors := responses.NewErrorResponse(responses.ErrorResponseModel{
			Message: constants.ErrMsgUnableToParseQueryParametars,
		})

		_ = json.NewEncoder(rw).Encode(errors)
		rw.WriteHeader(http.StatusBadRequest)

		return
    } 

	users := uc.us.GetTop(request)

	err = json.NewEncoder(rw).Encode(users)
	if err != nil {
		errors := responses.NewErrorResponse(responses.ErrorResponseModel{
			Message: ErrorUnableToMarshalJson.Error(),
		})

		_ = json.NewEncoder(rw).Encode(errors)
		rw.WriteHeader(http.StatusInternalServerError)
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

		rw.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(rw).Encode(errors)
		
		return
	}

	user, err := uc.us.GetById(uint(id))
	
	if err != nil {
		errors := responses.NewErrorResponse(responses.ErrorResponseModel{
			Message: err.Error(),
		})

		rw.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(rw).Encode(errors)
		
		return
	}

	err = json.NewEncoder(rw).Encode(user)
	if err != nil {
		errors := responses.NewErrorResponse(responses.ErrorResponseModel{
			Message: ErrorUnableToMarshalJson.Error(),
		})

		rw.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(rw).Encode(errors)
	}
}

// swagger:route GET /api/me users loggedInUser
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
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	
	id, err := strconv.ParseUint(sub, 10, 64)
	if err != nil {
		errors := responses.NewErrorResponse(responses.ErrorResponseModel{
			Message: constants.ErrMsgUnableToConvertId,
		})

		rw.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(rw).Encode(errors)

		return
	}

	user, errRes := uc.us.GetPersonalInfo(uint(id))

	if errRes != nil {
		rw.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(rw).Encode(errRes)

		return
	}

	err = json.NewEncoder(rw).Encode(user)
	if err != nil {
		errors := responses.NewErrorResponse(responses.ErrorResponseModel{
			Message: ErrorUnableToMarshalJson.Error(),
		})

		rw.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(rw).Encode(errors)
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
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseUint(sub, 10, 64)
	if err != nil {
		errors := responses.NewErrorResponse(responses.ErrorResponseModel{
			Message: constants.ErrMsgUnableToConvertId,
		})

		rw.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(rw).Encode(errors)

		return
	}

	var req requests.ChangePasswordRequest

	decoder := json.NewDecoder(r.Body)
	
	err = decoder.Decode(&req)
    if err != nil {
		errors := responses.NewErrorResponse(responses.ErrorResponseModel{
			Message: constants.ErrMsgUnableToParseJSONBody,
		})

		rw.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(rw).Encode(errors)

		return
    } 

	errRes := uc.us.ChangePassword(uint(id), req)

	if errRes != nil {
		rw.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(rw).Encode(errRes)

		return
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
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseUint(sub, 10, 64)
	if err != nil {
		errors := responses.NewErrorResponse(responses.ErrorResponseModel{
			Message: constants.ErrMsgUnableToConvertId,
		})

		rw.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(rw).Encode(errors)

		return
	}

	var req requests.UserUpdateRequest

	decoder := json.NewDecoder(r.Body)
	
	err = decoder.Decode(&req)
    if err != nil {
		errors := responses.NewErrorResponse(responses.ErrorResponseModel{
			Message: constants.ErrMsgUnableToParseJSONBody,
		})

		rw.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(rw).Encode(errors)
		
		return
    } 

	user, err := uc.us.Update(uint(id), req)

	if err != nil {
		if err == constants.ErrUserNotFound {
			rw.WriteHeader(http.StatusNotFound)
		} else if err == constants.ErrEmailIsTaken {
			rw.WriteHeader(http.StatusBadRequest)
		}

		resErr := utils.ConvertToErrorResponse(err)
		_ = json.NewEncoder(rw).Encode(resErr)

		return
	}

	err = json.NewEncoder(rw).Encode(user)
	if err != nil {
		errors := responses.NewErrorResponse(responses.ErrorResponseModel{
			Message: ErrorUnableToMarshalJson.Error(),
		})

		rw.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(rw).Encode(errors)
	}
}

// swagger:route GET /api/users/{id}/questions users listUserQuestions
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

		rw.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(rw).Encode(errors)

		return
    } 

	vars := mux.Vars(r)

	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		errors := responses.NewErrorResponse(responses.ErrorResponseModel{
			Message: constants.ErrMsgUnableToConvertId,
		})

		rw.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(rw).Encode(errors)
		
		return
	}

	questions, err := uc.us.GetQuestions(uint(id), request)
	if err != nil {
		resErr := utils.ConvertToErrorResponse(err)
		
		rw.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(rw).Encode(resErr)
		
		return
	}

	err = json.NewEncoder(rw).Encode(questions)
	if err != nil {
		errors := responses.NewErrorResponse(responses.ErrorResponseModel{
			Message: ErrorUnableToMarshalJson.Error(),
		})

		rw.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(rw).Encode(errors)
	}
}
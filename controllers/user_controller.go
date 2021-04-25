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

// swagger:route GET /api/users users users
// Returns a list of users
//
// parameters:
// + name: pageNumber
//	 in: query
//	 schema: int
// + name: pageSize
//	 in: query
//	 schema: int
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

// swagger:route GET /api/users/{id} users user
// Returns a single user
//
// responses:
//	200: UserResponse
func (uc *UserController) GetById(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		errors := responses.NewErrorResponse(responses.ErrorResponseModel{
			Message: "Unable to convert id",
		})

		out, _ := json.Marshal(errors)

		http.Error(rw, string(out), http.StatusNotFound)
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
	sub, err := utils.ExtractSubFromJwt(r)

	if err != nil {
		http.Error(rw, "", http.StatusBadRequest)
		return;
	}

	id, err := strconv.ParseUint(sub, 10, 64)
	if err != nil {
		errors := responses.NewErrorResponse(responses.ErrorResponseModel{
			Message: "Unable to convert id",
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
			Message: constants.UnableToParseJSONBody,
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

// swagger:route PUT /api/users users user
//
// responses:
//	200: UserResponse
func (uc *UserController) Update(rw http.ResponseWriter, r *http.Request) {
	sub, err := utils.ExtractSubFromJwt(r)

	if err != nil {
		http.Error(rw, "", http.StatusBadRequest)
		return;
	}

	id, err := strconv.ParseUint(sub, 10, 64)
	if err != nil {
		errors := responses.NewErrorResponse(responses.ErrorResponseModel{
			Message: "Unable to convert id",
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
			Message: constants.UnableToParseJSONBody,
		})

		out, _ := json.Marshal(errors)

		http.Error(rw, string(out), http.StatusBadRequest)
		return
    } 

	user, errRes := uc.us.Update(uint(id), req)

	if errRes != nil {
		out, _ := json.Marshal(errRes)

		http.Error(rw, string(out), http.StatusBadRequest)
		return;
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
// responses:
//	200: QuestionsResponse
func (uc *UserController) GetQuestions(rw http.ResponseWriter, r *http.Request) {
	var request requests.QuestionSearchRequest

	err := decoder.Decode(&request, r.URL.Query())
    if err != nil {
		http.Error(rw, "Unable to parse query parametars.", http.StatusBadRequest)
		return
    } 

	vars := mux.Vars(r)

	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		errors := responses.NewErrorResponse(responses.ErrorResponseModel{
			Message: "Unable to convert id",
		})

		out, _ := json.Marshal(errors)

		http.Error(rw, string(out), http.StatusNotFound)
		return
	}

	questions, errRes := uc.us.GetQuestions(uint(id), request)
	if errRes != nil {
		out, _ := json.Marshal(errRes)

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
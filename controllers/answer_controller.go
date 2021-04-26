package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/EnisMulic/Ask.it.Backend/constants"
	"github.com/EnisMulic/Ask.it.Backend/contracts/requests"
	"github.com/EnisMulic/Ask.it.Backend/contracts/responses"
	"github.com/EnisMulic/Ask.it.Backend/services"
	"github.com/EnisMulic/Ask.it.Backend/utils"
	"github.com/gorilla/mux"
)


type AnswerController struct {
	as *services.AnswerService
}

func NewAnswerController(as *services.AnswerService) *AnswerController {
	return &AnswerController{as}
}

// swagger:route PUT /api/answers/{id} answers answer
//
// responses:
//	200: AnswerResponse
func (ac *AnswerController) Update (rw http.ResponseWriter, r *http.Request) {
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

	sub, err := utils.ExtractSubFromJwt(r)

	if err != nil {
		http.Error(rw, "", http.StatusBadRequest)
		return;
	}

	userId, err := strconv.ParseUint(sub, 10, 64)
	if err != nil {
		errors := responses.NewErrorResponse(responses.ErrorResponseModel{
			Message: "Unable to convert id",
		})

		out, _ := json.Marshal(errors)

		http.Error(rw, string(out), http.StatusBadRequest)
		return
	}

	var req requests.AnswerUpdateRequest

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

	user, errRes := ac.as.Update(uint(id), uint(userId), req)

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

// swagger:route DELETE /api/answers/{id} answers bool
//
func (ac *AnswerController) Delete (rw http.ResponseWriter, r *http.Request) {
	
}

// swagger:route POST /api/answers/{id}/like answers answer
//
func (ac *AnswerController) Like (rw http.ResponseWriter, r *http.Request) {

}


// swagger:route POST /api/answers/{id}/like/undo answers answer
//
func (ac *AnswerController) LikeUndo (rw http.ResponseWriter, r *http.Request) {

}

// swagger:route POST /api/answers/{id}/dislike answers answer
//
func (ac *AnswerController) Dislike (rw http.ResponseWriter, r *http.Request) {
	
}

// swagger:route POST /api/answers/{id}/dislike/undo answers answer
//
func (ac *AnswerController) DislikeUndo (rw http.ResponseWriter, r *http.Request) {
	
}
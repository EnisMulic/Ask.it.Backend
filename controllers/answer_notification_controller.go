package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/EnisMulic/Ask.it.Backend/constants"
	"github.com/EnisMulic/Ask.it.Backend/contracts/responses"
	"github.com/EnisMulic/Ask.it.Backend/services"
	"github.com/EnisMulic/Ask.it.Backend/utils"
	"github.com/gorilla/mux"
)

type AnswerNotificationController struct {
	ans *services.AnswerNotificationService
}

func NewAnswerNotificationController(ans *services.AnswerNotificationService) *AnswerNotificationController {
	return &AnswerNotificationController{ans}
}

// swagger:route POST /api/answer-notifications/{id} answer-notifications markRead
//
// responses:
//	200: 
//  400: ErrorResponse
//  403: ErrorResponse
func (anc *AnswerNotificationController) MarkRead(rw http.ResponseWriter, r *http.Request) {
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

	sub, err := utils.ExtractSubFromJwt(r)

	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	userId, err := strconv.ParseUint(sub, 10, 64)
	if err != nil {
		errors := responses.NewErrorResponse(responses.ErrorResponseModel{
			Message: constants.ErrMsgUnableToConvertId,
		})

		rw.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(rw).Encode(errors)

		return
	}

	err = anc.ans.MarkRead(uint(id), uint(userId))
	if err != nil {
		if err == constants.ErrForbidden {
			rw.WriteHeader(http.StatusForbidden)
		} else {
			rw.WriteHeader(http.StatusBadRequest)
		}
		
		errRes := utils.ConvertToErrorResponse(err)
		_ = json.NewEncoder(rw).Encode(errRes)

		return
	}
}
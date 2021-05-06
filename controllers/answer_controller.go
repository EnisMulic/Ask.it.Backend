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

// swagger:route PUT /api/answers/{id} answers answerUpdate
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
//	200: AnswerResponse
//	400: ErrorResponse
//	403: ErrorResponse
//	404: ErrorResponse
//  500: ErrorResponse
func (ac *AnswerController) Update (rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		errors := responses.NewErrorResponse(responses.ErrorResponseModel{
			Message: constants.ErrMsgUnableToConvertId,
		})

		_ = json.NewEncoder(rw).Encode(errors)
		rw.WriteHeader(http.StatusBadRequest)

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
			Message: constants.ErrMsgUnableToConvertUserId,
		})

		_ = json.NewEncoder(rw).Encode(errors)
		rw.WriteHeader(http.StatusBadRequest)

		return
	}

	var req requests.AnswerUpdateRequest

	decoder := json.NewDecoder(r.Body)
	
	err = decoder.Decode(&req)
    if err != nil {
		errors := responses.NewErrorResponse(responses.ErrorResponseModel{
			Message: constants.ErrMsgUnableToParseJSONBody,
		})

		_ = json.NewEncoder(rw).Encode(errors)
		rw.WriteHeader(http.StatusBadRequest)

		return
    } 

	user, err := ac.as.Update(uint(id), uint(userId), req)
	
	if err != nil {
		errRes := utils.ConvertToErrorResponse(err)

		_ = json.NewEncoder(rw).Encode(errRes)
		rw.WriteHeader(http.StatusBadRequest)

		if err == constants.ErrAnswerNotFound {
			rw.WriteHeader(http.StatusNotFound)
		} else if err == constants.ErrForbidden {
			rw.WriteHeader(http.StatusForbidden)
		} else {
			rw.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	err = json.NewEncoder(rw).Encode(user)
	if err != nil {
		errors := responses.NewErrorResponse(responses.ErrorResponseModel{
			Message: constants.ErrMsgUnableToMarshalJson,
		})

		_ = json.NewEncoder(rw).Encode(errors)
		rw.WriteHeader(http.StatusInternalServerError)
	}
}

// swagger:route DELETE /api/answers/{id} answers bool
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
//	204: 
//	400: ErrorResponse
//	403: ErrorResponse
//	404: ErrorResponse
//  500: ErrorResponse
func (ac *AnswerController) Delete (rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		errors := responses.NewErrorResponse(responses.ErrorResponseModel{
			Message: constants.ErrMsgUnableToConvertId,
		})

		_ = json.NewEncoder(rw).Encode(errors)
		rw.WriteHeader(http.StatusBadRequest)

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
			Message: constants.ErrMsgUnableToConvertUserId,
		})

		_ = json.NewEncoder(rw).Encode(errors)
		rw.WriteHeader(http.StatusBadRequest)

		return
	}

	err = ac.as.Delete(uint(id), uint(userId))

	if err != nil {
		errRes := utils.ConvertToErrorResponse(err)

		_ = json.NewEncoder(rw).Encode(errRes)

		if err == constants.ErrAnswerNotFound {
			rw.WriteHeader(http.StatusNotFound)
		} else if err == constants.ErrForbidden {
			rw.WriteHeader(http.StatusForbidden)
		} else {
			rw.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}

// swagger:route POST /api/answers/{id}/like answers likeAnswer
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
//	200: 
//	400: ErrorResponse
//	404: ErrorResponse
//  500: ErrorResponse
func (ac *AnswerController) Like (rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		errors := responses.NewErrorResponse(responses.ErrorResponseModel{
			Message: constants.ErrMsgUnableToConvertId,
		})

		_ = json.NewEncoder(rw).Encode(errors)
		rw.WriteHeader(http.StatusNotFound)

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

		_ = json.NewEncoder(rw).Encode(errors)
		rw.WriteHeader(http.StatusBadRequest)

		return
	}

	err = ac.as.Like(uint(id), uint(userId))
	if err != nil {
		errRes := utils.ConvertToErrorResponse(err)

		_ = json.NewEncoder(rw).Encode(errRes)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
}


// swagger:route POST /api/answers/{id}/like/undo answers likeAnswerUndo
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
//	200: 
//	400: ErrorResponse
//	404: ErrorResponse
//  500: ErrorResponse
func (ac *AnswerController) LikeUndo (rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		errors := responses.NewErrorResponse(responses.ErrorResponseModel{
			Message: constants.ErrMsgUnableToConvertId,
		})

		_ = json.NewEncoder(rw).Encode(errors)
		rw.WriteHeader(http.StatusBadRequest)

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
			Message: constants.ErrMsgUnableToConvertUserId,
		})

		_ = json.NewEncoder(rw).Encode(errors)
		rw.WriteHeader(http.StatusBadRequest)

		return
	}

	err = ac.as.LikeUndo(uint(id), uint(userId))
	if err != nil {
		errRes := utils.ConvertToErrorResponse(err)
		
		_ = json.NewEncoder(rw).Encode(errRes)
		rw.WriteHeader(http.StatusInternalServerError)

		return
	}
}

// swagger:route POST /api/answers/{id}/dislike answers dislikeAnswer
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
//	200: 
//	400: ErrorResponse
//	404: ErrorResponse
//  500: ErrorResponse
func (ac *AnswerController) Dislike (rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		errors := responses.NewErrorResponse(responses.ErrorResponseModel{
			Message: constants.ErrMsgUnableToConvertId,
		})

		_ = json.NewEncoder(rw).Encode(errors)
		rw.WriteHeader(http.StatusBadRequest)

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
			Message: constants.ErrMsgUnableToConvertUserId,
		})

		_ = json.NewEncoder(rw).Encode(errors)
		rw.WriteHeader(http.StatusBadRequest)

		return
	}

	err = ac.as.Dislike(uint(id), uint(userId))
	if err != nil {
		errRes := utils.ConvertToErrorResponse(err)
		
		_ = json.NewEncoder(rw).Encode(errRes)
		rw.WriteHeader(http.StatusInternalServerError)

		return
	}
}

// swagger:route POST /api/answers/{id}/dislike/undo answers dislikeAnswerUndo
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
//	200: 
//	400: ErrorResponse
//	404: ErrorResponse
//  500: ErrorResponse
func (ac *AnswerController) DislikeUndo (rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		errors := responses.NewErrorResponse(responses.ErrorResponseModel{
			Message: constants.ErrMsgUnableToConvertId,
		})

		_ = json.NewEncoder(rw).Encode(errors)
		rw.WriteHeader(http.StatusBadRequest)
		
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
			Message: constants.ErrMsgUnableToConvertUserId,
		})

		_ = json.NewEncoder(rw).Encode(errors)
		rw.WriteHeader(http.StatusBadRequest)	

		return
	}

	err = ac.as.DislikeUndo(uint(id), uint(userId))
	if err != nil {
		errRes := utils.ConvertToErrorResponse(err)
		
		_ = json.NewEncoder(rw).Encode(errRes)
		rw.WriteHeader(http.StatusInternalServerError)

		return
	}
}
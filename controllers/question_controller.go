package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/EnisMulic/Ask.it.Backend/constants"
	"github.com/EnisMulic/Ask.it.Backend/contracts/requests"
	"github.com/EnisMulic/Ask.it.Backend/contracts/responses"
	"github.com/EnisMulic/Ask.it.Backend/services"
	"github.com/EnisMulic/Ask.it.Backend/utils"
	"github.com/gorilla/mux"
)


type QuestionController struct {
	l *log.Logger
	qs *services.QuestionService
}

func NewQuestionController(l *log.Logger, qs *services.QuestionService) *QuestionController {
	return &QuestionController{l, qs}
}

// swagger:route GET /api/questions questions questions 
//
// responses:
//	200: QuestionsResponse
//	400: ErrorResponse
//  500: ErrorResponse
func (qc *QuestionController) Get(rw http.ResponseWriter, r *http.Request) {
	var request requests.QuestionSearchRequest

	err := decoder.Decode(&request, r.URL.Query())
    if err != nil {
        log.Println("Error in GET parameters : ", err)

		errors := responses.NewErrorResponse(responses.ErrorResponseModel{
			Message: constants.ErrMsgUnableToParseQueryParametars,
		})

		out, _ := json.Marshal(errors)

		http.Error(rw, string(out), http.StatusBadRequest)
		return
    } 

	users := qc.qs.Get(request)

	err = json.NewEncoder(rw).Encode(users)
	if err != nil {
		errors := responses.NewErrorResponse(responses.ErrorResponseModel{
			Message: constants.ErrMsgUnableToMarshalJson,
		})

		out, _ := json.Marshal(errors)

		http.Error(rw, string(out), http.StatusInternalServerError)
		return
	}
}

// swagger:route GET /api/questions-top questions questions 
//
// parameters: PaginationQuery
//
// responses:
//	200: QuestionsResponse
//	400: ErrorResponse
//  500: ErrorResponse
func (qc *QuestionController) GetHot(rw http.ResponseWriter, r *http.Request) {
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

	users := qc.qs.GetHot(request)

	err = json.NewEncoder(rw).Encode(users)
	if err != nil {
		errors := responses.NewErrorResponse(responses.ErrorResponseModel{
			Message: constants.ErrMsgUnableToMarshalJson,
		})

		out, _ := json.Marshal(errors)

		http.Error(rw, string(out), http.StatusInternalServerError)
		return
	}
}

// swagger:route GET /api/questions/{id} questions question
//
// parameters:
// + name: id
//	 in: path
//	 schema: int
//
// responses:
//	200: QuestionResponse
//	400: ErrorResponse
//	404: ErrorResponse
//  500: ErrorResponse
func (qc *QuestionController) GetById(rw http.ResponseWriter, r *http.Request) {
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

	user, err := qc.qs.GetById(uint(id))

	if err != nil {
		errRes := utils.ConvertToErrorResponse(err)

		out, _ := json.Marshal(errRes)

		http.Error(rw, string(out), http.StatusNotFound)
		return
	}

	err = json.NewEncoder(rw).Encode(user)
	if err != nil {
		errors := responses.NewErrorResponse(responses.ErrorResponseModel{
			Message: constants.ErrMsgUnableToMarshalJson,
		})

		out, _ := json.Marshal(errors)

		http.Error(rw, string(out), http.StatusInternalServerError)
	}
}

// swagger:route POST /api/questions questions question
//
// responses:
//	204: QuestionResponse
//	400: ErrorResponse
//  500: ErrorResponse
func (qc *QuestionController) Create(rw http.ResponseWriter, r *http.Request) {
	sub, err := utils.ExtractSubFromJwt(r)
	
	if err != nil {
		http.Error(rw, "", http.StatusBadRequest)
		return;
	}

	userId, err := strconv.ParseUint(sub, 10, 64)
	if err != nil {
		errors := responses.NewErrorResponse(responses.ErrorResponseModel{
			Message: constants.ErrMsgUnableToConvertUserId,
		})

		out, _ := json.Marshal(errors)

		http.Error(rw, string(out), http.StatusBadRequest)
		return
	}

	var req requests.QuestionInsertRequest

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

	question, err := qc.qs.Create(uint(userId), req)
	if err != nil {
		errRes := utils.ConvertToErrorResponse(err)

		out, _ := json.Marshal(errRes)


		http.Error(rw, string(out), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(rw).Encode(question)
	if err != nil {
		errors := responses.NewErrorResponse(responses.ErrorResponseModel{
			Message: constants.ErrMsgUnableToMarshalJson,
		})

		out, _ := json.Marshal(errors)

		http.Error(rw, string(out), http.StatusInternalServerError)
	}
}

// swagger:route DELETE /api/questions/{id} questions bool
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
func (qc *QuestionController) Delete(rw http.ResponseWriter, r *http.Request) {
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
	
	sub, err := utils.ExtractSubFromJwt(r)
	
	if err != nil {
		http.Error(rw, "", http.StatusBadRequest)
		return;
	}

	userId, err := strconv.ParseUint(sub, 10, 64)
	if err != nil {
		errors := responses.NewErrorResponse(responses.ErrorResponseModel{
			Message: constants.ErrMsgUnableToConvertUserId,
		})

		out, _ := json.Marshal(errors)

		http.Error(rw, string(out), http.StatusBadRequest)
		return
	}

	err = qc.qs.Delete(uint(id), uint(userId))
	if err != nil {
		errRes := utils.ConvertToErrorResponse(err)
		out, _ := json.Marshal(errRes)

		if err == constants.ErrQuestionNotFound {
			http.Error(rw, string(out), http.StatusNotFound)
		} else if err == constants.ErrForbidden {
			http.Error(rw, string(out), http.StatusForbidden)
		}
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}

// swagger:route POST /api/questions/{id}/like questions question
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
func (qc *QuestionController) Like (rw http.ResponseWriter, r *http.Request) {
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
	
	sub, err := utils.ExtractSubFromJwt(r)
	
	if err != nil {
		http.Error(rw, "", http.StatusBadRequest)
		return;
	}

	userId, err := strconv.ParseUint(sub, 10, 64)
	if err != nil {
		errors := responses.NewErrorResponse(responses.ErrorResponseModel{
			Message: constants.ErrMsgUnableToConvertUserId,
		})

		out, _ := json.Marshal(errors)

		http.Error(rw, string(out), http.StatusBadRequest)
		return
	}

	err = qc.qs.Like(uint(id), uint(userId))
	if err != nil {
		resErr := utils.ConvertToErrorResponse(err)
		out, _ := json.Marshal(resErr)

		http.Error(rw, string(out), http.StatusInternalServerError)
		return
	}
}

// swagger:route POST /api/questions/{id}/like/undo questions question
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
func (qc *QuestionController) LikeUndo (rw http.ResponseWriter, r *http.Request) {
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
	
	sub, err := utils.ExtractSubFromJwt(r)
	
	if err != nil {
		http.Error(rw, "", http.StatusBadRequest)
		return;
	}

	userId, err := strconv.ParseUint(sub, 10, 64)
	if err != nil {
		errors := responses.NewErrorResponse(responses.ErrorResponseModel{
			Message: constants.ErrMsgUnableToConvertUserId,
		})

		out, _ := json.Marshal(errors)

		http.Error(rw, string(out), http.StatusBadRequest)
		return
	}

	err = qc.qs.LikeUndo(uint(id), uint(userId))
	if err != nil {
		resErr := utils.ConvertToErrorResponse(err)
		out, _ := json.Marshal(resErr)

		http.Error(rw, string(out), http.StatusInternalServerError)
		return
	}
}

// swagger:route POST /api/questions/{id}/dislike questions question
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
func (qc *QuestionController) Dislike (rw http.ResponseWriter, r *http.Request) {
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
	
	sub, err := utils.ExtractSubFromJwt(r)
	
	if err != nil {
		http.Error(rw, "", http.StatusBadRequest)
		return;
	}

	userId, err := strconv.ParseUint(sub, 10, 64)
	if err != nil {
		errors := responses.NewErrorResponse(responses.ErrorResponseModel{
			Message: constants.ErrMsgUnableToConvertId,
		})

		out, _ := json.Marshal(errors)

		http.Error(rw, string(out), http.StatusBadRequest)
		return
	}

	err = qc.qs.Dislike(uint(id), uint(userId))
	if err != nil {
		resErr := utils.ConvertToErrorResponse(err)
		out, _ := json.Marshal(resErr)

		http.Error(rw, string(out), http.StatusInternalServerError)
		return
	}
}

// swagger:route POST /api/questions/{id}/dislike/undo questions question
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
func (qc *QuestionController) DislikeUndo (rw http.ResponseWriter, r *http.Request) {
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
	
	sub, err := utils.ExtractSubFromJwt(r)
	
	if err != nil {
		http.Error(rw, "", http.StatusBadRequest)
		return;
	}

	userId, err := strconv.ParseUint(sub, 10, 64)
	if err != nil {
		errors := responses.NewErrorResponse(responses.ErrorResponseModel{
			Message: constants.ErrMsgUnableToConvertId,
		})

		out, _ := json.Marshal(errors)

		http.Error(rw, string(out), http.StatusBadRequest)
		return
	}

	err = qc.qs.DislikeUndo(uint(id), uint(userId))
	if err != nil {
		resErr := utils.ConvertToErrorResponse(err)
		out, _ := json.Marshal(resErr)

		http.Error(rw, string(out), http.StatusInternalServerError)
		return
	}
}

// swagger:route POST /api/questions/{id}/answers questions answer
//
// parameters:
// + name: id
//	 in: path
//	 schema: int
//
// responses:
//	204: AnswerResponse
//	400: ErrorResponse
//  500: ErrorResponse
func (qc *QuestionController) CreateAnswer (rw http.ResponseWriter, r *http.Request) {
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

	sub, err := utils.ExtractSubFromJwt(r)
	
	if err != nil {
		http.Error(rw, "", http.StatusBadRequest)
		return;
	}

	userId, err := strconv.ParseUint(sub, 10, 64)
	if err != nil {
		errors := responses.NewErrorResponse(responses.ErrorResponseModel{
			Message: constants.ErrMsgUnableToConvertId,
		})

		out, _ := json.Marshal(errors)

		http.Error(rw, string(out), http.StatusBadRequest)
		return
	}

	var req requests.AnswerInsertRequest

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

	question, err := qc.qs.CreateAnswer(uint(id), uint(userId), req)
	if err != nil {
		resErr := utils.ConvertToErrorResponse(err)
		out, _ := json.Marshal(resErr)

		http.Error(rw, string(out), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(rw).Encode(question)
	if err != nil {
		errors := responses.NewErrorResponse(responses.ErrorResponseModel{
			Message: constants.ErrMsgUnableToMarshalJson,
		})

		out, _ := json.Marshal(errors)

		http.Error(rw, string(out), http.StatusInternalServerError)
		return
	}
	//rw.WriteHeader(http.StatusCreated)
}
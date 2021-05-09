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

// swagger:route GET /api/questions questions questionsSearch 
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

		rw.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(rw).Encode(errors)

		return
    } 

	users := qc.qs.Get(request)

	err = json.NewEncoder(rw).Encode(users)
	if err != nil {
		errors := responses.NewErrorResponse(responses.ErrorResponseModel{
			Message: constants.ErrMsgUnableToMarshalJson,
		})

		rw.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(rw).Encode(errors)

		return
	}
}

// swagger:route GET /api/questions-top questions hotQuestionsSearch
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

		rw.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(rw).Encode(errors)

		return
    } 

	users := qc.qs.GetHot(request)

	err = json.NewEncoder(rw).Encode(users)
	if err != nil {
		errors := responses.NewErrorResponse(responses.ErrorResponseModel{
			Message: constants.ErrMsgUnableToMarshalJson,
		})

		rw.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(rw).Encode(errors)
		
		return
	}
}

// swagger:route GET /api/questions/{id} questions getQuestion
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

		rw.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(rw).Encode(errors)

		return
	}

	user, err := qc.qs.GetById(uint(id))

	if err != nil {
		errRes := utils.ConvertToErrorResponse(err)

		rw.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(rw).Encode(errRes)

		return
	}

	err = json.NewEncoder(rw).Encode(user)
	if err != nil {
		errors := responses.NewErrorResponse(responses.ErrorResponseModel{
			Message: constants.ErrMsgUnableToMarshalJson,
		})

		rw.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(rw).Encode(errors)
	}
}

// swagger:route POST /api/questions questions questionInsert
//
// security:
//  - Bearer: []
//
// responses:
//	204: QuestionResponse
//	400: ErrorResponse
//  500: ErrorResponse
func (qc *QuestionController) Create(rw http.ResponseWriter, r *http.Request) {
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

		rw.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(rw).Encode(errors)

		return
	}

	var req requests.QuestionInsertRequest

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

	question, err := qc.qs.Create(uint(userId), req)
	if err != nil {
		errRes := utils.ConvertToErrorResponse(err)

		rw.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(rw).Encode(errRes)

		return
	}

	err = json.NewEncoder(rw).Encode(question)
	if err != nil {
		errors := responses.NewErrorResponse(responses.ErrorResponseModel{
			Message: constants.ErrMsgUnableToMarshalJson,
		})

		rw.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(rw).Encode(errors)
	}
}

// swagger:route DELETE /api/questions/{id} questions deleteQuestion
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
func (qc *QuestionController) Delete(rw http.ResponseWriter, r *http.Request) {
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
			Message: constants.ErrMsgUnableToConvertUserId,
		})

		rw.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(rw).Encode(errors)

		return
	}

	err = qc.qs.Delete(uint(id), uint(userId))
	if err != nil {
		if err == constants.ErrQuestionNotFound {
			rw.WriteHeader(http.StatusNotFound)
		} else if err == constants.ErrForbidden {
			rw.WriteHeader(http.StatusForbidden)
		}

		errRes := utils.ConvertToErrorResponse(err)
		_ = json.NewEncoder(rw).Encode(errRes)

		return
	}

	rw.WriteHeader(http.StatusNoContent)
}

// swagger:route POST /api/questions/{id}/like questions likeQuestion
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
func (qc *QuestionController) Like (rw http.ResponseWriter, r *http.Request) {
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
			Message: constants.ErrMsgUnableToConvertUserId,
		})

		rw.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(rw).Encode(errors)

		return
	}

	err = qc.qs.Like(uint(id), uint(userId))
	if err != nil {
		resErr := utils.ConvertToErrorResponse(err)

		rw.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(rw).Encode(resErr)

		return
	}
}

// swagger:route POST /api/questions/{id}/like/undo questions likeQuestionUndo
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
func (qc *QuestionController) LikeUndo (rw http.ResponseWriter, r *http.Request) {
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
			Message: constants.ErrMsgUnableToConvertUserId,
		})

		rw.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(rw).Encode(errors)

		return
	}

	err = qc.qs.LikeUndo(uint(id), uint(userId))
	if err != nil {
		resErr := utils.ConvertToErrorResponse(err)
		
		rw.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(rw).Encode(resErr)

		return
	}
}

// swagger:route POST /api/questions/{id}/dislike questions dislikeQuestion
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
func (qc *QuestionController) Dislike (rw http.ResponseWriter, r *http.Request) {
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
		return;
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

	err = qc.qs.Dislike(uint(id), uint(userId))
	if err != nil {
		resErr := utils.ConvertToErrorResponse(err)
		
		rw.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(rw).Encode(resErr)

		return
	}
}

// swagger:route POST /api/questions/{id}/dislike/undo questions undoQuestionDislike
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
func (qc *QuestionController) DislikeUndo (rw http.ResponseWriter, r *http.Request) {
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
		return;
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

	err = qc.qs.DislikeUndo(uint(id), uint(userId))
	if err != nil {
		resErr := utils.ConvertToErrorResponse(err)

		rw.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(rw).Encode(resErr)

		return
	}
}

// swagger:route POST /api/questions/{id}/answers questions answerInsert
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

		rw.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(rw).Encode(errors)

		return
	}

	sub, err := utils.ExtractSubFromJwt(r)
	
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return;
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

	var req requests.AnswerInsertRequest

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

	question, err := qc.qs.CreateAnswer(uint(id), uint(userId), req)
	if err != nil {
		resErr := utils.ConvertToErrorResponse(err)

		rw.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(rw).Encode(resErr)

		return
	}

	rw.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(rw).Encode(question)

	if err != nil {
		errors := responses.NewErrorResponse(responses.ErrorResponseModel{
			Message: constants.ErrMsgUnableToMarshalJson,
		})

		rw.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(rw).Encode(errors)

		return
	}
}
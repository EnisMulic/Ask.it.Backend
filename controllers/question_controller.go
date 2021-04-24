package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/EnisMulic/Ask.it.Backend/contracts/requests"
	"github.com/EnisMulic/Ask.it.Backend/services"
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
func (qc *QuestionController) Get(rw http.ResponseWriter, r *http.Request) {
	var request requests.QuestionSearchRequest

	err := decoder.Decode(&request, r.URL.Query())
    if err != nil {
        log.Println("Error in GET parameters : ", err)
		http.Error(rw, "Unable to parse query parametars.", http.StatusBadRequest)
		return
    } 

	users := qc.qs.Get(request)

	err = json.NewEncoder(rw).Encode(users)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

// swagger:route GET /api/questions/{id} questions question
//
// responses:
//	200: QuestionResponse
func (qc *QuestionController) GetById(rw http.ResponseWriter, r *http.Request) {

}

// swagger:route POST /api/questions questions question
//
// responses:
//	204: QuestionResponse
func (qc *QuestionController) Create(rw http.ResponseWriter, r *http.Request) {

}

// swagger:route DELETE /api/questions/{id} questions bool
//
func (qc *QuestionController) Delete(rw http.ResponseWriter, r *http.Request) {

}

// swagger:route POST /api/questions/{id}/like questions question
//
func (qc *QuestionController) Like (rw http.ResponseWriter, r *http.Request) {

}

// swagger:route POST /api/questions/{id}/like/undo questions question
//
func (qc *QuestionController) LikeUndo (rw http.ResponseWriter, r *http.Request) {

}

// swagger:route POST /api/questions/{id}/dislike questions question
//
func (qc *QuestionController) Dislike (rw http.ResponseWriter, r *http.Request) {
	
}

// swagger:route POST /api/questions/{id}/dislike/undo questions question
//
func (qc *QuestionController) DislikeUndo (rw http.ResponseWriter, r *http.Request) {
	
}

// swagger:route GET /api/questions/{id}/answers questions answers
//
// responses:
//	200: AnswersResponse
func (qc *QuestionController) GetAnswers (rw http.ResponseWriter, r *http.Request) {

}

// swagger:route POST /api/questions/{id}/answers questions answer
//
// responses:
//	204: AnswerResponse
func (qc *QuestionController) CreateAnswer (rw http.ResponseWriter, r *http.Request) {

}

// swagger:route PUT /api/questions/{id}/answers/{id} questions answer
//
// responses:
//	200: AnswerResponse
func (qc *QuestionController) EditAnswer (rw http.ResponseWriter, r *http.Request) {

}

// swagger:route DELETE /api/questions/{id}/answers/{id} questions bool
//
func (qc * QuestionController) DeleteAnswer (rw http.ResponseWriter, r *http.Request) {

}

// swagger:route POST /api/questions/{id}/answers/{id}/like questions answer
//
func (qc *QuestionController) LikeAnswer (rw http.ResponseWriter, r *http.Request) {

}


// swagger:route POST /api/questions/{id}/answers/{id}/like/undo questions answer
//
func (qc *QuestionController) LikeAnswerUndo (rw http.ResponseWriter, r *http.Request) {

}

// swagger:route POST /api/questions/{id}/answers/{id}/dislike questions answer
//
func (qc *QuestionController) DislikeAnswer (rw http.ResponseWriter, r *http.Request) {
	
}

// swagger:route POST /api/questions/{id}/answers/{id}/dislike/undo questions answer
//
func (qc *QuestionController) DislikeAnswerUndo (rw http.ResponseWriter, r *http.Request) {
	
}
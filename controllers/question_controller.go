package controllers

import (
	"log"
	"net/http"
)


type QuestionController struct {
	l *log.Logger
}

func NewQuestionController(l *log.Logger) *QuestionController {
	return &QuestionController{l}
}

// swagger:route GET /api/questions questions questions 
//
func (qc *QuestionController) Get(rw http.ResponseWriter, r *http.Request) {

}

// swagger:route GET /api/questions/{id} questions question
//
func (qc *QuestionController) GetById(rw http.ResponseWriter, r *http.Request) {

}

// swagger:route POST /api/questions questions question
//
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
func (qc *QuestionController) GetAnswers (rw http.ResponseWriter, r *http.Request) {

}

// swagger:route POST /api/questions/{id}/answers questions answer
//
func (qc *QuestionController) CreateAnswer (rw http.ResponseWriter, r *http.Request) {

}

// swagger:route PUT /api/questions/{id}/answers/{id} questions answer
//
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
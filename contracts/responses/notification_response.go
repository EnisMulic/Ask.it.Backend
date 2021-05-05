package responses

// swagger:model AnswerNotification
type AnswerNotification struct {
	ID         uint   `json:"id"`
	QuestionID uint   `json:"questionId"`
	AnswerID   uint   `json:"answerId"`
	UserID     uint   `json:"userId"`
	Content    string `json:"content"`
	IsRead     bool   `json:"isRead"`
}
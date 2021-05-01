package responses

type AnswerNotification struct {
	ID         uint
	QuestionID uint
	AnswerID   uint
	UserID     uint
	Content    string
	IsRead     bool
}
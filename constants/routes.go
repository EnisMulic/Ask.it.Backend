package constants

const (
	ApiRoot = "/api"

	AuthBaseRoute = ApiRoot + "/auth"
	LoginRoute    = AuthBaseRoute + "/login"
	RegisterRoute = AuthBaseRoute + "/register"

	UsersBaseRoute          = ApiRoot + "/users"
	GetUsersRoute           = UsersBaseRoute
	GetUserByIdRoute        = UsersBaseRoute + "/{id}"
	GetMeRoute              = UsersBaseRoute + "/me"
	ChangeUserPasswordRoute = UsersBaseRoute + "/change-password"
	UpdateUserRoute         = UsersBaseRoute
	GetUsersQuestionsRoute  = UsersBaseRoute + "/{id}/questions"

	QuestionsBaseRoute        = ApiRoot + "/questions"
	GetQuestionsRoute         = QuestionsBaseRoute
	GetQuestionByIdRoute      = QuestionsBaseRoute + "/{id}"
	CreateQuestionRoute       = QuestionsBaseRoute
	DeleteQuestionRoute       = QuestionsBaseRoute + "/{id}"
	LikeQuestionRoute         = QuestionsBaseRoute + "/{id}/like"
	LikeQuestionUndoRoute     = QuestionsBaseRoute + "/{id}/like/undo"
	DislikeQuestionRoute      = QuestionsBaseRoute + "/{id}/dislike"
	DislikeQuestionUndoRoute  = QuestionsBaseRoute + "/{id}/dislike/undo"
	GetQuestionAnswersRoute   = QuestionsBaseRoute + "/{id}/answers"
	CreateQuestionAnswerRoute = QuestionsBaseRoute + "/{id}/answers"

	AnswerBaseRoute        = ApiRoot + "/answers"
	UpdateAnswerRoute      = AnswerBaseRoute + "/{id}"
	DeleteAnswerRoute      = AnswerBaseRoute + "/{id}"
	LikeAnswerRoute        = AnswerBaseRoute + "/{id}/like"
	LikeAnswerUndoRoute    = AnswerBaseRoute + "/{id}/like/undo"
	DislikeAnswerRoute     = AnswerBaseRoute + "/{id}/dislike"
	DislikeAnswerUndoRoute = AnswerBaseRoute + "/{id}/dislike/undo"
)
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

	QuestionsBaseRoute             = ApiRoot + "/questions"
	GetQuestionsRoute              = QuestionsBaseRoute
	GetQuestionByIdRoute           = QuestionsBaseRoute + "/{id}"
	CreateQuestionRoute            = QuestionsBaseRoute
	DeleteQuestionRoute            = QuestionsBaseRoute + "/{id}"
	LikeQuestionRoute              = QuestionsBaseRoute + "/{id}/like"
	LikeQuestionUndoRoute          = QuestionsBaseRoute + "/{id}/like/undo"
	DislikeQuestionRoute           = QuestionsBaseRoute + "/{id}/dislike"
	DislikeQuestionUndoRoute       = QuestionsBaseRoute + "/{id}/dislike/undo"
	GetQuestionAnswersRoute        = QuestionsBaseRoute + "/{id}/answers"
	CreateQuestionAnswerRoute      = QuestionsBaseRoute + "/{id}/answers"
	EditQuestionAnswerRoute        = QuestionsBaseRoute + "/{id}/answers/{answer_id}"
	DeleteQuestionAnswerRoute      = QuestionsBaseRoute + "/{id}/answers/{answer_id}"
	LikeQuestionAnswerRoute        = QuestionsBaseRoute + "/{id}/answers/{answer_id}/like"
	LikeQuestionAnswerUndoRoute    = QuestionsBaseRoute + "/{id}/answers/{answer_id}/like/undo"
	DislikeQuestionAnswerRoute     = QuestionsBaseRoute + "/{id}/answers/{answer_id}/dislike"
	DislikeQuestionAnswerUndoRoute = QuestionsBaseRoute + "/{id}/answers/{answer_id}/dislike/undo"
)
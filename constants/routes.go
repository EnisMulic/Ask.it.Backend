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
	GetQuestionAnswersRoute        = QuestionsBaseRoute + "/{id}/answer"
	CreateQuestionAnswerRoute      = QuestionsBaseRoute + "/{id}/answer"
	EditQuestionAnswerRoute        = QuestionsBaseRoute + "/{id}/answer/{answer_id}"
	DeleteQuestionAnswerRoute      = QuestionsBaseRoute + "/{id}/answer/{answer_id}"
	LikeQuestionAnswerRoute        = QuestionsBaseRoute + "/{id}/answer/{answer_id}/like"
	LikeQuestionAnswerUndoRoute    = QuestionsBaseRoute + "/{id}/answer/{answer_id}/like/undo"
	DislikeQuestionAnswerRoute     = QuestionsBaseRoute + "/{id}/answer/{answer_id}/dislike"
	DislikeQuestionAnswerUndoRoute = QuestionsBaseRoute + "/{id}/answer/{answer_id}/dislike/undo"
)
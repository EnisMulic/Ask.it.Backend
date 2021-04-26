package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/EnisMulic/Ask.it.Backend/constants"
	"github.com/EnisMulic/Ask.it.Backend/controllers"
	"github.com/EnisMulic/Ask.it.Backend/database"
	"github.com/EnisMulic/Ask.it.Backend/repositories"
	"github.com/EnisMulic/Ask.it.Backend/services"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	logger := log.New(os.Stdout, "ask.it.api", log.LstdFlags)

	dsn := os.Getenv("CONNECTION_STRING")
	db, dbErr := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	
	if dbErr != nil {
		log.Fatalf(dbErr.Error())
	}
	
	database.Migrate(db)

	userRepo := repositories.NewUserRepository(db)
	questionRepo := repositories.NewQuestionRepository(db)
	questionRatingRepo := repositories.NewUserQuestionRatingRepository(db)
	answerRepo := repositories.NewAnswerRepository(db)
	
	authSevice := services.NewAuthService(userRepo)
	userService := services.NewUserService(userRepo, questionRepo)
	questionService := services.NewQuestionService(questionRepo, questionRatingRepo, answerRepo)

	ac := controllers.NewAuthController(logger, authSevice)
	uc := controllers.NewUserController(logger, userService)
	qc := controllers.NewQuestionController(logger, questionService)

	r := mux.NewRouter()

	// auth routers
	authPostRouter := r.Methods(http.MethodPost).Subrouter()
	authPostRouter.HandleFunc(constants.LoginRoute, ac.Login)
	authPostRouter.HandleFunc(constants.RegisterRoute, ac.Register)

	// users routers
	usersGetRouter := r.Methods(http.MethodGet).Subrouter()
	usersGetRouter.HandleFunc(constants.GetUsersRoute, uc.Get)
	usersGetRouter.HandleFunc(constants.GetUserByIdRoute, uc.GetById)
	usersGetRouter.HandleFunc(constants.GetMeRoute, uc.GetMe)
	usersGetRouter.HandleFunc(constants.GetUsersQuestionsRoute, uc.GetQuestions)

	usersPostRoutes := r.Methods(http.MethodPost).Subrouter()
	usersPostRoutes.HandleFunc(constants.ChangeUserPasswordRoute, uc.ChangePassword)
	usersPostRoutes.HandleFunc(constants.UpdateUserRoute, uc.Update)

	// questions routers
	questionsGetRouter := r.Methods(http.MethodGet).Subrouter()
	questionsGetRouter.HandleFunc(constants.GetQuestionsRoute, qc.Get)
	questionsGetRouter.HandleFunc(constants.GetQuestionByIdRoute, qc.GetById)
	questionsGetRouter.HandleFunc(constants.GetQuestionAnswersRoute, qc.GetAnswers)
	
	questionsPostRouter := r.Methods(http.MethodPost).Subrouter()

	questionsPostRouter.HandleFunc(constants.CreateQuestionRoute, qc.Create)
	questionsPostRouter.HandleFunc(constants.CreateQuestionAnswerRoute, qc.CreateAnswer)

	questionsPostRouter.HandleFunc(constants.LikeQuestionRoute, qc.Like)
	questionsPostRouter.HandleFunc(constants.LikeQuestionUndoRoute, qc.LikeUndo)

	questionsPostRouter.HandleFunc(constants.DislikeQuestionRoute, qc.Dislike)
	questionsPostRouter.HandleFunc(constants.DislikeQuestionUndoRoute, qc.DislikeUndo)

	questionsPostRouter.HandleFunc(constants.LikeQuestionAnswerRoute, qc.LikeAnswer)
	questionsPostRouter.HandleFunc(constants.LikeQuestionAnswerUndoRoute, qc.LikeAnswerUndo)
	
	questionsPostRouter.HandleFunc(constants.DislikeQuestionAnswerRoute, qc.DislikeAnswer)
	questionsPostRouter.HandleFunc(constants.DislikeQuestionAnswerUndoRoute, qc.DislikeAnswerUndo)

	questionsPutRouter := r.Methods(http.MethodPut).Subrouter()
	questionsPutRouter.HandleFunc(constants.UpdateQuestionAnswerRoute, qc.UpdateAnswer)

	questionsDeleteRouter := r.Methods(http.MethodDelete).Subrouter()
	questionsDeleteRouter.HandleFunc(constants.DeleteQuestionRoute, qc.Delete)
	questionsDeleteRouter.HandleFunc(constants.DeleteQuestionAnswerRoute, qc.DeleteAnswer)

	r.PathPrefix("/swagger/").Handler(http.StripPrefix("/swagger/", http.FileServer(http.Dir("./swaggerui/"))))

	addr := os.Getenv("API_ADDRESS")
	srv := &http.Server {
		Handler: r,
		Addr: addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}

	err = srv.ListenAndServe()

	if err != nil {
		log.Fatal(err.Error())
	}
}
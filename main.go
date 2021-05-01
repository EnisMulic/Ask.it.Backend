package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/EnisMulic/Ask.it.Backend/constants"
	"github.com/EnisMulic/Ask.it.Backend/controllers"
	"github.com/EnisMulic/Ask.it.Backend/database"
	"github.com/EnisMulic/Ask.it.Backend/middleware"
	"github.com/EnisMulic/Ask.it.Backend/repositories"
	"github.com/EnisMulic/Ask.it.Backend/services"
	"github.com/EnisMulic/Ask.it.Backend/websockets"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
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

	pool := websockets.NewPool()
    go pool.Start()

	userRepo := repositories.NewUserRepository(db)
	questionRepo := repositories.NewQuestionRepository(db)
	questionRatingRepo := repositories.NewUserQuestionRatingRepository(db)
	answerRepo := repositories.NewAnswerRepository(db)
	answerRatingRepo := repositories.NewUserAnswerRatingRepository(db)

	authSevice := services.NewAuthService(userRepo)
	userService := services.NewUserService(userRepo, questionRepo)
	questionService := services.NewQuestionService(questionRepo, questionRatingRepo, answerRepo, pool)
	answerService := services.NewAnswerRepository(answerRepo, answerRatingRepo)

	ac := controllers.NewAuthController(logger, authSevice)
	uc := controllers.NewUserController(logger, userService)
	qc := controllers.NewQuestionController(logger, questionService)
	answc := controllers.NewAnswerController(answerService)
	
	r := mux.NewRouter().StrictSlash(true)

	nh := websockets.NewNotificationHandler()

	r.HandleFunc(constants.NotificationRoute, func(w http.ResponseWriter, r *http.Request) {
        nh.ServeWS(pool, w, r)
    })

	userGetRouter := r.Methods(http.MethodGet).Subrouter()
	userGetRouter.HandleFunc(constants.GetMeRoute, uc.GetMe)
	userGetRouter.Use(middleware.IsAuthorized)
	userGetRouter.Use(middleware.AddContentType)

	// auth routers
	authPostRouter := r.Methods(http.MethodPost).Subrouter()
	authPostRouter.Use(middleware.AddContentType)
	authPostRouter.HandleFunc(constants.LoginRoute, ac.Login)
	authPostRouter.HandleFunc(constants.RegisterRoute, ac.Register)

	// users routers
	usersGetRouter := r.Methods(http.MethodGet).Subrouter()
	usersGetRouter.HandleFunc(constants.GetUsersRoute, uc.Get)
	usersGetRouter.HandleFunc(constants.GetUserByIdRoute, uc.GetById)
	usersGetRouter.HandleFunc(constants.GetUsersQuestionsRoute, uc.GetQuestions)
	usersGetRouter.HandleFunc(constants.GetTopUsersRoute, uc.GetTop)
	usersGetRouter.Use(middleware.AddContentType)

	usersPostRoutes := r.Methods(http.MethodPost).Subrouter()
	usersPostRoutes.HandleFunc(constants.ChangeUserPasswordRoute, uc.ChangePassword)
	usersPostRoutes.HandleFunc(constants.UpdateUserRoute, uc.Update)
	usersPostRoutes.Use(middleware.AddContentType)
	usersPostRoutes.Use(middleware.IsAuthorized)

	// questions routers
	questionsGetRouter := r.Methods(http.MethodGet).Subrouter()
	questionsGetRouter.HandleFunc(constants.GetQuestionsRoute, qc.Get)
	questionsGetRouter.HandleFunc(constants.GetQuestionByIdRoute, qc.GetById)
	questionsGetRouter.HandleFunc(constants.GetHotQuestionsRoute, qc.GetHot)
	questionsGetRouter.Use(middleware.AddContentType)

	questionsPostRouter := r.Methods(http.MethodPost).Subrouter()
	
	questionsPostRouter.HandleFunc(constants.CreateQuestionRoute, qc.Create)
	questionsPostRouter.HandleFunc(constants.CreateQuestionAnswerRoute, qc.CreateAnswer)

	questionsPostRouter.HandleFunc(constants.LikeQuestionRoute, qc.Like)
	questionsPostRouter.HandleFunc(constants.LikeQuestionUndoRoute, qc.LikeUndo)

	questionsPostRouter.HandleFunc(constants.DislikeQuestionRoute, qc.Dislike)
	questionsPostRouter.HandleFunc(constants.DislikeQuestionUndoRoute, qc.DislikeUndo)

	questionsPostRouter.Use(middleware.AddContentType)
	questionsPostRouter.Use(middleware.IsAuthorized)

	questionsDeleteRouter := r.Methods(http.MethodDelete).Subrouter()
	questionsDeleteRouter.HandleFunc(constants.DeleteQuestionRoute, qc.Delete)
	questionsDeleteRouter.Use(middleware.AddContentType)
	questionsDeleteRouter.Use(middleware.IsAuthorized)
	
	// answer routers
	answerPostRouter := r.Methods(http.MethodPost).Subrouter()
	answerPostRouter.HandleFunc(constants.LikeAnswerRoute, answc.Like)
	answerPostRouter.HandleFunc(constants.LikeAnswerUndoRoute, answc.LikeUndo)
	
	questionsPostRouter.HandleFunc(constants.DislikeAnswerRoute, answc.Dislike)
	questionsPostRouter.HandleFunc(constants.DislikeAnswerUndoRoute, answc.DislikeUndo)

	questionsPostRouter.Use(middleware.AddContentType)
	questionsPostRouter.Use(middleware.IsAuthorized)

	answerPutRouter := r.Methods(http.MethodPut).Subrouter()
	answerPutRouter.HandleFunc(constants.UpdateAnswerRoute, answc.Update)
	answerPutRouter.Use(middleware.AddContentType)
	answerPutRouter.Use(middleware.IsAuthorized)

	answerDeleteRouter := r.Methods(http.MethodDelete).Subrouter()
	answerDeleteRouter.HandleFunc(constants.DeleteAnswerRoute, answc.Delete)
	answerDeleteRouter.Use(middleware.AddContentType)
	answerDeleteRouter.Use(middleware.IsAuthorized)

	r.PathPrefix("/swagger/").Handler(http.StripPrefix("/swagger/", http.FileServer(http.Dir("./swaggerui/"))))
	
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{os.Getenv("CLIENT_APP")},
		AllowCredentials: true,
		AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	addr := os.Getenv("API_ADDRESS")
	srv := &http.Server {
		Handler: cors.Handler(r),
		Addr: addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}
	
	err = srv.ListenAndServe()

	if err != nil {
		log.Fatal(err.Error())
	}
}
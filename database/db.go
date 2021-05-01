package database

import (
	"log"

	"github.com/EnisMulic/Ask.it.Backend/domain"
	"gorm.io/gorm"
)


func Migrate(db *gorm.DB) {
	
	db.AutoMigrate(&domain.User{})
	db.AutoMigrate(&domain.Question{})
	db.AutoMigrate(&domain.Answer{})
	db.AutoMigrate(&domain.UserQuestionRating{})
	db.AutoMigrate(&domain.UserAnswerRating{})
	db.AutoMigrate(&domain.AnswerNotification{})

	err := db.SetupJoinTable(&domain.User{}, "QuestionRatings", &domain.UserQuestionRating{})

	if err != nil {
		log.Fatalf(err.Error())
	}

	err = db.SetupJoinTable(&domain.User{}, "AnswerRatings", &domain.UserAnswerRating{})

	if err != nil {
		log.Fatalf(err.Error())
	}
}
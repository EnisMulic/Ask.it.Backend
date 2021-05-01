package repositories

import (
	"github.com/EnisMulic/Ask.it.Backend/domain"
	"gorm.io/gorm"
)


type AnswerNotificationRepository struct {
	db *gorm.DB
}

func NewAnswerNotificationRepository(db *gorm.DB) *AnswerNotificationRepository {
	return &AnswerNotificationRepository{db}
}

func (anr *AnswerNotificationRepository) Create(notification domain.AnswerNotification) (domain.AnswerNotification, error) {
	result := anr.db.Create(&notification)
	return notification, result.Error
}
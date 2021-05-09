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

func (anr *AnswerNotificationRepository) GetById (id uint) (domain.AnswerNotification, error) {
	var notification domain.AnswerNotification
	result := anr.db.First(&notification, id)
	return notification, result.Error 
}

func (anr *AnswerNotificationRepository) Create(notification domain.AnswerNotification) (domain.AnswerNotification, error) {
	result := anr.db.Create(&notification)
	return notification, result.Error
}

func (anr *AnswerNotificationRepository) Update(notification domain.AnswerNotification) (domain.AnswerNotification, error) {
	result := anr.db.Model(&notification).Updates(domain.AnswerNotification{
		IsRead: true,
	})

	return notification, result.Error
}
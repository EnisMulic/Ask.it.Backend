package services

import (
	"github.com/EnisMulic/Ask.it.Backend/constants"
	"github.com/EnisMulic/Ask.it.Backend/repositories"
)


type AnswerNotificationService struct {
	anRepo *repositories.AnswerNotificationRepository
}

func NewAnswerNotificationService(anRepo *repositories.AnswerNotificationRepository) *AnswerNotificationService {
	return &AnswerNotificationService{anRepo}
}

func (ans *AnswerNotificationService) MarkRead (notificationId uint, userId uint) error {
	notification, err := ans.anRepo.GetById(notificationId)
	
	if err != nil {
		return constants.ErrGeneric
	}

	if notification.IsRead {
		return constants.ErrGeneric
	}

	if notification.UserID != userId {
		return constants.ErrForbidden
	}

	notification, err = ans.anRepo.Update(notification)

	if err != nil {
		return constants.ErrGeneric
	}

	return nil
}
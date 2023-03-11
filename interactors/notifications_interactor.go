package interactors

import (
	"github.com/pkg/errors"
	"microservice_clean_design/app/core"
	"microservice_clean_design/domain"
)

type NotificationsInteractor struct {
	log               core.Logger
	notificationsRepo domain.NotificationsRepository
}

func NewNotificationsUCase(log core.Logger, notificationsRepo domain.NotificationsRepository) *NotificationsInteractor {
	return &NotificationsInteractor{
		log:               log,
		notificationsRepo: notificationsRepo,
	}
}

func (i *NotificationsInteractor) GetAllNotifications(userId int32) (domain.GetAllNotificationsResponse, error) {
	notifications, err := i.notificationsRepo.All()
	if err != nil {
		return domain.GetAllNotificationsResponse{}, errors.Wrap(err, "Cannot fetch all notifications")
	}

	return domain.GetAllNotificationsResponse{
		StatusCode:    domain.Success,
		Notifications: notifications,
	}, nil
}

func (i *NotificationsInteractor) CreateNotification(name string) (domain.CreateNotificationResponse, error) {
	notification := &domain.Notification{
		Title: "Title",
	}
	err := i.notificationsRepo.Insert(notification)
	if err != nil {
		return domain.CreateNotificationResponse{}, errors.Wrap(err, "Cannot insert task")
	}

	return domain.CreateNotificationResponse{
		StatusCode: domain.Success,
		Id:         notification.Id,
	}, nil
}

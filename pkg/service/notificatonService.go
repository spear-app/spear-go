package service

import "github.com/spear-app/spear-go/pkg/domain/notification"



type NotificationService interface{
	Create(*notification.Notification) error
}

type DefaultNotificationService struct{
	repo notification.NotificationRepository
}

func (s DefaultNotificationService) Create(notificationObj *notification.Notification) error {
	return s.repo.Create(notificationObj)
}

func NewNotificationService(repository notification.NotificationRepository) DefaultNotificationService {
	return DefaultNotificationService{repo: repository}
}
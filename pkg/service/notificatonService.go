package service

import "github.com/spear-app/spear-go/pkg/domain/notification"



type NotificationService interface{
	Create(*notification.Notification) error
	ReadByNotificationID(int)(notification.Notification,error)
}

type DefaultNotificationService struct{
	repo notification.NotificationRepository
}

func (s DefaultNotificationService) Create(notificationObj *notification.Notification) error {
	return s.repo.Create(notificationObj)
}

func (s DefaultNotificationService) ReadByNotificationID(id int)(notification.Notification,error){
	return s.repo.ReadByNotificationID(id)
}

func NewNotificationService(repository notification.NotificationRepository) DefaultNotificationService {
	return DefaultNotificationService{repo: repository}
}
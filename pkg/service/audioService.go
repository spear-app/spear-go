package service

import "github.com/spear-app/spear-go/pkg/domain/notification"

type AudioService interface {
	SoundDetection(*notification.Notification) error
}

type DefaultAudioService struct {
	repo notification.NotificationRepository
}

func (s DefaultNotificationService) SoundDetection(notificationObj *notification.Notification) error {
	return s.repo.Create(notificationObj)
}

func NewAudioService(repository notification.NotificationRepository) DefaultNotificationService {
	return DefaultNotificationService{repo: repository}
}

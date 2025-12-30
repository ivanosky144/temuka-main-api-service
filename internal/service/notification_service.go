package service

import (
	"context"
	"errors"

	"github.com/temuka-api-service/internal/model"
	"github.com/temuka-api-service/internal/repository"
)

type NotificationService interface {
	GetNotificationsByUser(ctx context.Context, userID int) ([]model.Notification, error)
}

type NotificationServiceImpl struct {
	NotificationRepository repository.NotificationRepository
}

func NewNotificationService(notifRepo repository.NotificationRepository) NotificationService {
	return &NotificationServiceImpl{
		NotificationRepository: notifRepo,
	}
}

func (s *NotificationServiceImpl) GetNotificationsByUser(ctx context.Context, userID int) ([]model.Notification, error) {
	notifications, err := s.NotificationRepository.GetNotificationsByUserID(ctx, userID)
	if err != nil {
		return nil, errors.New("error retrieving notifications")
	}
	return notifications, nil
}

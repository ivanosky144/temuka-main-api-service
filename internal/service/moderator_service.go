package service

import (
	"context"
	"errors"
	"strconv"

	"github.com/temuka-api-service/internal/dto"
	"github.com/temuka-api-service/internal/model"
	"github.com/temuka-api-service/internal/repository"
)

type ModeratorService interface {
	SendModeratorRequest(ctx context.Context, data dto.SendModeratorRequest) error
	RemoveModerator(ctx context.Context, moderatorID int) error
}

type ModeratorServiceImpl struct {
	ModeratorRepository    repository.ModeratorRepository
	NotificationRepository repository.NotificationRepository
}

func NewModeratorService(moderatorRepo repository.ModeratorRepository, notificationRepo repository.NotificationRepository) ModeratorService {
	return &ModeratorServiceImpl{
		ModeratorRepository:    moderatorRepo,
		NotificationRepository: notificationRepo,
	}
}

func (s *ModeratorServiceImpl) SendModeratorRequest(ctx context.Context, data dto.SendModeratorRequest) error {
	notification := model.Notification{
		UserID:  data.CommunityMemberID,
		Type:    "request",
		Message: "You have been requested to be a moderator in community with ID " + strconv.Itoa(data.CommunityID),
	}

	if err := s.NotificationRepository.CreateNotification(ctx, &notification); err != nil {
		return errors.New("error creating moderator notification")
	}

	return nil
}

func (s *ModeratorServiceImpl) RemoveModerator(ctx context.Context, moderatorID int) error {
	if err := s.ModeratorRepository.DeleteModerator(ctx, moderatorID); err != nil {
		return errors.New("error removing moderator")
	}
	return nil
}

package repository

import (
	"context"
	"fmt"

	"github.com/temuka-api-service/internal/model"
	database "github.com/temuka-api-service/util/database"
)

type ModeratorRepository interface {
	CreateModerator(ctx context.Context, moderator *model.Moderator) error
	GetModeratorsByCommunityID(ctx context.Context, communityId int) ([]model.Moderator, error)
	DeleteModerator(ctx context.Context, id int) error
}

type ModeratorRepositoryImpl struct {
	db database.PostgresWrapper
}

func NewModeratorRepository(db database.PostgresWrapper) ModeratorRepository {
	return &ModeratorRepositoryImpl{
		db: db,
	}
}

func (r *ModeratorRepositoryImpl) CreateModerator(ctx context.Context, moderator *model.Moderator) error {
	if err := r.db.Create(ctx, moderator); err != nil {
		return fmt.Errorf("failed to create moderator: %w", err)
	}
	return nil
}

func (r *ModeratorRepositoryImpl) GetModeratorsByCommunityID(ctx context.Context, communityId int) ([]model.Moderator, error) {
	var moderators []model.Moderator

	if err := r.db.Where(ctx, &moderators, "community_id = ?", communityId); err != nil {
		return nil, fmt.Errorf("failed to get moderators: %w", err)
	}

	return moderators, nil
}

func (r *ModeratorRepositoryImpl) DeleteModerator(ctx context.Context, id int) error {
	if err := r.db.Delete(ctx, &model.Moderator{}, id); err != nil {
		return fmt.Errorf("failed to delete moderator: %w", err)
	}
	return nil
}

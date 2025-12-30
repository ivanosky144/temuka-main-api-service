package repository

import (
	"context"
	"fmt"

	"github.com/temuka-api-service/internal/model"
	database "github.com/temuka-api-service/util/database"
)

type LocationRepository interface {
	AddLocation(ctx context.Context, location *model.Location) error
	UpdateLocation(ctx context.Context, id int, location *model.Location) error
	GetLocations(ctx context.Context) ([]model.Location, error)
	DeleteLocation(ctx context.Context, id int) error
	GetLocationById(ctx context.Context, id int) (*model.Location, error)
}

type LocationRepositoryImpl struct {
	db database.PostgresWrapper
}

func NewLocationRepository(db database.PostgresWrapper) LocationRepository {
	return &LocationRepositoryImpl{
		db: db,
	}
}

func (r *LocationRepositoryImpl) AddLocation(ctx context.Context, location *model.Location) error {
	if err := r.db.Create(ctx, location); err != nil {
		return fmt.Errorf("failed to add location: %w", err)
	}
	return nil
}

func (r *LocationRepositoryImpl) DeleteLocation(ctx context.Context, id int) error {
	if err := r.db.Delete(ctx, &model.Location{}, id); err != nil {
		return fmt.Errorf("failed to delete location: %w", err)
	}
	return nil
}

func (r *LocationRepositoryImpl) UpdateLocation(ctx context.Context, id int, location *model.Location) error {
	q := r.db.Model(ctx, &model.Location{}).Where("id = ?", id)

	if err := q.Updates(location).Error; err != nil {
		return fmt.Errorf("failed to update location: %w", err)
	}

	return nil
}

func (r *LocationRepositoryImpl) GetLocations(ctx context.Context) ([]model.Location, error) {
	var locations []model.Location

	if err := r.db.Find(ctx, &locations); err != nil {
		return nil, fmt.Errorf("failed to get locations: %w", err)
	}

	return locations, nil
}

func (r *LocationRepositoryImpl) GetLocationById(ctx context.Context, id int) (*model.Location, error) {
	var location model.Location

	if err := r.db.First(ctx, &location, id); err != nil {
		return nil, fmt.Errorf("failed to get location by id: %w", err)
	}

	return &location, nil
}

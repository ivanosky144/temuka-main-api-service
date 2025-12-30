package repository

import (
	"context"
	"fmt"

	"github.com/temuka-api-service/internal/model"
	database "github.com/temuka-api-service/util/database"
)

type ReportRepository interface {
	CreateReport(ctx context.Context, report *model.Report) error
	DeleteReport(ctx context.Context, id int) error
}

type ReportRepositoryImpl struct {
	db database.PostgresWrapper
}

func NewReportRepository(db database.PostgresWrapper) ReportRepository {
	return &ReportRepositoryImpl{
		db: db,
	}
}

func (r *ReportRepositoryImpl) CreateReport(ctx context.Context, report *model.Report) error {
	if err := r.db.Create(ctx, report); err != nil {
		return fmt.Errorf("failed to create report: %w", err)
	}
	return nil
}

func (r *ReportRepositoryImpl) DeleteReport(ctx context.Context, id int) error {
	if err := r.db.Delete(ctx, &model.Report{}, id); err != nil {
		return fmt.Errorf("failed to delete report: %w", err)
	}
	return nil
}

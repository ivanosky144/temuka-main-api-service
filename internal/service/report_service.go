package service

import (
	"context"

	"github.com/temuka-api-service/internal/model"
	"github.com/temuka-api-service/internal/repository"
)

type ReportService interface {
	CreateReport(ctx context.Context, report *model.Report) error
	DeleteReport(ctx context.Context, id int) error
}

type ReportServiceImpl struct {
	ReportRepository repository.ReportRepository
}

func NewReportService(reportRepo repository.ReportRepository) ReportService {
	return &ReportServiceImpl{
		ReportRepository: reportRepo,
	}
}

func (s *ReportServiceImpl) CreateReport(ctx context.Context, report *model.Report) error {
	return s.ReportRepository.CreateReport(ctx, report)
}

func (s *ReportServiceImpl) DeleteReport(ctx context.Context, id int) error {
	return s.ReportRepository.DeleteReport(ctx, id)
}

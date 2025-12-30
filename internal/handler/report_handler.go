package handler

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/temuka-api-service/internal/dto"
	"github.com/temuka-api-service/internal/model"
	"github.com/temuka-api-service/internal/service"
	"github.com/temuka-api-service/util/rest"
)

type ReportHandler interface {
	CreateReport(w http.ResponseWriter, r *http.Request)
	DeleteReport(w http.ResponseWriter, r *http.Request)
}

type ReportHandlerImpl struct {
	ReportService service.ReportService
}

func NewReportHandler(reportService service.ReportService) ReportHandler {
	return &ReportHandlerImpl{
		ReportService: reportService,
	}
}

func (h *ReportHandlerImpl) CreateReport(w http.ResponseWriter, r *http.Request) {
	var request dto.CreateReportRequest

	if err := rest.ReadRequest(r, &request); err != nil {
		rest.WriteResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
		return
	}

	report := &model.Report{
		CommentID: &request.CommentID,
		PostID:    &request.PostID,
		Reason:    request.Reason,
	}

	if err := h.ReportService.CreateReport(r.Context(), report); err != nil {
		rest.WriteResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	response := dto.MessageResponse{
		Message: "Report has been created successfully",
	}
	rest.WriteResponse(w, http.StatusOK, response)
}

func (h *ReportHandlerImpl) DeleteReport(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	reportIDStr := vars["id"]

	reportID, err := strconv.Atoi(reportIDStr)
	if err != nil {
		rest.WriteResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid report ID"})
		return
	}

	if err := h.ReportService.DeleteReport(r.Context(), reportID); err != nil {
		rest.WriteResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	response := dto.MessageResponse{
		Message: "Report has been deleted",
	}
	rest.WriteResponse(w, http.StatusOK, response)
}

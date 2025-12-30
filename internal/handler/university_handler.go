package handler

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/temuka-api-service/internal/dto"
	"github.com/temuka-api-service/internal/service"
	"github.com/temuka-api-service/util/rest"
)

type UniversityHandler interface {
	AddUniversity(w http.ResponseWriter, r *http.Request)
	UpdateUniversity(w http.ResponseWriter, r *http.Request)
	GetUniversityDetail(w http.ResponseWriter, r *http.Request)
	GetUniversities(w http.ResponseWriter, r *http.Request)
	AddReview(w http.ResponseWriter, r *http.Request)
	GetUniversityReviews(w http.ResponseWriter, r *http.Request)
}

type UniversityHandlerImpl struct {
	UniversityService service.UniversityService
}

func NewUniversityHandler(service service.UniversityService) UniversityHandler {
	return &UniversityHandlerImpl{
		UniversityService: service,
	}
}

func (h *UniversityHandlerImpl) AddUniversity(w http.ResponseWriter, r *http.Request) {
	var req dto.AddUniversityRequest
	if err := rest.ReadRequest(r, &req); err != nil {
		rest.WriteResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
		return
	}

	university, err := h.UniversityService.AddUniversity(r.Context(), req)
	if err != nil {
		rest.WriteResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	rest.WriteResponse(w, http.StatusOK, map[string]interface{}{
		"message": "University has been added",
		"data":    university,
	})
}

func (h *UniversityHandlerImpl) UpdateUniversity(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		rest.WriteResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid university ID"})
		return
	}

	var req dto.UpdateUniversityRequest
	if err := rest.ReadRequest(r, &req); err != nil {
		rest.WriteResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
		return
	}

	university, err := h.UniversityService.UpdateUniversity(r.Context(), id, req)
	if err != nil {
		rest.WriteResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	rest.WriteResponse(w, http.StatusOK, map[string]interface{}{
		"message": "University has been updated",
		"data":    university,
	})
}

func (h *UniversityHandlerImpl) GetUniversityDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	slug := vars["slug"]

	university, err := h.UniversityService.GetUniversityDetail(r.Context(), slug)
	if err != nil {
		rest.WriteResponse(w, http.StatusNotFound, map[string]string{"error": err.Error()})
		return
	}

	rest.WriteResponse(w, http.StatusOK, map[string]interface{}{
		"message": "University detail retrieved",
		"data":    university,
	})
}

func (h *UniversityHandlerImpl) GetUniversities(w http.ResponseWriter, r *http.Request) {
	universities, err := h.UniversityService.GetUniversities(r.Context())
	if err != nil {
		rest.WriteResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	rest.WriteResponse(w, http.StatusOK, map[string]interface{}{
		"message": "Universities retrieved successfully",
		"data":    universities,
	})
}

func (h *UniversityHandlerImpl) AddReview(w http.ResponseWriter, r *http.Request) {
	var req dto.AddReviewRequest
	if err := rest.ReadRequest(r, &req); err != nil {
		rest.WriteResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
		return
	}

	review, err := h.UniversityService.AddReview(r.Context(), req)
	if err != nil {
		rest.WriteResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	rest.WriteResponse(w, http.StatusOK, map[string]interface{}{
		"message": "Review added successfully",
		"data":    review,
	})
}

func (h *UniversityHandlerImpl) GetUniversityReviews(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	universityIDStr := vars["university_id"]

	universityID, err := strconv.Atoi(universityIDStr)
	if err != nil {
		rest.WriteResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid university ID"})
		return
	}

	reviews, err := h.UniversityService.GetUniversityReviews(r.Context(), universityID)
	if err != nil {
		rest.WriteResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	rest.WriteResponse(w, http.StatusOK, map[string]interface{}{
		"message": "University reviews retrieved",
		"data":    reviews,
	})
}

package controller

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/temuka-api-service/internal/model"
	"github.com/temuka-api-service/internal/repository"
	httputil "github.com/temuka-api-service/pkg/http"
)

type UniversityController interface {
	AddUniversity(w http.ResponseWriter, r *http.Request)
	UpdateUniversity(w http.ResponseWriter, r *http.Request)
	GetUniversityDetail(w http.ResponseWriter, r *http.Request)
	GetUniversities(w http.ResponseWriter, r *http.Request)
	AddReview(w http.ResponseWriter, r *http.Request)
	GetUniversityReviews(w http.ResponseWriter, r *http.Request)
}

type UniversityControllerImpl struct {
	UniversityRepository repository.UniversityRepository
	ReviewRepository     repository.ReviewRepository
}

func NewUniversityController(universityRepo repository.UniversityRepository, reviewRepo repository.ReviewRepository) UniversityController {
	return &UniversityControllerImpl{
		UniversityRepository: universityRepo,
		ReviewRepository:     reviewRepo,
	}
}

func (c *UniversityControllerImpl) AddUniversity(w http.ResponseWriter, r *http.Request) {
	var requestBody struct {
		Name          string `json:"name"`
		Summary       string `json:"summary"`
		LocationID    int    `json:"location_id"`
		Website       string `json:"website"`
		Address       string `json:"address"`
		MinTuition    int    `json:"min_tuition"`
		MaxTuition    int    `json:"max_tuition"`
		TotalMajors   int    `json:"total_majors"`
		Logo          string `json:"logo"`
		Type          string `json:"type"`
		Accreditation string `json:"accreditation"`
	}

	if err := httputil.ReadRequest(r, &requestBody); err != nil {
		httputil.WriteResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
		return
	}

	newUniversity := model.University{
		Name:        requestBody.Name,
		Slug:        strings.ReplaceAll(strings.ToLower(requestBody.Name), " ", "_"),
		Summary:     requestBody.Summary,
		LocationID:  requestBody.LocationID,
		Website:     requestBody.Website,
		Address:     requestBody.Address,
		TotalMajors: &requestBody.TotalMajors,
		MinTuition:  requestBody.MinTuition,
		MaxTuition:  requestBody.MaxTuition,
		Logo:        requestBody.Logo,
	}

	if err := c.UniversityRepository.CreateUniversity(context.Background(), &newUniversity); err != nil {
		httputil.WriteResponse(w, http.StatusInternalServerError, map[string]string{"error": "Error creating new university"})
		return
	}

	response := struct {
		Message string           `json:"message"`
		Data    model.University `json:"data"`
	}{
		Message: "University has been added",
		Data:    newUniversity,
	}

	httputil.WriteResponse(w, http.StatusOK, response)
}

func (c *UniversityControllerImpl) UpdateUniversity(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	universityIDstr := vars["id"]

	universityID, err := strconv.Atoi(universityIDstr)
	if err != nil {
		httputil.WriteResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid university id"})
		return
	}

	university, err := c.UniversityRepository.GetUniversityDetailByID(context.Background(), universityID)
	if err != nil {
		httputil.WriteResponse(w, http.StatusNotFound, map[string]string{"error": "University not found"})
		return
	}

	var requestBody struct {
		Name          string `json:"name"`
		Summary       string `json:"summary"`
		LocationID    int    `json:"location_id"`
		Website       string `json:"website"`
		Address       string `json:"address"`
		MinTuition    int    `json:"min_tuition"`
		MaxTuition    int    `json:"max_tuition"`
		TotalMajors   int    `json:"total_majors"`
		Logo          string `json:"logo"`
		Type          string `json:"type"`
		Accreditation string `json:"accreditation"`
	}

	if err := httputil.ReadRequest(r, &requestBody); err != nil {
		httputil.WriteResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
		return
	}

	updatedUniversity := model.University{
		Name:          requestBody.Name,
		Slug:          strings.ReplaceAll(strings.ToLower(requestBody.Name), " ", "_"),
		Summary:       requestBody.Summary,
		LocationID:    requestBody.LocationID,
		Website:       requestBody.Website,
		Address:       requestBody.Address,
		TotalMajors:   &requestBody.TotalMajors,
		MinTuition:    requestBody.MinTuition,
		MaxTuition:    requestBody.MaxTuition,
		Type:          requestBody.Type,
		Accreditation: requestBody.Accreditation,
		Logo:          requestBody.Logo,
	}

	university.ID = universityID

	if err := c.UniversityRepository.UpdateUniversity(context.Background(), university.ID, &updatedUniversity); err != nil {
		httputil.WriteResponse(w, http.StatusInternalServerError, map[string]string{"error": "Error updating university"})
		return
	}

	response := struct {
		Message string           `json:"message"`
		Data    model.University `json:"data"`
	}{
		Message: "University has been updated",
		Data:    updatedUniversity,
	}

	httputil.WriteResponse(w, http.StatusOK, response)
}

func (c *UniversityControllerImpl) GetUniversityDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	universitySlugstr := vars["slug"]

	university, err := c.UniversityRepository.GetUniversityDetailBySlug(context.Background(), universitySlugstr)
	if err != nil {
		httputil.WriteResponse(w, http.StatusNotFound, map[string]string{"error": "University not found"})
		return
	}

	response := struct {
		Message string           `json:"message"`
		Data    model.University `json:"data"`
	}{
		Message: "University detail has been retrieved",
		Data:    *university,
	}

	httputil.WriteResponse(w, http.StatusOK, response)
}

func (c *UniversityControllerImpl) GetUniversities(w http.ResponseWriter, r *http.Request) {

	universities, err := c.UniversityRepository.GetUniversities(context.Background())
	if err != nil {
		httputil.WriteResponse(w, http.StatusOK, map[string]string{"error": "Universities not found"})
		return
	}

	response := struct {
		Message string             `json:"message"`
		Data    []model.University `json:"data"`
	}{
		Message: "Data has been retrieved successfully",
		Data:    universities,
	}

	httputil.WriteResponse(w, http.StatusOK, response)
}

func (c *UniversityControllerImpl) AddReview(w http.ResponseWriter, r *http.Request) {

	var requestBody struct {
		UserID       int    `json:"user_id"`
		UniversityID int    `json:"university_id"`
		Text         string `json:"text"`
		Rating       int    `json:"rating`
	}

	if err := httputil.ReadRequest(r, &requestBody); err != nil {
		httputil.WriteResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
		return
	}

	newUniversityReview := model.Review{
		UserID:       requestBody.UserID,
		UniversityID: requestBody.UniversityID,
		Text:         requestBody.Text,
	}

	if err := c.ReviewRepository.CreateReview(context.Background(), &newUniversityReview); err != nil {
		httputil.WriteResponse(w, http.StatusInternalServerError, map[string]string{"error": "Error creating new review"})
		return
	}

	university, err := c.UniversityRepository.GetUniversityDetailByID(context.Background(), requestBody.UniversityID)
	if err != nil {
		httputil.WriteResponse(w, http.StatusNotFound, map[string]string{"error": "University not found"})
		return
	}

	universityRating := 0
	if university.Rating != nil {
		universityRating = *university.Rating
	}

	universityTotalReviews := 0
	if university.TotalReviews != nil {
		universityTotalReviews = *university.TotalReviews
	}

	universityTotalReviews = universityTotalReviews + 1
	newUniversityRating := (universityRating*universityTotalReviews + requestBody.Rating) / (universityTotalReviews)

	university.Rating = &newUniversityRating
	university.TotalReviews = &universityTotalReviews

	if err := c.UniversityRepository.UpdateUniversity(context.Background(), requestBody.UniversityID, university); err != nil {
		httputil.WriteResponse(w, http.StatusInternalServerError, map[string]string{"error": "Failed to update university rating"})
		return
	}

	response := struct {
		Message string       `json:"message"`
		Data    model.Review `json:"data"`
	}{
		Message: "New review has been added",
		Data:    newUniversityReview,
	}

	httputil.WriteResponse(w, http.StatusOK, response)
}

func (c *UniversityControllerImpl) GetUniversityReviews(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	universityIDstr := vars["university_id"]

	universityID, err := strconv.Atoi(universityIDstr)
	if err != nil {
		httputil.WriteResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid university id"})
		return
	}

	universityReviews, err := c.ReviewRepository.GetReviewsByUniversityID(context.Background(), universityID)
	if err != nil {
		httputil.WriteResponse(w, http.StatusInternalServerError, map[string]string{"error": "Error retrieving university reviews"})
		return
	}

	response := struct {
		Message string         `json:"message"`
		Data    []model.Review `json:"data"`
	}{
		Message: "Data has been retrieved successfully",
		Data:    universityReviews,
	}

	httputil.WriteResponse(w, http.StatusOK, response)
}

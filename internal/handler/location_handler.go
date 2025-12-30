package handler

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/temuka-api-service/internal/dto"
	"github.com/temuka-api-service/internal/service"
	"github.com/temuka-api-service/util/rest"
)

type LocationHandler interface {
	AddLocation(w http.ResponseWriter, r *http.Request)
	UpdateLocation(w http.ResponseWriter, r *http.Request)
	GetLocations(w http.ResponseWriter, r *http.Request)
}

type LocationHandlerImpl struct {
	locationService service.LocationService
}

func NewLocationHandler(s service.LocationService) LocationHandler {
	return &LocationHandlerImpl{
		locationService: s,
	}
}

func (h *LocationHandlerImpl) AddLocation(w http.ResponseWriter, r *http.Request) {
	var req dto.AddLocationRequest

	if err := rest.ReadRequest(r, &req); err != nil {
		rest.WriteResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
		return
	}

	location, err := h.locationService.AddLocation(r.Context(), &req)
	if err != nil {
		rest.WriteResponse(w, http.StatusInternalServerError, map[string]string{"error": "Error creating new location"})
		return
	}

	response := dto.MessageResponse{
		Message: "Location has been added",
		Data:    location,
	}

	rest.WriteResponse(w, http.StatusOK, response)
}

func (h *LocationHandlerImpl) UpdateLocation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		rest.WriteResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid location id"})
		return
	}

	var req dto.UpdateLocationRequest
	if err := rest.ReadRequest(r, &req); err != nil {
		rest.WriteResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
		return
	}

	location, err := h.locationService.UpdateLocation(r.Context(), id, &req)
	if err != nil {
		rest.WriteResponse(w, http.StatusNotFound, map[string]string{"error": err.Error()})
		return
	}

	response := dto.MessageResponse{
		Message: "Location has been updated",
		Data:    location,
	}

	rest.WriteResponse(w, http.StatusOK, response)
}

func (h *LocationHandlerImpl) GetLocations(w http.ResponseWriter, r *http.Request) {
	locations, err := h.locationService.GetLocations(r.Context())
	if err != nil {
		rest.WriteResponse(w, http.StatusInternalServerError, map[string]string{"error": "Locations not found"})
		return
	}

	response := dto.MessageResponse{
		Message: "Data has been retrieved successfully",
		Data:    locations,
	}

	rest.WriteResponse(w, http.StatusOK, response)
}

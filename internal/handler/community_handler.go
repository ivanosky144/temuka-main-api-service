package handler

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/temuka-api-service/internal/dto"
	"github.com/temuka-api-service/internal/service"
	rest "github.com/temuka-api-service/util/rest"
)

type CommunityHandler interface {
	CreateCommunity(w http.ResponseWriter, r *http.Request)
	GetCommunities(w http.ResponseWriter, r *http.Request)
	UpdateCommunity(w http.ResponseWriter, r *http.Request)
	DeleteCommunity(w http.ResponseWriter, r *http.Request)
	JoinCommunity(w http.ResponseWriter, r *http.Request)
	GetCommunityPosts(w http.ResponseWriter, r *http.Request)
	GetCommunityDetail(w http.ResponseWriter, r *http.Request)
	GetUserJoinedCommunities(w http.ResponseWriter, r *http.Request)
}

type CommunityHandlerImpl struct {
	CommunityService service.CommunityService
}

func NewCommunityHandler(service service.CommunityService) CommunityHandler {
	return &CommunityHandlerImpl{CommunityService: service}
}

func (h *CommunityHandlerImpl) CreateCommunity(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateCommunityRequest
	if err := rest.ReadRequest(r, &req); err != nil {
		rest.WriteResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
		return
	}

	community, err := h.CommunityService.CreateCommunity(r.Context(), req)
	if err != nil {
		rest.WriteResponse(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	rest.WriteResponse(w, http.StatusOK, map[string]interface{}{
		"message": "Community has been created",
		"data":    community,
	})
}

func (h *CommunityHandlerImpl) GetCommunities(w http.ResponseWriter, r *http.Request) {
	communities, err := h.CommunityService.GetCommunities(r.Context())
	if err != nil {
		rest.WriteResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	rest.WriteResponse(w, http.StatusOK, map[string]interface{}{
		"message": "Communities have been retrieved",
		"data":    communities,
	})
}

func (h *CommunityHandlerImpl) UpdateCommunity(w http.ResponseWriter, r *http.Request) {
	var req dto.UpdateCommunityRequest
	if err := rest.ReadRequest(r, &req); err != nil {
		rest.WriteResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
		return
	}

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		rest.WriteResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid community ID"})
		return
	}

	community, err := h.CommunityService.UpdateCommunity(r.Context(), id, req)
	if err != nil {
		rest.WriteResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	rest.WriteResponse(w, http.StatusOK, map[string]interface{}{
		"message": "Community has been updated",
		"data":    community,
	})
}

func (h *CommunityHandlerImpl) DeleteCommunity(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		rest.WriteResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid community ID"})
		return
	}

	if err := h.CommunityService.DeleteCommunity(r.Context(), id); err != nil {
		rest.WriteResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	rest.WriteResponse(w, http.StatusOK, map[string]string{"message": "Community has been deleted"})
}

func (h *CommunityHandlerImpl) JoinCommunity(w http.ResponseWriter, r *http.Request) {
	var req dto.JoinCommunityRequest
	if err := rest.ReadRequest(r, &req); err != nil {
		rest.WriteResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
		return
	}

	id, err := strconv.Atoi(mux.Vars(r)["community_id"])
	if err != nil {
		rest.WriteResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid community ID"})
		return
	}

	if err := h.CommunityService.JoinCommunity(r.Context(), id, req); err != nil {
		rest.WriteResponse(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	rest.WriteResponse(w, http.StatusOK, map[string]string{"message": "Successfully joined the community"})
}

func (h *CommunityHandlerImpl) GetCommunityPosts(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		rest.WriteResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid community ID"})
		return
	}

	filters := make(map[string]interface{})
	if topic := r.URL.Query().Get("topic"); topic != "" {
		filters["topic"] = topic
	}
	if sort := r.URL.Query().Get("sort"); sort != "" {
		filters["sort"] = sort
	}
	if sortBy := r.URL.Query().Get("sort_by"); sortBy != "" {
		filters["sort_by"] = sortBy
	}

	posts, err := h.CommunityService.GetCommunityPosts(r.Context(), id, filters)
	if err != nil {
		rest.WriteResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	rest.WriteResponse(w, http.StatusOK, map[string]interface{}{
		"message": "Community posts have been retrieved",
		"data":    posts,
	})
}

func (h *CommunityHandlerImpl) GetCommunityDetail(w http.ResponseWriter, r *http.Request) {
	slug := mux.Vars(r)["slug"]
	community, err := h.CommunityService.GetCommunityDetail(r.Context(), slug)
	if err != nil {
		rest.WriteResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	rest.WriteResponse(w, http.StatusOK, map[string]interface{}{
		"message": "Community detail has been retrieved",
		"data":    community,
	})
}

func (h *CommunityHandlerImpl) GetUserJoinedCommunities(w http.ResponseWriter, r *http.Request) {
	var req dto.GetUserJoinedCommunitiesRequest
	if err := rest.ReadRequest(r, &req); err != nil {
		rest.WriteResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
		return
	}

	communities, err := h.CommunityService.GetUserJoinedCommunities(r.Context(), req)
	if err != nil {
		rest.WriteResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	rest.WriteResponse(w, http.StatusOK, map[string]interface{}{
		"message": "User communities have been retrieved",
		"data":    communities,
	})
}

package handler

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/temuka-api-service/internal/dto"
	"github.com/temuka-api-service/internal/service"
	"github.com/temuka-api-service/util/rest"
)

type UserHandler interface {
	SearchUsers(w http.ResponseWriter, r *http.Request)
	GetUserDetail(w http.ResponseWriter, r *http.Request)
	CreateUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	FollowUser(w http.ResponseWriter, r *http.Request)
	GetFollowers(w http.ResponseWriter, r *http.Request)
}

type UserHandlerImpl struct {
	UserService service.UserService
}

func NewUserHandler(userService service.UserService) UserHandler {
	return &UserHandlerImpl{
		UserService: userService,
	}
}

func (h *UserHandlerImpl) SearchUsers(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	data := dto.SearchUsersDTO{Name: name}
	users, err := h.UserService.SearchUsers(r.Context(), data)
	if err != nil {
		rest.WriteResponse(w, http.StatusNotFound, map[string]string{"error": err.Error()})
		return
	}

	response := struct {
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}{
		Message: "Search results retrieved successfully",
		Data:    users,
	}

	rest.WriteResponse(w, http.StatusOK, response)
}

func (h *UserHandlerImpl) GetUserDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIDstr := vars["id"]

	userID, err := strconv.Atoi(userIDstr)
	if err != nil {
		rest.WriteResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
		return
	}

	data := dto.GetUserDetailDTO{UserID: userID}
	user, err := h.UserService.GetUserDetail(r.Context(), data)
	if err != nil {
		rest.WriteResponse(w, http.StatusNotFound, map[string]string{"error": err.Error()})
		return
	}

	response := struct {
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}{
		Message: "User detail retrieved successfully",
		Data:    user,
	}

	rest.WriteResponse(w, http.StatusOK, response)
}

func (h *UserHandlerImpl) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateUserDTO
	if err := rest.ReadRequest(r, &req); err != nil {
		rest.WriteResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
		return
	}

	user, err := h.UserService.CreateUser(r.Context(), req)
	if err != nil {
		rest.WriteResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	response := struct {
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}{
		Message: "User created successfully",
		Data:    user,
	}

	rest.WriteResponse(w, http.StatusOK, response)
}

func (h *UserHandlerImpl) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIDstr := vars["id"]

	userID, err := strconv.Atoi(userIDstr)
	if err != nil {
		rest.WriteResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
		return
	}

	var req dto.UpdateUserDTO
	if err := rest.ReadRequest(r, &req); err != nil {
		rest.WriteResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
		return
	}

	req.UserID = userID

	if err := h.UserService.UpdateUser(r.Context(), req); err != nil {
		rest.WriteResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	response := map[string]string{"message": "User updated successfully"}
	rest.WriteResponse(w, http.StatusOK, response)
}

func (h *UserHandlerImpl) FollowUser(w http.ResponseWriter, r *http.Request) {
	var req dto.FollowUserDTO
	if err := rest.ReadRequest(r, &req); err != nil {
		rest.WriteResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
		return
	}

	if err := h.UserService.FollowUser(r.Context(), req); err != nil {
		rest.WriteResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	response := map[string]string{"message": "User followed successfully"}
	rest.WriteResponse(w, http.StatusOK, response)
}

func (h *UserHandlerImpl) GetFollowers(w http.ResponseWriter, r *http.Request) {
	var req dto.GetFollowersDTO
	if err := rest.ReadRequest(r, &req); err != nil {
		rest.WriteResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
		return
	}

	followers, err := h.UserService.GetFollowers(r.Context(), req)
	if err != nil {
		rest.WriteResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	response := struct {
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}{
		Message: "Followers list retrieved successfully",
		Data:    followers,
	}

	rest.WriteResponse(w, http.StatusOK, response)
}

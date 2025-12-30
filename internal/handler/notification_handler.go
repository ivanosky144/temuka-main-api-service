package handler

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/temuka-api-service/internal/dto"
	"github.com/temuka-api-service/internal/service"
	"github.com/temuka-api-service/util/rest"
)

type NotificationHandler interface {
	GetNotificationsByUser(w http.ResponseWriter, r *http.Request)
}

type NotificationHandlerImpl struct {
	NotificationService service.NotificationService
}

func NewNotificationHandler(notificationService service.NotificationService) NotificationHandler {
	return &NotificationHandlerImpl{
		NotificationService: notificationService,
	}
}

func (h *NotificationHandlerImpl) GetNotificationsByUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIDStr := vars["user_id"]

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		rest.WriteResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
		return
	}

	notifications, err := h.NotificationService.GetNotificationsByUser(r.Context(), userID)
	if err != nil {
		rest.WriteResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	response := dto.NotificationResponse{
		Message: "User notifications have been retrieved",
		Data:    notifications,
	}

	rest.WriteResponse(w, http.StatusOK, response)
}

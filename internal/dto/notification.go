package dto

import "github.com/temuka-api-service/internal/model"

type NotificationResponse struct {
	Message string               `json:"message"`
	Data    []model.Notification `json:"data"`
}

package dto

type MessageResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

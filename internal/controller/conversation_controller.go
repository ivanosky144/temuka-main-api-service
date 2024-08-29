package controller

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/temuka-api-service/internal/model"
	"github.com/temuka-api-service/internal/repository"
	httputil "github.com/temuka-api-service/pkg/http"
)

type ConversationController interface {
	AddConversation(w http.ResponseWriter, r *http.Request)
	DeleteConversation(w http.ResponseWriter, r *http.Request)
}

type ConversationControllerImpl struct {
	ConversationRepository repository.ConversationRepository
	UserRepository         repository.UserRepository
}

func NewConversationRepository(conversationRepo repository.ConversationRepository, userRepo repository.UserRepository) ConversationController {
	return &ConversationControllerImpl{
		ConversationRepository: conversationRepo,
		UserRepository:         userRepo,
	}
}

func (c *ConversationControllerImpl) AddConversation(w http.ResponseWriter, r *http.Request) {
	var requestBody struct {
		Title  string `json:"title"`
		UserID int    `json:"user_id"`
	}

	if err := httputil.ReadRequest(r, &requestBody); err != nil {
		httputil.WriteResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}

	newConversation := model.Conversation{
		UserID: requestBody.UserID,
		Title:  requestBody.Title,
	}

	if err := c.ConversationRepository.CreateConversation(context.Background(), &newConversation); err != nil {
		httputil.WriteResponse(w, http.StatusInternalServerError, map[string]string{"error": "Error creating comment"})
		return
	}

	response := struct {
		Message string             `json:"message"`
		Data    model.Conversation `json:"data"`
	}{
		Message: "Comment has been added",
		Data:    newConversation,
	}

	httputil.WriteResponse(w, http.StatusOK, response)
}

func (c *ConversationControllerImpl) DeleteConversation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	conversationIDstr := vars["id"]

	commentID, err := strconv.Atoi(conversationIDstr)
	if err != nil {
		httputil.WriteResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid comment id"})
		return
	}

	if err := c.ConversationRepository.DeleteConversation(context.Background(), commentID); err != nil {
		httputil.WriteResponse(w, http.StatusInternalServerError, map[string]string{"error": "Error deleting comment"})
		return
	}

	response := struct {
		Message string `json:"message"`
	}{
		Message: "Comment has been deleted",
	}

	httputil.WriteResponse(w, http.StatusOK, response)
}
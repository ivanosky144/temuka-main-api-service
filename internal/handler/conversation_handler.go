package handler

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/temuka-api-service/internal/dto"
	"github.com/temuka-api-service/internal/service"
	rest "github.com/temuka-api-service/util/rest"
)

type ConversationHandler interface {
	AddConversation(w http.ResponseWriter, r *http.Request)
	AddMessage(w http.ResponseWriter, r *http.Request)
	AddParticipant(w http.ResponseWriter, r *http.Request)
	GetConversationsByUserID(w http.ResponseWriter, r *http.Request)
	GetConversationDetail(w http.ResponseWriter, r *http.Request)
	DeleteConversation(w http.ResponseWriter, r *http.Request)
	RetrieveMessages(w http.ResponseWriter, r *http.Request)
}

type ConversationHandlerImpl struct {
	ConversationService service.ConversationService
}

func NewConversationHandler(conversationService service.ConversationService) ConversationHandler {
	return &ConversationHandlerImpl{ConversationService: conversationService}
}

func (h *ConversationHandlerImpl) AddConversation(w http.ResponseWriter, r *http.Request) {
	var req dto.AddConversationRequest
	if err := rest.ReadRequest(r, &req); err != nil {
		rest.WriteResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}

	conversation, err := h.ConversationService.AddConversation(r.Context(), req)
	if err != nil {
		rest.WriteResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	resp := dto.MessageResponse{Message: "Conversation has been created", Data: conversation}
	rest.WriteResponse(w, http.StatusOK, resp)
}

func (h *ConversationHandlerImpl) AddMessage(w http.ResponseWriter, r *http.Request) {
	var req dto.AddMessageRequest
	if err := rest.ReadRequest(r, &req); err != nil {
		rest.WriteResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}

	message, err := h.ConversationService.AddMessage(r.Context(), req)
	if err != nil {
		rest.WriteResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	resp := dto.MessageResponse{Message: "Message has been created", Data: message}
	rest.WriteResponse(w, http.StatusOK, resp)
}

func (h *ConversationHandlerImpl) AddParticipant(w http.ResponseWriter, r *http.Request) {
	var req dto.AddParticipantRequest
	if err := rest.ReadRequest(r, &req); err != nil {
		rest.WriteResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}

	if err := h.ConversationService.AddParticipant(r.Context(), req); err != nil {
		rest.WriteResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	resp := dto.MessageResponse{Message: "Participant has been added"}
	rest.WriteResponse(w, http.StatusOK, resp)
}

func (h *ConversationHandlerImpl) GetConversationsByUserID(w http.ResponseWriter, r *http.Request) {
	userIDstr := mux.Vars(r)["user_id"]
	userID, err := strconv.Atoi(userIDstr)
	if err != nil {
		rest.WriteResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid user id"})
		return
	}

	conversations, err := h.ConversationService.GetConversationsByUserID(r.Context(), userID)
	if err != nil {
		rest.WriteResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	resp := dto.MessageResponse{Message: "Conversations have been retrieved", Data: conversations}
	rest.WriteResponse(w, http.StatusOK, resp)
}

func (h *ConversationHandlerImpl) GetConversationDetail(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		rest.WriteResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid conversation id"})
		return
	}

	conversation, err := h.ConversationService.GetConversationDetail(r.Context(), id)
	if err != nil {
		rest.WriteResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	resp := dto.MessageResponse{Message: "Conversation detail has been retrieved", Data: conversation}
	rest.WriteResponse(w, http.StatusOK, resp)
}

func (h *ConversationHandlerImpl) DeleteConversation(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		rest.WriteResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid conversation id"})
		return
	}

	if err := h.ConversationService.DeleteConversation(r.Context(), id); err != nil {
		rest.WriteResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	resp := dto.MessageResponse{Message: "Conversation has been deleted"}
	rest.WriteResponse(w, http.StatusOK, resp)
}

func (h *ConversationHandlerImpl) RetrieveMessages(w http.ResponseWriter, r *http.Request) {
	conversationIDstr := mux.Vars(r)["conversation_id"]
	conversationID, err := strconv.Atoi(conversationIDstr)
	if err != nil {
		rest.WriteResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid conversation id"})
		return
	}

	messages, err := h.ConversationService.RetrieveMessages(r.Context(), conversationID)
	if err != nil {
		rest.WriteResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	resp := dto.MessageResponse{Message: "Messages have been retrieved", Data: messages}
	rest.WriteResponse(w, http.StatusOK, resp)
}

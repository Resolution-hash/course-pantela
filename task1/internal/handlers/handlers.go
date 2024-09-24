package handlers

import (
	"context"

	"github.com/Resolution-hash/course-pantela/task1/internal/messageService"
	"github.com/Resolution-hash/course-pantela/task1/internal/web/messages"
)

type Handler struct {
	Service *messageService.MessageService
}

// DeleteMessagesId implements messages.StrictServerInterface.
func (h *Handler) DeleteMessagesId(ctx context.Context, request messages.DeleteMessagesIdRequestObject) (messages.DeleteMessagesIdResponseObject, error) {
	messageID := request.Id

	err := h.Service.DeleteMessageByID(messageID)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// PatchMessagesId implements messages.StrictServerInterface.
func (h *Handler) PatchMessagesId(ctx context.Context, request messages.PatchMessagesIdRequestObject) (messages.PatchMessagesIdResponseObject, error) {
	messageRequest := request.Body
	messageID := request.Id

	messageToUpdate := messageService.Message{Text: *messageRequest.Message}
	_, err := h.Service.UpdateMessageByID(messageID, messageToUpdate)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// GetMessages implements messages.StrictServerInterface.
func (h *Handler) GetMessages(_ context.Context, _ messages.GetMessagesRequestObject) (messages.GetMessagesResponseObject, error) {
	allMessages, err := h.Service.GetAllMessage()
	if err != nil {
		return nil, err
	}

	response := messages.GetMessages200JSONResponse{}

	for _, msg := range allMessages {
		message := messages.Message{
			Id:      &msg.ID,
			Message: &msg.Text,
		}
		response = append(response, message)
	}

	return response, nil
}

// PostMessages implements messages.StrictServerInterface.
func (h *Handler) PostMessages(_ context.Context, request messages.PostMessagesRequestObject) (messages.PostMessagesResponseObject, error) {
	messageRequest := request.Body
	println("messageRequest", messageRequest)
	messageToCreate := messageService.Message{Text: *messageRequest.Message}
	println("messageToCreate", messageToCreate.Text)
	createdMessage, err := h.Service.CreateMessage(messageToCreate)
	if err != nil {
		return nil, err
	}

	response := messages.PostMessages201JSONResponse{
		Id:      &createdMessage.ID,
		Message: &createdMessage.Text,
	}

	return response, nil
}

func NewHandler(service *messageService.MessageService) *Handler {
	return &Handler{Service: service}
}

// func (h *Handler) GetMesssageHandler(w http.ResponseWriter, r *http.Request) {
// 	messages, err := h.Service.GetAllMessage()
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	w.Header().Set("Content-Type", "application/json")

// 	err = json.NewEncoder(w).Encode(messages)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}
// }

// func (h *Handler) PostMessageHandler(w http.ResponseWriter, r *http.Request) {
// 	var message messageService.Message
// 	err := json.NewDecoder(r.Body).Decode(&message)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	createdMessage, err := h.Service.CreateMessage(message)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	err = json.NewEncoder(w).Encode(createdMessage)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}
// }

// func (h *Handler) DeleteMessageHandler(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	stringId := vars["id"]

// 	id, err := strconv.ParseInt(stringId, 10, 32)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	err = h.Service.DeleteMessageByID(int(id))
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusNoContent)
// }

// func (h *Handler) PatchMessageHandler(w http.ResponseWriter, r *http.Request) {
// 	var message messageService.Message
// 	err := json.NewDecoder(r.Body).Decode(&message)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	vars := mux.Vars(r)
// 	stringId := vars["id"]

// 	id, err := strconv.ParseInt(stringId, 10, 0)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	updatedMessage, err := h.Service.UpdateMessageByID(int(id), message)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	err = json.NewEncoder(w).Encode(updatedMessage)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}
// }
